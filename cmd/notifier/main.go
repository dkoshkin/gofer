package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dkoshkin/gofer/pkg/dependency"
	"github.com/dkoshkin/gofer/pkg/dependency/manager"
	"github.com/dkoshkin/gofer/pkg/notifier"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

var (
	version   string
	buildDate string
)

const (
	dependenciesYAMLEnv = "DEPENDENCIES_YAML"

	datastoreCredentialsBase64Env = "DATASTORE_CREDENTIALS_BASE64"
	datastoreProjectIDEnv         = "DATASTORE_PROJECT_ID"
	datastoreCollectionEnv        = "DATASTORE_COLLECTION"
	datastoreDocEnv               = "DATASTORE_DOC"

	sendgridAPIKeyEnv      = "SENDGRID_API_KEY"
	notifierSenderNameEnv  = "NOTIFIER_SENDER_NAME"
	notifierSenderEmailEnv = "NOTIFIER_SENDER_EMAIL"
	notifierSubjectEnv     = "NOTIFIER_SUBJECT"
	notifierContactsEnv    = "NOTIFIER_CONTACTS"
)

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handler Triggered")

	var manifest dependency.Manifest
	err := json.NewDecoder(r.Body).Decode(&manifest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := run(&manifest)
	if err != nil {
		http.Error(w, fmt.Errorf("got an error: %v", err).Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func run(manifest *dependency.Manifest) (*dependency.Manifest, error) {
	projectID, collection, doc, err := checkDatastoreEnvs()
	if err != nil {
		return nil, fmt.Errorf("error reading env: %v", err)
	}
	sendgridAPIKey, notifierSenderName, notifierSenderEmail, notifierSubject, contacts, err := checkNotifierEnvs()
	if err != nil {
		return nil, fmt.Errorf("error reading env: %v", err)
	}

	var rw manager.ReadWriter
	credentialsBase64Bytes := os.Getenv(datastoreCredentialsBase64Env)
	if len(credentialsBase64Bytes) != 0 {
		credentialsJSONBytes, err := base64.StdEncoding.DecodeString(credentialsBase64Bytes)
		if err != nil {
			return nil, err
		}
		rw, err = manager.NewFirestoreManagerWithCredentialsJSON(projectID, collection, doc, credentialsJSONBytes)
	} else {
		rw, err = manager.NewFirestoreManager(projectID, collection, doc)
	}
	if err != nil {
		return nil, fmt.Errorf("error setting up datastore: %v", err)
	}

	_, err = rw.Init("", manifest.Dependencies...)
	if err != nil {
		return nil, fmt.Errorf("error initializing dependencies: %v", err)
	}

	newDependencies, updatedDependencies, existingDependencies, err := findDifferences(rw)
	if err != nil {
		return nil, err
	}

	err = notifier.NewEmailNotifier(sendgridAPIKey, notifierSenderName, notifierSenderEmail, notifierSubject, contacts).Send(newDependencies, updatedDependencies)
	if err != nil {
		return nil, fmt.Errorf("error sending with notifier: %v", err)
	}

	updated, err := updateInStore(rw, newDependencies, updatedDependencies, existingDependencies)
	if err != nil {
		return nil, fmt.Errorf("error updating dependencies in the store: %v", err)
	}

	return updated, nil
}

func checkDatastoreEnvs() (projectID string, collection string, doc string, err error) {
	if projectID = os.Getenv(datastoreProjectIDEnv); projectID == "" {
		err = fmt.Errorf("env %s must be set", datastoreProjectIDEnv)
		return
	}
	if collection = os.Getenv(datastoreCollectionEnv); collection == "" {
		err = fmt.Errorf("env %s must be set", datastoreCollectionEnv)
		return
	}
	if doc = os.Getenv(datastoreDocEnv); doc == "" {
		err = fmt.Errorf("env %s must be set", datastoreDocEnv)
		return
	}
	return
}

func checkNotifierEnvs() (sendgridAPIKey string, notifierSenderName string, notifierSenderEmail string, notifierSubject string, contacts []notifier.Contacts, err error) {
	if sendgridAPIKey = os.Getenv(sendgridAPIKeyEnv); sendgridAPIKey == "" {
		err = fmt.Errorf("env %s must be set", sendgridAPIKeyEnv)
		return
	}
	if notifierSenderName = os.Getenv(notifierSenderNameEnv); notifierSenderName == "" {
		err = fmt.Errorf("env %s must be set", notifierSenderNameEnv)
		return
	}
	if notifierSenderEmail = os.Getenv(notifierSenderEmailEnv); notifierSenderEmail == "" {
		err = fmt.Errorf("env %s must be set", notifierSenderEmailEnv)
		return
	}
	if notifierSubject = os.Getenv(notifierSubjectEnv); notifierSubject == "" {
		err = fmt.Errorf("env %s must be set", notifierSubjectEnv)
		return
	}
	var contactsString string
	if contactsString = os.Getenv(notifierContactsEnv); contactsString == "" {
		err = fmt.Errorf("env %s must be set", notifierContactsEnv)
		return
	}
	for _, contactString := range strings.Split(contactsString, "|") {
		contact := strings.Split(contactString, ":")
		contacts = append(contacts, notifier.Contacts{
			Name:    contact[0],
			Address: contact[1],
		})
	}

	return
}

func findDifferences(rw manager.ReadWriter) ([]dependency.Spec, []dependency.Spec, []dependency.Spec, error) {
	manifest, err := rw.Read()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error reading from datastore: %v", err)
	}
	_, dependenciesMap, err := manifest.ToMap()
	if err != nil {
		return nil, nil, nil, err
	}

	updatedManifest, err := manifest.Latest()
	if err != nil {
		fmt.Errorf("error getting updating dependencies: %v", err)
	}

	_, updatedDependenciesMap, err := updatedManifest.ToMap()
	if err != nil {
		return nil, nil, nil, err
	}

	newDependencies := make([]dependency.Spec, 0)
	updatedDependencies := make([]dependency.Spec, 0)
	existingDependencies := make([]dependency.Spec, 0)

	for k, v := range dependenciesMap {
		if dependency, ok := updatedDependenciesMap[k]; !ok {
			newDependencies = append(newDependencies, v)
		} else if v.Version != dependency.LatestVersion {
			updatedDependencies = append(updatedDependencies, dependency)
		} else {
			existingDependencies = append(existingDependencies, dependency)
		}
	}

	return newDependencies, updatedDependencies, existingDependencies, nil
}

func updateInStore(rw manager.ReadWriter, newDependencies []dependency.Spec, updatedDependencies []dependency.Spec, existingDependencies []dependency.Spec) (*dependency.Manifest, error) {
	dependencies := make([]dependency.Spec, 0)
	dependencies = append(dependencies, newDependencies...)
	dependencies = append(dependencies, updatedDependencies...)
	dependencies = append(dependencies, existingDependencies...)

	manifest := dependency.Manifest{Dependencies: dependencies}
	return &manifest, rw.Write(manifest)
}

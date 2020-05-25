package main

import (
	"fmt"
	"github.com/dkoshkin/gofer/pkg/dependency"
	"github.com/dkoshkin/gofer/pkg/dependency/manager"
	"github.com/dkoshkin/gofer/pkg/notifier"
	"os"
	"reflect"
)

var (
	version   string
	buildDate string
)

const (
	dependenciesYAMLEnv = "DEPENDENCIES_YAML"

	credentialsJSONEnv = "TEST_CREDENTIALS_JSON"

	projectIDEnv           = "PROJECT_ID"
	datastoreCollectionEnv = "DATASTORE_COLLECTION"
	datastoreDocEnv        = "DATASTORE_DOC"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading env: %v", err)
		os.Exit(1)
	}
}

func run() error {
	projectID, collection, doc, err := checkDatastoreEnvs()
	if err != nil {
		return fmt.Errorf("error reading env: %v", err)
	}

	var rw manager.ReadWriter
	credentialsJSONBytes := []byte(os.Getenv(credentialsJSONEnv))
	if len(credentialsJSONBytes) != 0 {
		rw, err = manager.NewFirestoreManagerWithCredentialsJSON(projectID, collection, doc, credentialsJSONBytes)
	} else {
		rw, err = manager.NewFirestoreManager(projectID, collection, doc)
	}
	if err != nil {
		return fmt.Errorf("error setting up datastore: %v", err)
	}

	dependenciesJSON := os.Getenv(dependenciesYAMLEnv)
	if dependenciesJSON != "" {
		manifest, err := dependency.FromBytes([]byte(dependenciesJSON))
		if err != nil {
			return fmt.Errorf("error unmarshalling from bytes: %v", err)
		}
		_, err = rw.Init("", manifest.Dependencies...)
		if err != nil {
			return fmt.Errorf("error initializing dependencies: %v", err)
		}
	}

	newDependencies, updatedDependencies, err := findDifferences(rw)
	if err != nil {
		return err
	}

	err = notifier.NewIOWriter(os.Stdout).Send(newDependencies, updatedDependencies)
	if err != nil {
		return fmt.Errorf("error sending with notifier: %v", err)
	}

	err = updateInStore(rw, newDependencies, updatedDependencies)
	if err != nil {
		return fmt.Errorf("error updating dependencies in the store: %v", err)
	}

	return nil
}

func checkDatastoreEnvs() (projectID string, collection string, doc string, err error) {
	if projectID = os.Getenv(projectIDEnv); projectID == "" {
		err = fmt.Errorf("env %s must be set", projectIDEnv)
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

func findDifferences(rw manager.ReadWriter) ([]dependency.Spec, []dependency.Spec, error) {
	manifest, err := rw.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("error reading from datastore: %v", err)
	}
	_, dependenciesMap, err := manifest.ToMap()
	if err != nil {
		return nil, nil, err
	}

	updatedManifest, err := manifest.Latest()
	if err != nil {
		fmt.Errorf("error getting updating dependencies: %v", err)
	}

	_, updatedDependenciesMap, err := updatedManifest.ToMap()
	if err != nil {
		return nil, nil, err
	}

	newDependencies := make([]dependency.Spec, 0)
	updatedDependencies := make([]dependency.Spec, 0)

	for k, v := range dependenciesMap {
		if dependency, ok := updatedDependenciesMap[k]; !ok {
			newDependencies = append(newDependencies, v)
		} else if !reflect.DeepEqual(v, dependency) {
			updatedDependencies = append(updatedDependencies, dependency)
		}
	}

	return newDependencies, updatedDependencies, nil
}

func updateInStore(rw manager.ReadWriter, newDependencies []dependency.Spec, updatedDependencies []dependency.Spec) error {
	dependencies := make([]dependency.Spec, 0)
	dependencies = append(dependencies, newDependencies...)
	dependencies = append(dependencies, updatedDependencies...)

	manifest := dependency.Manifest{Dependencies: dependencies}
	return rw.Write(manifest)
}

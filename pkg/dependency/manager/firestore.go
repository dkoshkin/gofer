package manager

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/dkoshkin/gofer/pkg/dependency"
	"google.golang.org/api/option"
	"strings"
)

// FirestoreManager persists the manifest to a Firebase database
type FirestoreManager struct {
	client     *firestore.Client
	collection string
	doc        string
}

func NewFirestoreManager(projectID, collection, doc string) (ReadWriter, error) {
	client, err := newClient(projectID)
	if err != nil {
		return nil, err
	}
	return FirestoreManager{client: client, collection: collection, doc: doc}, nil
}

func NewFirestoreManagerWithCredentialsJSON(projectID, collection, doc string, credentialsJSON []byte) (ReadWriter, error) {
	client, err := newClient(projectID, option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		return nil, err
	}

	return FirestoreManager{client: client, collection: collection, doc: doc}, nil
}

func newClient(projectID string, opts ...option.ClientOption) (*firestore.Client, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not create client: %v", err)
	}

	return client, nil
}

func (m FirestoreManager) Init(apiVersion string, dependencies ...dependency.Spec) (*dependency.Manifest, error) {
	dependenciesMap, err := m.readFromFirestore()
	if err != nil && !notFoundErr(err) {
		return nil, err
	}

	newDependenciesMap := make(map[string]dependency.Spec, 0)
	for _, dep := range dependencies {
		hash, err := dep.Hash()
		if err != nil {
			return nil, err
		}
		if _, ok := dependenciesMap[hash]; dependenciesMap == nil || !ok {
			newDependenciesMap[hash] = dep
		}
	}

	err = m.writeToFirestore(newDependenciesMap)
	if err != nil {
		return nil, err
	}

	return &dependency.Manifest{APIVersion: apiVersion, Dependencies: dependencies}, nil
}

func (m FirestoreManager) Read() (*dependency.Manifest, error) {
	dependenciesMap, err := m.readFromFirestore()
	if err != nil {
		return nil, err
	}
	manifest := &dependency.Manifest{}
	// TODO set version
	return manifest.FromMap("", dependenciesMap), nil
}

func (m FirestoreManager) Write(manifest dependency.Manifest) error {
	_, dependenciesMap, err := manifest.ToMap()
	if err != nil {
		return err
	}
	return m.writeToFirestore(dependenciesMap)
}

func (m FirestoreManager) readFromFirestore() (map[string]dependency.Spec, error) {
	doc := m.client.Collection(m.collection).Doc(m.doc)
	docsnap, err := doc.Get(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error reading from firestore: %v", err)
	}

	var dependenciesMap map[string]dependency.Spec
	if err = docsnap.DataTo(&dependenciesMap); err != nil {
		return nil, fmt.Errorf("error converting data: %v", err)
	}

	return dependenciesMap, nil
}

func (m FirestoreManager) writeToFirestore(dependenciesMap map[string]dependency.Spec) error {
	if len(dependenciesMap) == 0 {
		return nil
	}
	doc := m.client.Collection(m.collection).Doc(m.doc)
	if _, err := doc.Set(context.TODO(), dependenciesMap); err != nil {
		return fmt.Errorf("error writing to firestore: %v", err)
	}

	return nil
}

func notFoundErr(err error) bool {
	return strings.Contains(err.Error(), "rpc error: code = NotFound")
}

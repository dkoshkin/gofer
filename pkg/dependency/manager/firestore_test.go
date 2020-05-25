package manager

import (
	"encoding/base64"
	"fmt"
	"github.com/dkoshkin/gofer/pkg/dependency"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

const (
	datastoreCredentialsBase64Env = "DATASTORE_CREDENTIALS_BASE64"
	datastoreProjectIDEnv         = "DATASTORE_PROJECT_ID"

	collection = "dependencies-test"
)

func TestWriteRead(t *testing.T) {
	var credentialsJSONBytes []byte
	credentialsBase64Bytes := os.Getenv(datastoreCredentialsBase64Env)
	if len(credentialsBase64Bytes) == 0 {
		t.Fatalf("%s is not set", datastoreCredentialsBase64Env)
	} else {
		var err error
		credentialsJSONBytes, err = base64.StdEncoding.DecodeString(credentialsBase64Bytes)
		if err != nil {
			t.Fatal(err)
		}
	}
	projectID := os.Getenv(datastoreProjectIDEnv)
	if len(projectID) == 0 {
		t.Fatalf("%s is not set", datastoreProjectIDEnv)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rw, err := NewFirestoreManagerWithCredentialsJSON(projectID, collection, fmt.Sprintf("test-%d", r.Int()), credentialsJSONBytes)
	if err != nil {
		t.Fatal(err)
	}

	manifest := dependency.Manifest{
		Dependencies: []dependency.Spec{
			{
				Name:    "busybox",
				Type:    dependency.DockerType,
				Version: "1.28.1",
				Mask:    "1.28.[0-9]+",
			},
			{
				Name:    "https://github.com/kubernetes/kubernetes",
				Type:    dependency.GithubType,
				Version: "v1.17.5",
				Mask:    "v1.17.[0-9]+",
			},
		},
	}

	err = rw.Write(manifest)
	if err != nil {
		t.Fatal(err)
	}
	out, err := rw.Read()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*out, manifest) {
		t.Errorf("expected to be: %+v, got %+v", manifest, *out)
	}

	// write again updating one of the dependencies
	updatedManifest := dependency.Manifest{
		Dependencies: []dependency.Spec{
			{
				Name:    "busybox",
				Type:    dependency.DockerType,
				Version: "1.28.1",
				Mask:    "1.28.[0-9]+",
			},
			{
				Name:    "https://github.com/kubernetes/kubernetes",
				Type:    dependency.GithubType,
				Version: "v1.17.6",
				Mask:    "v1.17.[0-9]+",
			},
		},
	}
	err = rw.Write(updatedManifest)
	if err != nil {
		t.Fatal(err)
	}
	out, err = rw.Read()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*out, updatedManifest) {
		t.Errorf("expected to be: %+v, got %+v", updatedManifest, *out)
	}

}

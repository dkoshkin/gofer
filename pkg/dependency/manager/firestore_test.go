package manager

import (
	"fmt"
	"github.com/dkoshkin/gofer/pkg/dependency"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

const (
	credentialsJSONEnv = "FIRESTORE_TEST_CREDENTIALS_JSON"

	projectID           = "gofer-278221"
	credentialsJSONFile = "gofer-278221-4a60a7afccbd.json"

	collection = "dependencies-test"
)

func TestWriteRead(t *testing.T) {
	credentialsJSONBytes := []byte(os.Getenv(credentialsJSONEnv))
	if len(credentialsJSONBytes) == 0 {
		var err error
		credentialsJSONBytes, err = ioutil.ReadFile(filepath.Join("../../../", "hack", "env", credentialsJSONFile))
		if err != nil {
			t.Fatalf("%s is not set and got error trying to read file: %v", credentialsJSONEnv, err)
		}
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

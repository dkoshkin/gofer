package manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const (
	version = "v100.0.0"
)

func TestInit(t *testing.T) {
	f := filepath.Join(os.TempDir(), fmt.Sprintf("goferfiletesthappy%d", time.Now().UnixNano()))
	mngr := NewFileManager(f)
	manifest, err := mngr.Init(version)
	if err != nil {
		t.Fatalf("unexpected error running init: %v", err)
	}
	if manifest.APIVersion != version {
		t.Fatalf("expected version %q, instead got %v", version, manifest.APIVersion)
	}
	if len(manifest.Dependencies) != 0 {
		t.Fatalf("expected dependency length to be 0, instead got %d", len(manifest.Dependencies))
	}
}

func TestInitDirectoryExists(t *testing.T) {
	f := filepath.Join(os.TempDir(), fmt.Sprintf("goferfiletesthappy%d", time.Now().UnixNano()))
	mngr := NewFileManager(f)
	manifest, err := mngr.Init(version)
	if err != nil {
		t.Fatalf("unexpected error running init: %v", err)
	}
	if manifest.APIVersion != version {
		t.Fatalf("expected version %q, instead got %v", version, manifest.APIVersion)
	}
	if len(manifest.Dependencies) != 0 {
		t.Fatalf("expected dependency length to be 0, instead got %d", len(manifest.Dependencies))
	}
}
func TestInitFileExists(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "goferfiletestfileexists")
	if err != nil {
		t.Fatal(err)
	}
	mngr := NewFileManager(f.Name())
	_, err = mngr.Init(version)
	if err == nil {
		t.Fatalf("expected an error running init")
	}
	if _, ok := err.(*FileExistsError); !ok {
		t.Fatal("expected error to be of type 'FileExistsError'")
	}
}

func TestInitFileExistsAsDirectory(t *testing.T) {
	f, err := ioutil.TempDir(os.TempDir(), "goferfiletestfileexistsasdirectory")
	if err != nil {
		t.Fatal(err)
	}
	mngr := NewFileManager(f)
	_, err = mngr.Init(version)
	if err == nil {
		t.Fatalf("expected an error running init")
	}
	if !strings.Contains(err.Error(), "found a directory, could not create file") {
		t.Fatalf("expected error to contain 'found a directory, could not create file', instead got %q", err.Error())
	}
}

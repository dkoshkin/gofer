package manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dkoshkin/gofer/pkg/dependency"
	yaml "gopkg.in/yaml.v2"
)

// New returns an error that formats as the given text.
func NewFileExistsError(file string) error {
	return &FileExistsError{file}
}

// FileExistsError is an error returned when a file with the same name as the config already exists
type FileExistsError struct {
	file string
}

func (e *FileExistsError) Error() string {
	return fmt.Sprintf("file %q already exists, refusing to overwrite", e.file)
}

// FileManager persists the manifest to a file
type FileManager struct {
	filePath string
}

func NewFileManager(filePath string) ReadWriter {
	return FileManager{filePath: filePath}
}

func (m FileManager) Init(apiVersion string) (*dependency.Manifest, error) {
	manifestFile := filepath.Clean(m.filePath)
	fi, err := os.Stat(manifestFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("could not determine if file %q exists", manifestFile)
		}
	}
	// error is empty, already exists
	if err == nil {
		if fi.IsDir() {
			return nil, fmt.Errorf("found a directory, could not create file %q", manifestFile)
		}
		return nil, NewFileExistsError(manifestFile)
	}
	// continue if file does not exist and try to create
	// check if we should create a dir first
	dir := filepath.Dir(m.filePath)
	fi, err = os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("could not determine if file %q exists", manifestFile)
		}
		// dir does not exist, try to create it
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("could not create directory %q: %v", dir, err)
		}
	}
	f, err := os.Create(manifestFile)
	if err != nil {
		return nil, fmt.Errorf("could not create new manifest file: %v", err)
	}
	defer f.Close()

	manifest := &dependency.Manifest{APIVersion: apiVersion}

	data, err := yaml.Marshal(manifest)
	if err != nil {
		return nil, fmt.Errorf("could not marshal new manifest file: %v", err)
	}
	_, err = f.Write(data)
	if err != nil {
		return nil, fmt.Errorf("could not write new manifest file: %v", err)
	}

	return manifest, nil
}

// Read file and unmarshal the yaml
func (m FileManager) Read() (*dependency.Manifest, error) {
	manifestFile, err := validFilepath(m.filePath)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return nil, fmt.Errorf("could not read manifest file: %v", err)
	}

	manifest := &dependency.Manifest{}
	err = yaml.Unmarshal(data, manifest)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal manifest file: %v", err)
	}
	return manifest, nil
}

// Marshal and write to yaml file
func (m FileManager) Write(manifest dependency.Manifest) error {
	_, err := validFilepath(m.filePath)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(m.filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("could not open manifest file %v", err)
	}
	defer file.Close()

	// marshal file and write to it
	data, err := yaml.Marshal(manifest)
	if err != nil {
		return fmt.Errorf("could not marshal manifest data: %v", err)
	}
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("could not write to manifest file: %v", err)
	}

	// save changes
	err = file.Sync()
	if err != nil {
		return fmt.Errorf("could not save changes to manifest file: %v", err)
	}

	return nil
}

func validFilepath(file string) (string, error) {
	manifestFile := filepath.Clean(file)
	fi, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file %q does not exist", file)
		}
		return "", fmt.Errorf("could not determine if file %q exists", file)
	}
	if fi.IsDir() {
		return "", fmt.Errorf("expected %q to be a file, intead found a directory", file)
	}
	return manifestFile, nil
}

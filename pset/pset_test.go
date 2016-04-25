package pset

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGenerateFiles(t *testing.T) {
	err := os.RemoveAll("1.1.6")
	if err != nil {
		t.Log(err)
	}

	generateFiles("1.1.6.1-5")

	identical, err := contentsIdentical("1.1.6", "pregen/1.1.6")
	if err != nil {
		t.Error(err)
	}
	if !identical {
		t.Error("Contents were not identical")
	}
}

func contentsIdentical(actualDir string, expectedDir string) (bool, error) {

	expectedDirFiles, err := ioutil.ReadDir(expectedDir)
	if err != nil {
		return false, err
	}

	for _, fileInfo := range expectedDirFiles {
		matches, err := fileMatchesCounterpart(fileInfo.Name(), actualDir, expectedDir)
		if err != nil {
			return false, err
		}
		if !matches {
			return false, nil
		}
	}

	return true, nil
}

func fileMatchesCounterpart(fileName string, actualDir string, expectedDir string) (bool, error) {
	log.Printf("Checking if %s matches", fileName)
	expectedBytes, err := ioutil.ReadFile(filepath.Join(expectedDir, fileName))
	if err != nil {
		return false, err
	}
	actualBytes, err := ioutil.ReadFile(filepath.Join(actualDir, fileName))
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(expectedBytes, actualBytes), nil
}

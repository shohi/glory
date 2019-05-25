package util

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestGetWorkDir(t *testing.T) {
	dir1, _ := os.Executable()
	dir2, _ := os.Getwd()
	dir3, _ := filepath.Abs(filepath.Dir("./"))

	log.Printf("exeutable: %v", dir1)
	log.Printf("getwd: %v", dir2)
	log.Printf("filepath: %v", dir3)
}

func TestReadDir(t *testing.T) {
	dir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(dir)

	for _, f := range files {
		log.Printf("file: %v", filepath.Join(dir, f.Name()))
	}

}

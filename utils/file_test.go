package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setTest(dir string) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Errorf("failed to create test directory: '%s', error: '%s'", dir, err.Error())
	}
}

func removeTest(dir string) {
	if err := os.RemoveAll(dir); err != nil { 
		fmt.Errorf("%s", err.Error()) 
	} 
}

func compareFile(file1, file2 string) bool {
	// Check file size ...

	f1, err := os.Open(file1)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	chunkSize := 64000
	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}

func createTestFile(msg, filePath string, perm os.FileMode) {
	d1 := []byte(msg)
	err := ioutil.WriteFile(filePath, d1, perm)
	if err != nil {
		panic(err)
	}
}

func Test_Exists(t *testing.T) {
	testDir := "./test_directory"

	setTest(testDir)
	defer removeTest(testDir)

	filePath := testDir + "/hello"
	t.Run("file is not exists", func (t *testing.T) {
		assert.EqualValues(t, false, Exists(filePath))
	})

	// create file
	createTestFile("Hello test", filePath, 0755)

	t.Run("file is exists", func (t *testing.T) {
		assert.EqualValues(t, true, Exists(filePath))
	})
}

func Test_CreateIfNotExists(t *testing.T) {
	testDir := "./test_directory"

	setTest(testDir)
	defer removeTest(testDir)

	path := testDir + "/test_create_dir"

	t.Run("directory is not exists", func (t *testing.T) {
		if err := CreateIfNotExists(path, 0755); err != nil {
			panic(err)
		}
		assert.DirExists(t, path)
	})

	t.Run("directory is exists", func (t *testing.T) {
		if err := CreateIfNotExists(path, 0755); err != nil {
			panic(err)
		}
		assert.DirExists(t, path)
	})
}

func Test_Copy(t *testing.T) {
	testDir := "./test_directory"

	setTest(testDir)
	defer removeTest(testDir)

	fileName := "test_file"
	sourceDir := testDir + "/source/"
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		t.Errorf("failed to create destination directory: '%s', error: '%s'", sourceDir, err.Error())
	}

	destDir := testDir + "/destination"
	if err := os.MkdirAll(destDir, 0755); err != nil {
		t.Errorf("failed to create destination directory: '%s', error: '%s'", destDir, err.Error())
	}

	sourceFile := sourceDir + "/" + fileName
	destFile := destDir + "/" + fileName

	// create file
	createTestFile("hello test", sourceFile, 0755)

	t.Run("copy file should success", func (t *testing.T) {
		if err := Copy(sourceFile, destFile); err != nil {
			panic(err)
		}
		assert.EqualValues(t, true, compareFile(sourceFile, destFile), "Destination file is not same as the source file")
	})

	t.Run("throw error if copy directory", func (t *testing.T) {
		err := Copy(sourceDir, destDir)
		assert.Error(t, err)
	})

	t.Run("same source and destination", func (t *testing.T) {
		err := Copy(sourceFile, sourceFile)
		assert.Error(t, err)
		if err != nil {
			assert.EqualValues(t, "source and direction are duplicated", err.Error())
		}
	})
}

func Test_CopyDirectory(t *testing.T) {
	testDir := "./test_directory"

	setTest(testDir)
	defer removeTest(testDir)

	sourceDir := testDir + "/source/"
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		t.Errorf("failed to create destination directory: '%s', error: '%s'", sourceDir, err.Error())
	}

	destDir := testDir + "/destination"
	if err := os.MkdirAll(destDir, 0755); err != nil {
		t.Errorf("failed to create destination directory: '%s', error: '%s'", destDir, err.Error())
	}

	sourceFiles := []string{ sourceDir + "/test_1", sourceDir + "/test_2", sourceDir + "/test_3" }
	for i, path := range sourceFiles {
		createTestFile("hello" + strconv.Itoa(i), path, 0755)
	}

	subDir := sourceDir + "/test_sub"
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Errorf("failed to create test sub directory: '%s', error: '%s'", subDir, err.Error())
	}
	subFiles := []string{ subDir + "/test_1", subDir + "/test_2" }
	for i, path := range subFiles {
		createTestFile("hello" + strconv.Itoa(i), path, 0755)
	}
	
	t.Run("copy all file in directory", func (t *testing.T) {
		if err := CopyDirectory(sourceDir, destDir); err != nil {
			panic(err)
		}

		// check dir
		readSourceDir,_ := ioutil.ReadDir(sourceDir)
		readDestDir,_ := ioutil.ReadDir(destDir)
		assert.EqualValues(t, len(readSourceDir), len(readDestDir), "Number of source and destination is not match")
		
		// check sub dir
		readSourceDir,_ = ioutil.ReadDir(subDir)
		readDestDir,_ = ioutil.ReadDir(destDir + "/test_sub")
		assert.EqualValues(t, len(readSourceDir), len(readDestDir), "Number of source and destination is not match")
	})
}

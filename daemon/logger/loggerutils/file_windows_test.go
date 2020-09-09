package loggerutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

<<<<<<< HEAD
	"gotest.tools/assert"
=======
	"gotest.tools/v3/assert"
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
)

func TestOpenFileDelete(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", t.Name())
	assert.NilError(t, err)
	defer os.RemoveAll(tmpDir)

	f, err := openFile(filepath.Join(tmpDir, "test.txt"), os.O_CREATE|os.O_RDWR, 644)
	assert.NilError(t, err)
	defer f.Close()

	assert.NilError(t, os.RemoveAll(f.Name()))
}

func TestOpenFileRename(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", t.Name())
	assert.NilError(t, err)
	defer os.RemoveAll(tmpDir)

	f, err := openFile(filepath.Join(tmpDir, "test.txt"), os.O_CREATE|os.O_RDWR, 0644)
	assert.NilError(t, err)
	defer f.Close()

	assert.NilError(t, os.Rename(f.Name(), f.Name()+"renamed"))
}

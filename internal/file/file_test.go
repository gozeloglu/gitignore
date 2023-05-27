package file

import (
	"os"
	"path"
	"testing"
)

func TestSave(t *testing.T) {
	dir := os.TempDir()
	data := []byte(".gitignore file")
	err := Save(path.Join(dir, ".gitignore"), data)
	if err != nil {
		t.Error(err)
	}
}

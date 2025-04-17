package db_connection

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMigration(t *testing.T) {
	fail := "failed to migrate database:"
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	DbOpen()
	CloseDb()

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC
	containsFail := strings.Contains(out, fail)

	if containsFail {
		t.Errorf("No se deberia mostrar el error: %v", fail)
	}
}

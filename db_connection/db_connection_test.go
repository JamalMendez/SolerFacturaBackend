package db_connection

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDbOpen(t *testing.T) {
	DbOpen()

	if Db == nil {
		t.Error("Hubo un error al abrir la base de datos")
	}

	CloseDb()
}

func TestCloseDb(t *testing.T) {
	expected := "Database connection closed"
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
	containsExpected := strings.Contains(out, expected)

	if !containsExpected {
		t.Errorf("Se esperaba que la salida contuviera: %v", expected)
	}
}

package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	Main()

	_ = w.Close()

	res, _ := io.ReadAll(r)

	output := string(res)

	os.Stdout = stdOut

	if !strings.Contains(output, "57200.00") {
		t.Error("not correct final balance")
	}
}

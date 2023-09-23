package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage(&wg, "epsilon")

	wg.Wait()

	if msg != "epsilon" {
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"

	printMessage()

	_ = w.Close()

	res, _ := io.ReadAll(r)

	output := string(res)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon, but it's not there")
	}
}

func TestChallenge1(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	Challenge1()

	_ = w.Close()

	res, _ := io.ReadAll(r)
	output := string(res)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello Universe !!!") {
		t.Errorf("Expected to find Hello Universe !!!, but it's not there")
	}

	if !strings.Contains(output, "Hello Cosmos !!!") {
		t.Errorf("Expected to find Hello Cosmos !!!, but it's not there")
	}

	if !strings.Contains(output, "Hello World !!!") {
		t.Errorf("Expected to find Hello World !!!, but it's not there")
	}

}

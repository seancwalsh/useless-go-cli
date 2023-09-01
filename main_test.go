package main

import (
  "bytes"
  "os"
  "testing"
)

func TestUselessCli(t *testing.T) {
  // Backup original args and restore them after the test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

  // Capture the standard output
	var buf bytes.Buffer
	out = &buf
	defer func() { out = os.Stdout }()

  t.Run("NoArgs", func(t *testing.T) {
		os.Args = []string{"useless"}
		main()

		expected := "\nError: you must pass your name as an arg into this useless CLI\n\nUsage: useless <name>\n\n"
		if buf.String() != expected {
			t.Errorf("Expected: %q, got: %q", expected, buf.String())
		}
		buf.Reset()
	})

  t.Run("WithNameArg", func(t *testing.T) {
		os.Args = []string{"useless", "John"}
		main()

		expected := "Hello, this is a useless CLI tool. Your name is John\n"
		if buf.String() != expected {
			t.Errorf("Expected: %q, got: %q", expected, buf.String())
		}
		buf.Reset()
	})
}

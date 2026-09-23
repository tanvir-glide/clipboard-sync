package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"crypto/sha256"
)

func clipboard_read() (string, error) {
  var cmd *exec.Cmd

 	if runtime.GOOS == "darwin" {
		cmd = exec.Command("pbpaste")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("xclip", "-selection", "clipboard", "-o")
	} else {
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func sha256sum(str string) []byte {
  hash := sha256.Sum256([]byte(str))
  return hash[:]
}

func main() {
  text, err := clipboard_read()
  if err != nil {
    panic(err)
  }
  fmt.Println("clipboard:", text)
  fmt.Println("sha256hash:", sha256sum(text))
}

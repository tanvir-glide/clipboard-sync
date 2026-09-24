package main

import (
  "bytes"
  "fmt"
  "os/exec"
  "runtime"
  "crypto/sha256"
  "crypto/cipher"
  "encoding/base64"
  "io"
  "crypto/rand"
  "crypto/aes"
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

func clipboard_write(text string) error {
  var cmd *exec.Cmd

  if runtime.GOOS == "darwin" {
    cmd = exec.Command("pbcopy")
  } else if runtime.GOOS == "linux" {
    cmd = exec.Command("xclip", "-selection", "clipboard")
  } else {
    return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
  }

  cmd.Stdin = bytes.NewBufferString(text)
  return cmd.Run()
}


func sha256sum(str string) []byte {
  hash := sha256.Sum256([]byte(str))
  return hash[:]
}


func encrypt_aes256(plainText, key []byte) (string, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return "", err
  }

  gcm, err := cipher.NewGCM(block)
  if err != nil {
    return "", err
  }

  nonce := make([]byte, gcm.NonceSize())
  if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
          return "", err
  }

  encryptedText := gcm.Seal(nonce, nonce, plainText, nil)
  return base64.StdEncoding.EncodeToString(encryptedText), nil
}


func main() {
  text, err := clipboard_read()
  if err != nil {
    panic(err)
  }
  fmt.Println("clipboard:", text)
  fmt.Println("sha256hash:", sha256sum(text))
}

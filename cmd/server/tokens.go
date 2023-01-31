package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrInvalidToken = errors.New("invalid token")

var tokensFilename = filepath.Join(os.TempDir(), "tokens.txt")

func verifyToken(token string) error {
	content, err := os.ReadFile(tokensFilename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrInvalidToken
		}

		return fmt.Errorf("failed to read tokens file: %w", err)
	}

	lines := bytes.Split(content, []byte{'\n'})

	for _, line := range lines {
		if bytes.Equal(line, []byte(token)) {
			return nil
		}
	}

	return ErrInvalidToken
}

func storeToken(token string) error {
	file, err := os.OpenFile(tokensFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open tokens file: %w", err)
	}

	defer file.Close()

	if _, err := file.Write(append([]byte(token), '\n')); err != nil {
		return fmt.Errorf("failed to write token to file: %w", err)
	}

	return nil
}

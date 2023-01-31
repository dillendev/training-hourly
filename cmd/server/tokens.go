package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

var ErrInvalidToken = errors.New("invalid token")

func verifyToken(token string) error {
	content, err := os.ReadFile("tokens.txt")
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
	file, err := os.OpenFile("tokens.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open tokens file: %w", err)
	}

	defer file.Close()

	if _, err := file.Write(append([]byte(token), '\n')); err != nil {
		return fmt.Errorf("failed to write token to file: %w", err)
	}

	return nil
}

package io

import (
	"crypto/rand"
	"os"
	"path/filepath"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// GetTempFile creates a new temporary file
func GetTempFile() string {
	dir := os.TempDir()
	path := filepath.Join(dir, getNonce(10))
	return path
}

func getNonce(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for _, x := range getRandomBytes(n) {
		sb.WriteByte(letters[x%byte(len(letters))])
	}
	return sb.String()
}

func getRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

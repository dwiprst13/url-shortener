package main

import (
	"crypto/rand"
	"errors"
	"net/url"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func generateCode(n int) (string, error) {
	if n <= 0 {
		n = 6
	}
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	for i := range b {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}

	return string(b), nil
}

func normalizedURL(raw string) (string, error) {
	if raw == ""{
		return "", errors.New("empty url")
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	if u.Host == "" {
		return "", errors.New("invalid url: missing host")
	}
	return u.String(), nil
}
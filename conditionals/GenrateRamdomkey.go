package main

import "crypto/rand"

const latters ="abcdefghijklmnopqrstuvwxyz123456789"
func GenrateRamdomkey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, b := range bytes{
		bytes[i] = latters[b%byte(len(latters))]
	}
	return string(bytes),nil
}
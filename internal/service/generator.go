package service

import (
	"log"
	"math/rand"
)

type shortCodeGenerator struct {
	length int
	chars  string
}

func NewShortCodeGenerator(length int) URLGenerator {
	return &shortCodeGenerator{
		length: length,
		chars:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
}
func (g *shortCodeGenerator) Generate() string {

	result := make([]byte, g.length)

	for i := range result {
		result[i] = g.chars[rand.Intn(len(g.chars))]
	}
	log.Printf("shortCode generated")
	return string(result)
}

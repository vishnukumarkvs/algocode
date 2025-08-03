package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	mrand "math/rand"
	"time"
)

var levels = []string{"INFO", "WARNING", "ERROR", "TRACE"}

func RandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Below will return random string but lot of random symbols and spaces
	// return string(b), nil

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func LogGenerator() string {
	s, err := RandomString(16)
	if err != nil {
		log.Fatal(err)
	}

	level := levels[mrand.Intn(len(levels))]

	l := fmt.Sprintf("%s %s %s", time.Now(), level, s)

	return l
}

func main() {
	t := time.NewTicker(1 * time.Millisecond)
	defer t.Stop()

	for range t.C {
		l := LogGenerator()
		fmt.Println(l)
	}
}

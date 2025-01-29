package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error
	if name == "" {
			return "", errors.New("empty name")
	}
	
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail allmighty %v!, well met!",
	}

	return formats[rand.Intn(len(formats))]
}
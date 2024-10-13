package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	// Initializing Sentry with DSN from environment variables
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		log.Fatal("SENTRY_DSN environment variable is required")
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	// Initialization of the random number generator
	rand.Seed(time.Now().UnixNano())

	// Infinite loop for generating and sending different types of errors
	for {
		var err error
		switch rand.Intn(8) {
		case 0:
			err = &newCustomError{message: fmt.Sprintf("new custom error: %d", rand.Intn(1000))}
		case 1:
			err = &anotherNewCustomError{message: fmt.Sprintf("another new custom error: %d", rand.Intn(1000))}
		case 2:
			err = &yetAnotherNewCustomError{message: fmt.Sprintf("yet another new custom error: %d", rand.Intn(1000))}
		case 3:
			err = &complexNewCustomError{code: rand.Intn(100), message: fmt.Sprintf("complex new custom error: %d", rand.Intn(1000))}
		case 4:
			err = &yetAnotherComplexCustomError{code: rand.Intn(100), message: fmt.Sprintf("yet another complex custom error: %d", rand.Intn(1000))}
		case 5:
			err = &evenMoreComplexCustomError{code: rand.Intn(100), message: fmt.Sprintf("even more complex custom error: %d", rand.Intn(1000))}
		case 6:
			err = &simpleCustomError{message: fmt.Sprintf("simple custom error: %d", rand.Intn(1000))}
		case 7:
			err = &anotherSimpleCustomError{message: fmt.Sprintf("another simple custom error: %d", rand.Intn(1000))}
		}

		sentry.CaptureException(err)
		log.Printf("Sent error: %v to Sentry", err)
		time.Sleep(100 * time.Millisecond) // Delay to simulate real load
	}
}

type newCustomError struct {
	message string
}

func (e *newCustomError) Error() string {
	return e.message
}

type anotherNewCustomError struct {
	message string
}

func (e *anotherNewCustomError) Error() string {
	return e.message
}

type yetAnotherNewCustomError struct {
	message string
}

func (e *yetAnotherNewCustomError) Error() string {
	return e.message
}

type complexNewCustomError struct {
	code    int
	message string
}

func (e *complexNewCustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.code, e.message)
}

type yetAnotherComplexCustomError struct {
	code    int
	message string
}

func (e *yetAnotherComplexCustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.code, e.message)
}

type evenMoreComplexCustomError struct {
	code    int
	message string
}

func (e *evenMoreComplexCustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.code, e.message)
}

type simpleCustomError struct {
	message string
}

func (e *simpleCustomError) Error() string {
	return e.message
}

type anotherSimpleCustomError struct {
	message string
}

func (e *anotherSimpleCustomError) Error() string {
	return e.message
}

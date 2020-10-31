package main

import (
	limit "github.com/lu-moreira/go-concurrency-exercises/0-limit-crawler"
	consumer "github.com/lu-moreira/go-concurrency-exercises/1-producer-consumer"
	limitservice "github.com/lu-moreira/go-concurrency-exercises/3-limit-service-time"
	sigint "github.com/lu-moreira/go-concurrency-exercises/4-graceful-sigint"
	sessioncleaner "github.com/lu-moreira/go-concurrency-exercises/5-session-cleaner"
)

func main() {
	ex05()
}

// exercise 00
func ex00() {
	limit.Start()
}

// exercise 01
func ex01() {
	consumer.Start()
}

func ex03() {
	limitservice.Start()
}

func ex04() {
	sigint.Start()
}

func ex05() {
	sessioncleaner.Start()
}

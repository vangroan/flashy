package main

type contextKey int

const (
	operationIDKey contextKey = iota
	loggerKey      contextKey = iota
)

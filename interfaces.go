package main

type commitMessageProvider interface {
	Provide() commitMessageQueue
}

type commitMessageQueue func() *string

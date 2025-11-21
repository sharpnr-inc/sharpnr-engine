package sere

import "fmt"

type SERE struct {
	Events chan string
}

func NewSERE() *SERE {
	return &SERE{
		Events: make(chan string, 20),
	}
}

func (s *SERE) Start() {
	fmt.Println("🚀 SERE: Sharpnr Event Relay Engine is starting...")

	go func() {
		fmt.Println("✅ SERE: Ready! Listening for events...")

		for evt := range s.Events {
			fmt.Println("📡 SERE RECEIVED EVENT →", evt)
		}
	}()
}

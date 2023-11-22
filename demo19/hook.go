package main

import "fmt"

type HookFunc func(interface{})

var hooks []HookFunc

func registerHook(h HookFunc) {
	hooks = append(hooks, h)
}

func RunHooks(event interface{}) {
	for _, hook := range hooks {
		hook(event)
	}
}

func main() {
	registerHook(func(event interface{}) {
		fmt.Printf("hook1: %v\n", event)
	})

	registerHook(func(event interface{}) {
		fmt.Printf("hook2: %v\n", event)
	})

	RunHooks("testEvent")
}

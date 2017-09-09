/*
go-publisher generates type-safe publisher/subscriber implementation in Go.

Usage:
	go-publisher [flags] [event ...]

The flags are:
	-o
		Output filename.
	-p
		Package name.
	-t
		Publisher type name.
	-v
		Print version info.
	-tags
		Go build tags.

Debugging support:
	-no-gofmt
		Generate Go source code without formatting.

Examples
	package main

	import (
		"context"
		"log"
	)

	type subscriber struct{}

	func (self *subscriber) OnEventA(ctx context.Context) error {
		log.Printf("On%s with no args", ExampleEventA)
		return nil
	}

	func (self *subscriber) OnEventB(ctx context.Context, arg1 string, arg2 int) error {
		log.Printf("On%s with %v", ExampleEventB, []interface{}{arg1, arg2})
		return nil
	}

	//go:generate go-publisher -o example_publisher.go -t Example EventA 'EventB|arg1:string,arg2:int'
	func ExampleUsage() {
		subscription := NewExampleSubscription()
		subscription.Subscribe(&subscriber{})
		log.Printf("publish %s", ExampleEventA)
		if err := subscription.PublishEventA(context.Background()); err != nil {
			log.Printf("Publish%s: %s", ExampleEventA, err)
		}
		log.Printf("publish %s", ExampleEventB)
		if err := subscription.PublishEventB(context.Background(), "example", 9999); err != nil {
			log.Printf("PublishEventB: %s", err)
		}
	}
*/
package main

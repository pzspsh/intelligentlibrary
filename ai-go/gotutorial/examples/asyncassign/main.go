package main

import (
	"gotutorial/examples/asyncassign/input"
	"gotutorial/examples/asyncassign/output"
)

func main() {
	go func() {
		input.Input()
	}()
	output.Output()
}

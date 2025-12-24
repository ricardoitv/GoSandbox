package main

import (
	//fstreev1 "github.com/rcdmrl/go-sandbox/fstree/v1"
	fstreev2 "github.com/rcdmrl/go-sandbox/fstree/v2"
)

func main() {
	// pd1 := fstreev1.NewParallelDir("/Users/ricaamar/Documents/")
	// pd1.Run()

	pd2 := fstreev2.NewParallelDir("/Users/ricaamar/Documents/")
	pd2.Run()
}

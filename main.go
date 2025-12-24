package main

import (
	"github.com/rcdmrl/go-sandbox/fstree"
)

func main() {
	pd := fstree.NewParallelDir("/Users/ricaamar/Documents/")
	pd.Run()
}

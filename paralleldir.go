package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type ParallelDir struct {
	root *treeNode
}

type treeNode struct {
	name     string
	children []*treeNode
}

func NewParallelDir(baseDir string) *ParallelDir {
	return &ParallelDir{
		root: &treeNode{
			name:     baseDir,
			children: make([]*treeNode, 0),
		},
	}
}

func (pd *ParallelDir) baseDir() string {
	return pd.root.name
}

func (pd *ParallelDir) Run() {
	var wg sync.WaitGroup
	fmt.Println("Starting on", pd.baseDir())
	listDirsRecursively(pd.root, &wg)
	wg.Wait()
	fmt.Println(pd.root)
}

func listDirsRecursively(node *treeNode, wg *sync.WaitGroup) {
	files, err := os.ReadDir(node.name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			childNode := &treeNode{
				name:     file.Name(),
				children: make([]*treeNode, 0),
			}
			node.children = append(node.children, childNode)
			wg.Go(func() {
				listDirsRecursively(childNode, wg)
			})
		}
	}
}

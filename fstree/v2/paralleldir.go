package v2

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type treeNode struct {
	Name     string      `json:"name"`
	Children []*treeNode `json:"children"`
}

type nodeUpdate struct {
	Parent *treeNode
	Child  *treeNode
}

type ParallelDir struct {
	Root           *treeNode `json:"root"`
	nodeUpdateChan chan *nodeUpdate
}

func NewParallelDir(baseDir string) *ParallelDir {
	return &ParallelDir{
		Root: &treeNode{
			Name:     baseDir,
			Children: make([]*treeNode, 0),
		},
		nodeUpdateChan: make(chan *nodeUpdate),
	}
}

func (pd *ParallelDir) baseDir() string {
	return pd.Root.Name
}

// appendChildren consumes update requests sent to the ParallelDir.nodeUpdateChan and performs the update
func (pd *ParallelDir) appendChildren() {
	for updateRequest := range pd.nodeUpdateChan {
		updateRequest.Parent.Children = append(updateRequest.Parent.Children, updateRequest.Child)
	}
}

func (pd *ParallelDir) Run() {
	var wg sync.WaitGroup
	fmt.Println("Starting on", pd.baseDir())
	pd.Root.listDirsRecursively(&wg, pd.nodeUpdateChan)
	go pd.appendChildren()                // single goroutine to read from the update channel and update the children
	wg.Wait()                             // waiting for all the parallel traversal goroutines
	close(pd.nodeUpdateChan)              // hopefully this will "free" the goroutine handling the children update TODO: do I need a waitgroup for this as well?
	fmt.Println(pd.Root.toString())       // displaying tree as text
	fmt.Println(strings.Repeat("-", 100)) // separator
	fmt.Println(pd.toJson())              // displaying tree as JSON
}

func (pd *ParallelDir) toJson() string {
	bytes, err := json.Marshal(pd.Root)
	if err != nil {
		log.Fatal("Error converting final tree to JSON:", err.Error())
	}
	return string(bytes)
}

func (t *treeNode) toString() string {
	return t.toStringWithIndent("")
}

func (t *treeNode) toStringWithIndent(indent string) string {
	var result strings.Builder

	result.WriteString(indent + t.Name + "\n")
	for _, child := range t.Children {
		result.WriteString(child.toStringWithIndent(indent + "\t"))
	}

	return result.String()
}

func (t *treeNode) listDirsRecursively(wg *sync.WaitGroup, ch chan<- *nodeUpdate) {
	files, err := os.ReadDir(t.Name)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			childNode := &treeNode{
				Name:     filepath.Join(t.Name, file.Name()),
				Children: make([]*treeNode, 0),
			}
			//t.safeAppendChild(childNode)
			update := nodeUpdate{
				Parent: t,
				Child:  childNode,
			}
			ch <- &update

			wg.Go(func() {
				childNode.listDirsRecursively(wg, ch)
			})
		}
	}
}

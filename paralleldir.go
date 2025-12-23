package main

import (
	"fmt"
	"log"
	"os"
)

type ParallelDir struct {
	baseDir string
}

func NewParallelDir(baseDir string) *ParallelDir {
	return &ParallelDir{baseDir: baseDir}
}

func (pd *ParallelDir) Run() {
	fmt.Println("Starting on", pd.baseDir)
	for _, dir := range pd.listDirs() {
		fmt.Println("-", dir)
	}
}

// listDirs Returns a list of directories inside of the current base directory
func (pd *ParallelDir) listDirs() []string {
	files, err := os.ReadDir(pd.baseDir)
	if err != nil {
		log.Fatal(err)
	}

	dirs := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	return dirs
}

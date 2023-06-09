package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	reader, err := getInput()
	if err != nil {
		fmt.Println("getInput", err.Error())
		os.Exit(1)
	}
	project, err := ReadPom(reader)
	if err != nil {
		fmt.Println("Unmarshal", err.Error())
		os.Exit(2)
	}
	deps := project.Dependencies
	if deps != nil {
		for _, dependency := range deps.Dependency {
			fmt.Println(
				"group:", dependency.GroupId,
				"name:", dependency.ArtifactId,
				"version:", dependency.Version,
				"classifier:", dependency.Classifier,
				"scope:", dependency.Scope,
				"optional:", dependency.Optional,
			)
		}
	}
	props := project.Properties
	if props == nil {
		os.Exit(0)
	}
}

func getInput() (io.ReadCloser, error) {
	length := len(os.Args)
	if length == 1 {
		return stdin(), nil
	} else if length == 2 {
		filename := os.Args[1]
		return os.Open(filename)
	}
	return nil, fmt.Errorf("files can be accept only 1")
}

type StdIn struct {
	input *os.File
}

func (s *StdIn) Close() error {
	return nil
}

func (s *StdIn) Read(p []byte) (n int, err error) {
	return s.input.Read(p)
}

func stdin() io.ReadCloser {
	return &StdIn{input: os.Stdin}
}

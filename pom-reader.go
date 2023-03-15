package main

import (
	"encoding/xml"
	"fmt"
	"io"
)

func ReadPom(reader io.Reader) (Project, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("ReadPom.io-ReadAll: %w", err)
	}
	var project Project
	err = xml.Unmarshal(bytes, &project)
	if err != nil {
		return nil, fmt.Errorf("ReadPom.xml-Unmarshal: %w", err)
	}
	return project, nil
}

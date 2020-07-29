/*
Things to think about:
--
- We need to check if the filename is valid
- What kind of permissions do we need to set on the file
- Will need to create go files in the project
*/

package project

import (
	"fmt"
	"os"
)

type Project struct {
	Name string
}

func New(name string) *Project {
	return &Project{
		Name: name,
	}
}

func (p *Project) Create() error {
	err := os.Mkdir(p.Name, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Something went wrong while creating the file: %v", err)
	}
	fmt.Printf("Project %v was created successfully.", p.Name)
	return nil
}

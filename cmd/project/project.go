/*
Things to think about:
--
- We need to check if the filename is valid
- What kind of permissions do we need to set on the file
- Will need to create go files in the project
-- We need to create test for project package
*/

package project

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type Project struct {
	Name string
}

type MainFile struct {
	Project string
}

func New(name string) *Project {
	return &Project{
		Name: name,
	}
}

// returns file with hello world scaffolded.
func createMainFile(projectName string) {
	// How do we create this file in the directory? This will create file with 0666 permissions
	f, err := os.Create("main.go")
	if err != nil {
		fmt.Errorf("Something went wrong.")
	}

	data := MainFile{
		Project: projectName,
	}

	time.Sleep(2 * time.Second)

	tmpl := template.Must(template.ParseFiles("../templates/main.tmpl"))
	tmpl.Execute(f, data)
}

func (p *Project) Create() error {
	err := os.Mkdir(p.Name, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Something went wrong while creating the file: %v", err)
	}
	fmt.Printf("Project %v was created successfully.", p.Name)

	err = os.Chdir(p.Name)
	if err != nil {
		return fmt.Errorf("Something went wrong while trying to change directories: %v", err)
	}

	createMainFile(p.Name)

	return nil
}

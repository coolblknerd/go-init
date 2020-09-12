package api

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Api struct {
	PackageName string
	Handles     []Handle
}

type Handle struct {
	Name string
}

func New(name string, handles []string) *Api {
	var h []Handle
	for _, v := range handles {
		h = append(h, Handle{v})
	}

	return &Api{
		PackageName: name,
		Handles:     h,
	}
}

func createAPIFile(filename string, packageName string, handlers []Handle) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln("Sorry something went wrong")
	}

	api := &Api{
		PackageName: packageName,
		Handles:     handlers,
	}

	tmpl := template.Must(template.ParseFiles("../templates/api.tmpl"))
	tmpl.Execute(f, api)
}

func (a *Api) Create() error {
	err := os.Mkdir(a.PackageName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Something went wrong while creating the file: %v", err)
	}
	fmt.Printf("Project %v was created successfully.", a.PackageName)

	err = os.Chdir(a.PackageName)
	if err != nil {
		return fmt.Errorf("Something went wrong while trying to change directories: %v", err)
	}

	filename := fmt.Sprintf("%v.go", a.PackageName)

	createAPIFile(filename, a.PackageName, a.Handles)

	return nil
}

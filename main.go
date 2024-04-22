package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var path string = "./todos.json"

type Todo struct {
	Id     uint16
	Name   string
	Desc   string
	Status string
}

func main() {
	add := flag.NewFlagSet("add", flag.ExitOnError)
	name := add.String("n", "", "Name of the todo item (shorthand)")
	desc := add.String("d", "", "Description of the todo item (shorthand)")
	flag.Parse()

	var command string = flag.Arg(0)

	switch command {
		
	case "add":
		add.Parse(flag.Args()[1:])
		fmt.Println(*name, *desc)
		todo := Todo{
			Id:     1,
			Name:   *name,
			Desc:   *desc,
			Status: "uncomplete",
		}
		json, _ := json.Marshal(todo)
		os.WriteFile(path, json, 0777)
	case "get":
		var todos []Todo
		content, _ := os.ReadFile(path)
		json.Unmarshal(content, &todos)

	}
}

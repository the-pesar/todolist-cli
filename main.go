package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var path string = "./todos.json"

type Todo struct {
	Id     uint16
	Name   string
	Desc   string
	Status string
}

var todos = []Todo{}

func main() {
	add := flag.NewFlagSet("add", flag.ExitOnError)
	name := add.String("n", "", "Name of the todo item (shorthand)")
	desc := add.String("d", "", "Description of the todo item (shorthand)")

	list := flag.NewFlagSet("list", flag.ExitOnError)
	flag.Parse()

	var command string = flag.Arg(0)

	content, _ := os.ReadFile(path)
	json.Unmarshal(content, &todos)

	switch command {

	// add subcommand function
	case "add":
		add.Parse(flag.Args()[1:])
		fmt.Println(*name, *desc)

		todo := Todo{
			Id:     uint16(len(todos)) + 1,
			Name:   *name,
			Desc:   *desc,
			Status: "uncomplete",
		}

		todos = append(todos, todo)
		json, _ := json.Marshal(todos)
		os.WriteFile(path, json, 0777)

	// list subcommand function
	case "list":
		list.Parse(flag.Args()[1:])

		idStr := list.Arg(0)
		if idStr == "" {
			fmt.Println(todos)
			return
		}

		id, _ := strconv.ParseUint(idStr, 10, 16)

		for _, v := range todos {
			if v.Id == uint16(id) {
				fmt.Println(v)
			}
		}
	}
}

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
	namePtr := flag.String("n", "", "Name of the item")
	descPtr := flag.String("d", "", "Description of the item")

	add := flag.NewFlagSet("add", flag.ExitOnError)

	add.StringVar(namePtr, "n", "", "Name of the todo item (shorthand)")
	add.StringVar(descPtr, "d", "", "Description of the todo item (shorthand)")

	flag.Parse()
	add.Parse(flag.Args()[1:])

	var command string = flag.Arg(0)

	switch command {
	case "add":
		todo := Todo{
			Id:     1,
			Name:   *namePtr,
			Desc:   *descPtr,
			Status: "uncomplete",
		}
		fmt.Println(todo)
		json, _ := json.Marshal(todo)
		os.WriteFile(path, json, os.ModePerm)
	}

}

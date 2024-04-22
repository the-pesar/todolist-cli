package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path string = "./todos.json"

type Todo struct {
	Id     uint16
	Name   string
	Desc   string
	Status string
}

var todos = []Todo{}

/*
	+++++++++++++++description
  	+ name 6      +
	+++++++++++++++
*/

func tablize(todos []Todo) {
	var idLength, nameLength, descLength, statusLength = 2, 4, 11, 6

	for _, v := range todos {
		if len(v.Name) > int(nameLength) {
			nameLength = len(v.Name)
		}
		if len(v.Desc) > descLength {
			descLength = len(v.Desc)
		}
		if len(fmt.Sprint(v.Id)) > idLength {
			idLength = len(fmt.Sprint(v.Id))
		}
	}

	totalLength := idLength + nameLength + descLength + statusLength + 13

	var header = fmt.Sprintf("\n%v\n+ \033[1mid\033[0m %v+ \033[1mname\033[0m %v+ \033[1mdescription\033[0m %v+ \033[1mstatus\033[0m +\n%v\n",
		strings.Repeat("+", totalLength),
		strings.Repeat(" ", idLength - 2),
		strings.Repeat(" ", nameLength-4),
		strings.Repeat(" ", descLength-11),
		strings.Repeat("+", totalLength))

	var body string

	for _, v := range todos {
		body = body + fmt.Sprintf("+ %v%v + %v%v + %v%v + %v%v +\n%v\n",
			v.Id,
			strings.Repeat(" ", idLength - len(fmt.Sprint(v.Id))),
			v.Name,
			strings.Repeat(" ", nameLength-len(v.Name)),
			v.Desc,
			strings.Repeat(" ", descLength-len(v.Desc)),
			v.Status,
			strings.Repeat(" ", statusLength - len(v.Status)),
			strings.Repeat("+", totalLength))
	}

	fmt.Println(header + body)
}

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
			tablize(todos)
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

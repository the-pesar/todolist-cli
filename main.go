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

func tablize(todos []Todo) {
	// length of header columns name
	var idLength, nameLength, descLength, statusLength = 2, 4, 11, 6

	// increase size of column when content is long
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
		if len(v.Status) > statusLength {
			statusLength = len(v.Status)
		}
	}

	// divider row size
	totalLength := idLength + nameLength + descLength + statusLength + 13

	// creates header
	var header = fmt.Sprintf("\n%v\n+ \033[1mid\033[0m %v+ \033[1mname\033[0m %v+ \033[1mdescription\033[0m %v+ \033[1mstatus\033[0m %v+\n%v\n",
		strings.Repeat("+", totalLength),
		strings.Repeat(" ", idLength-2),
		strings.Repeat(" ", nameLength-4),
		strings.Repeat(" ", descLength-11),
		strings.Repeat(" ", statusLength-6),
		strings.Repeat("+", totalLength))

	var body string
	// adds empty row and shows not found message
	if len(todos) == 0 {
		body = fmt.Sprintf("+%vThere are no todos%v+\n%v\n",
			strings.Repeat(" ", (totalLength-19)/2),
			strings.Repeat(" ", (totalLength-19)/2),
			strings.Repeat("+", totalLength),
		)
		fmt.Println(header + body)

		return
	}

	// create body row by row
	for _, v := range todos {
		body = body + fmt.Sprintf("+ %v%v + %v%v + %v%v + %v%v +\n%v\n",
			v.Id,
			strings.Repeat(" ", idLength-len(fmt.Sprint(v.Id))),
			v.Name,
			strings.Repeat(" ", nameLength-len(v.Name)),
			v.Desc,
			strings.Repeat(" ", descLength-len(v.Desc)),
			v.Status,
			strings.Repeat(" ", statusLength-len(v.Status)),
			strings.Repeat("+", totalLength))
	}

	fmt.Println(header + body)
}

func main() {
	add := flag.NewFlagSet("add", flag.ExitOnError)
	name := add.String("n", "", "Name of the todo item (shorthand)")
	desc := add.String("d", "", "Description of the todo item (shorthand)")

	list := flag.NewFlagSet("list", flag.ExitOnError)
	remove := flag.NewFlagSet("remove", flag.ExitOnError)
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

		id, _ := strconv.ParseUint(list.Arg(0), 10, 16)
		if id == 0 {
			tablize(todos)
			return
		}

		for _, v := range todos {
			if v.Id == uint16(id) {
				tablize([]Todo{v})
			}
		}

	// remove subcommand function
	case "remove":
		remove.Parse(flag.Args()[1:])

		id, _ := strconv.ParseUint(remove.Arg(0), 10, 16)
		if id == 0 {
			fmt.Println("please give valid id to remove!")
			return
		}

		var newTodos []Todo
		for _, v := range todos {
			if v.Id != uint16(id) {
				newTodos = append(newTodos, v)
			}
		}

		todos = newTodos

		content, _ := json.Marshal(todos)

		os.WriteFile(path, content, 0777)
		fmt.Println("Todo was successfully deleted.")
	}
}

package main

import (
	"github.com/jacky-htg/shonet-frontend/libraries"
	"log"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]

	if arguments[0] == "help" {
		help := "\n This is for input database on SQL to elasticsearch database\n"+
		        " For argument you can include <name>:<table name> example name:articles\n"+
			    " Just for this time you just inputing 3 tables: 'articles', 'products`, and 'users'\n"+
		        " ..."
		log.Printf(help)
		return
	}

	var tables map[string]string
	tables = map[string]string{}

	for _, val := range arguments {
		instructions := strings.Split(val, ":")
		if 1 >= len(instructions) {
			log.Println("\n please insert valid argumen <name>:<value> example => table:tablename for more type : help")
			return
		}
		tables[instructions[0]] = instructions[1]
	}

	log.Printf("\n bulking table : %s..\n please wait..\n\n", tables["name"])
	status, err := libraries.BulkingAllDataFromSQL(tables)
	if !status {
		if err != nil {
			panic(err)
			return
		}
	}

	log.Printf("Success to insert SQL data %s to elastic search..", tables["name"])
}
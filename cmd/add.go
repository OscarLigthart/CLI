/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	// "fmt"
	"github.com/spf13/cobra"

	"encoding/json"
	"io/ioutil"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Command for adding a todo to the list",
	Long: `The todo will be added to the total list, based on a priority level	`,
	Run: func(cmd *cobra.Command, args []string) {
		addTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Data struct {
	Todos 		[]Todo
}

type Todo struct {
	Text		string
	Priority 	int
}

func addTodo(args []string) {

	// initialize data
	var data Data;

	// optionally load a json
	if _, err := os.Stat("todos.json"); os.IsNotExist(err) {
		data = Data{
		}
	} else {
		data = loadJson("todos.json")
	}
	
	// create new todo
	todo := Todo{
		Text: 		args[0],
		Priority:	len(data.Todos) + 1,
	}

	// append to the data struct
	data.Todos = append(data.Todos, todo)

	// rewrite the json
	file, _ := json.MarshalIndent(data, "", " ") 
	_ = ioutil.WriteFile("todos.json", file, 0644)
}

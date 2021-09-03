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
	"fmt"
	"os"
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

	// optionally load a json

	// Open our jsonFile
	jsonFile, err := os.Open("todos.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	
	b, err := ioutil.ReadAll(jsonFile)

	var data Data

	json.Unmarshal(b, &data)

	// for i := 0; i < len(data.Todos); i++ {
	// 	fmt.Println("Text: " + data.Todos[i].Text)
	// 	fmt.Println("Priority: " + strconv.Itoa(data.Todos[i].Priority))
	// }

	// create new todo
	todo := Todo{
		Text: 		args[0],
		Priority:	1,
	}

	// load the json file


	// 
	data.Todos = append(data.Todos, todo)
	
	// append the new todo datafile
	// data = Data{
    //     Todos:	   []Todo{todo},
    // }

	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("todos.json", file, 0644)
}
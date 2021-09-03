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
	"strconv"

	"github.com/spf13/cobra"

	"encoding/json"
	"io/ioutil"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		rmTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func remove(slice []Todo, s int) []Todo {
    return append(slice[:s], slice[s+1:]...)
}

func rmTodo(args []string){

	// initialize data
	var data Data;

	// optionally load a json
	if _, err := os.Stat("todos.json"); !os.IsNotExist(err) {
		data = loadJson("todos.json")
	} else {
		fmt.Println("There are no todos currently!")
	}

	// get the index to remove
	index, _ := strconv.Atoi(args[0]) - 1

	fmt.Println(index)

	// remove according to the integer
	data.Todos = remove(data.Todos, index)

	// remove the integer
	fmt.Println("Removed todo")

	// rewrite the json
	file, _ := json.MarshalIndent(data, "", " ") 
	_ = ioutil.WriteFile("todos.json", file, 0644)
}
package main

import (
	"fmt"

	"github.com/youaresofunny/gpplugin/types"
)

//Plugin is exported object, the core works with it
var Plugin *types.Plugin

//Init function runs when plugin connected to core, you can describe ur hooks functions here
func Init() {
	//Initialize your plugin with name
	Plugin = types.New("test")
	fmt.Printf("PLUGIN: %v \n", Plugin)
	//Add action to the loop
	Plugin.AddAction("init", initialize)
}

func initialize() {
	fmt.Println("INIT PLUGIN")
}

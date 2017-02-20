package main

import (
	"plugin"

	"github.com/youaresofunny/plugintest/types"
)

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Init")
	if err != nil {
		panic(err)
	}

	f.(func())()

	plugin, err := p.Lookup("Plugin")
	if err != nil {
		panic(err)
	}

	pl := *plugin.(**types.Plugin)

	pl.DoAction("init")

	/*	var data interface{} = "great"

		if res, ok := data.(int); ok {
			fmt.Println("[is an int] value =>", res)
		} else {
			fmt.Println("[not an int] value =>", data)
			//prints: [not an int] value => great (as expected)
		}*/
}

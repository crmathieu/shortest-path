package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonGraph struct {
	Nodes []jsonNode `json:"nodes"`
}

type jsonNode struct {
	Name      string          `json:"name"`
	Id        int             `json:"id"`
	Neighbors []jsonNeighbors `json:"neighbors"`
}

type jsonNeighbors struct {
	Id     int `json:"id"`
	Weight int `json:"weight"`
}

func ReadJsonGraph(filename string) *jsonGraph {

	var jgraph jsonGraph
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&jgraph)
	return &jgraph
}

package main

import (
	"os"
	"fmt"
	"encoding/json"
)
type jsonGraph struct {
	Nodes   	[]jsonNode 		`json:"nodes"`
}

type jsonNode struct {
	Name 		string			`json:"name"`
	Id			int				`json:"id"`
	Neighbors 	[]jsonNeighbors	`json:"neighbors"`
}

type jsonNeighbors struct {
	Id 			int 			`json:"id"`
	weight		int				`json:"weight"`
}

/*
            "name": "Fullstack",
            "id": 0,
            "neighbors": [
                {
                    "id": 1,
                    "weight": 7
                },
                {
                    "id": 2,
                    "weight": 2
                },
                {
                    "id": 3,
                    "weight": 2
                }
            ]
*/

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
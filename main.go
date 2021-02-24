package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Enemy Defines Enemy Info
type Enemy struct {
	Name  string
	Hp    string
	Param int
}

// KHBR Defines KHBR File Spec
type KHBR struct {
	Game       string
	Enemyswaps map[string]Enemy
}

func main() {

	args := os.Args[1:]
	var fn string = args[0]

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var spec KHBR

	err = json.Unmarshal(data, &spec)
	if err != nil {
		log.Fatal("error:", err)
	}

	// Check using a supported game
	if spec.Game != "kh2" {
		log.Fatal("error - game not supported: ", spec.Game)
	}
	fmt.Printf("Randomizing: %s", spec.Game)
}

package main

import (
	"fmt"
)

func main() {
	// composite literal: permite agrupar valores del mismo tipo.
	// []tipo{}
	g := struct {
		nombre    string
		raza      string
		edad      int
		colorOjos string
	}{
		nombre:    "Nana",
		raza:      "Simesa",
		edad:      10,
		colorOjos: "Azules",
	}

	fmt.Println(g)
}
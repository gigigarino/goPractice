package main

import (
	"fmt"
	"errors"
)

func boom ()error {
	return errors.New("alto error")
}

func main() {
	err := boom()

	if err != nil {
		fmt.Println("ocurrrio un error", err)
		return
	}
	fmt.Println("ANCHORS AWAY")
}
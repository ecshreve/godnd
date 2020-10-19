package api

import "fmt"

//go:generate genny

type APIReference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func Hello() {
	fmt.Println("hello")
}

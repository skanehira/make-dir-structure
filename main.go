package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Object struct {
	Path string `json:"path"`
}

type Item struct {
	Name     string `json:"name"`
	Children Items  `json:"children,omitempty"`
}

type Items []*Item

func (is *Items) Merge(path string, rest []string) {
	for _, item := range *is {
		if path == item.Name {
			item.Children.Merge(rest[0], rest[1:])
			return
		}
	}
	item := &Item{Name: path}
	if len(rest) != 0 {
		item.Children.Merge(rest[0], rest[1:])
	}
	*is = append(*is, item)
}

func makeItems(obj string) Items {
	var objects []Object
	if err := json.Unmarshal([]byte(obj), &objects); err != nil {
		log.Fatal(err)
	}

	var items Items

	for _, obj := range objects {
		paths := strings.Split(obj.Path, "/")
		items.Merge(paths[0], paths[1:])
	}

	return items
}

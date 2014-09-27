package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/hello", func() string {
		return "Hello, World!"
	})
	m.Patch("/", func() string {
		// update something
		return "Hello, Patch!"
	})

	m.Post("/", func() string {
		// create something
		return "Hello, Post!"
	})

	m.Put("/", func() string {
		// replace something
		return "Hello, Put!"
	})

	m.Delete("/", func() string {
		// destroy something
		return "Hello, Delete!"
	})

	m.Options("/", func() string {
		// http options
		return "Hello, Option!"
	})

	m.NotFound(func() string {
		// handle 404
		return "我©, GONE!"
	})

	// m.Run()
}

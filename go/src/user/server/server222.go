package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/hello/(?P<name>[a-zA-Z]+)", func(params martini.Params) string {
		return "Helloddd, " + params["name"]
	})
	m.Patch("/", func() string {
		// update something
		return "Hello, Patch!"
	})

	m.Post("/", func(res http.ResponseWriter, req *http.Request) (int, string) {
		// create something

		// fmt.Println(req.Method)
		// fmt.Println(req.Form["fr"])
		// fmt.Println(req.PostForm["fr"])

		req.ParseForm()
		body_byte, err := ioutil.ReadAll(req.Body)
		// header_byte := req.Header
		if err != nil {
			//statusCode = -200
			return 400, "error"
		}
		body := string(body_byte)
		fmt.Println(body, "\n")
		fmt.Println(req.PostForm["fe"], "\n")
		return 418, "Hello, Post!"
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

	m.Run()
}

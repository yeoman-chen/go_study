package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == '/' {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	t, err := template.ParseFiles("template/html/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	pathInfo := string.Trim(r.URL.Path, "/")
	parts := string.Split(pathInfo, "/")

	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1] + "Action")
	}

	login := &loginController{}

	controller := reflect.ValueOf(login)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)

	method.Call([]reflect.Value{responseValue, requestValue})
}

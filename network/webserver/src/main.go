package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func main() {
	log.Println("main")
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/ajax/", ajaxHandler)
	http.HandleFunc("/", NotFoundHandler)
	http.ListenAndServe(":8888", nil)
}

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
func adminHandler(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	pathinfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathinfo, "/")

	var action = ""

	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	admin := &adminController{}
	controller := reflect.ValueOf(admin)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}

	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})

}

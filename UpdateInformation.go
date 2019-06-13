package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type pagedata struct {
	username string `json:"username" validate:"required"`
	email string	`json:"email" validate:"required,email"`
	phone string	`json:"phone" validate:"required,len=11"`
	address string	`json:"address" validate:"required"`
	Authentication string `json:"Authentication"  validate:"required"`
	password string	`json:"password"  validate:"required"`
}


func main(){
	tmpl := template.Must(template.ParseFiles("templates/TableUpdate.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		pd := pagedata{
			username: r.FormValue("UserNameTxt"),
			email:    r.FormValue("EmailTxt"),
			address:  r.FormValue("Address"),
			password: r.FormValue("PasTxt"),
			phone:    r.FormValue("Phone"),
		}



		tpl.ExecuteTemplate(w, "TableUpdate.html", pd)


	})


		http.ListenAndServe(":8080", nil)
	}

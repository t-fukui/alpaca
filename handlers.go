package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/t-fukui/alpaca/models"
	"github.com/wcl48/valval"
)

var tpl *template.Template

type FormData struct {
	Message models.Message
	Mess    string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Messages := []models.Message{}
	db.Debug().Find(&Messages)
	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
	tpl.Execute(w, &Messages)
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/new.html"))
	tpl.Execute(w, FormData{models.Message{}, ""})
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	Message := models.Message{Name: r.FormValue("Name"),
		Body: r.FormValue("Body")}

	if err := models.MessageValidate(Message); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/new.html"))
		tpl.Execute(w, FormData{Message, Mess})
	} else {
		db.Debug().Create(&Message)
		http.Redirect(w, r, "/", 301)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	Message := models.Message{}
	id := strings.Split(r.URL.Path, "/")[2]
	db.Debug().First(&Message, id)
	tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/edit.html"))
	tpl.Execute(w, FormData{Message, ""})
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	Message := models.Message{}
	id := strings.Split(r.URL.Path, "/")[2]
	fmt.Println("here")
	fmt.Println(db.Debug().First(&Message, id))
	Message.Name = r.FormValue("Name")
	Message.Body = r.FormValue("Body")
	if err := models.MessageValidate(Message); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("templates/layout.html", "templates/edit.html"))
		tpl.Execute(w, FormData{Message, Mess})
	} else {
		db.Debug().Save(&Message)
		http.Redirect(w, r, "/", 301)
	}
}

package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"smartpi/dbcon"
	"smartpi/sensor"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "index.html")
	case "POST":
		s, err := sensor.SensorFromRequest(r)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(fmt.Sprintf("error decoding use {'Type':'sometype', 'Value':1.1}\nError: %e", err)))

		} else {
			fmt.Println(fmt.Sprintf("Type: %s,Unit: %s, Value: %f", s.Type, s.Unit, s.Value))
			dbcon.Insert(s)
			w.Write([]byte("ok"))
		}

	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("static/html/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	fmt.Println("sers")
	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else {
		w.Header().Set("Content-Type", "text/css")
	}

	data, err := ioutil.ReadFile(path[1:])
	if err != nil {
		fmt.Print(err)
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Print(err)
	}

}

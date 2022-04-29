package main

import (
	"fmt"
	"net/http"
	"sesi10-assignment/helpers"
	"strconv"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var data, status = helpers.GetStatus()
		water_str := strconv.Itoa(data[0])
		wind_str := strconv.Itoa(data[1])
		payload := map[string]string{
			"water":  water_str,
			"wind":   wind_str,
			"status": status,
		}
		t.Execute(w, payload)
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

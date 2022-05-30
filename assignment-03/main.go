package main

import (
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type Water struct {
	Number int
	Status string
}
type Wind struct {
	Number int
	Status string
}

var (
	PORT = ":9999"
)

func main() {
	http.HandleFunc("/", waterAndWindSensor)

	log.Default().Println("server running at port :", PORT)
	http.ListenAndServe(PORT, nil)

}

func waterAndWindSensor(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rand.Seed(time.Now().UnixNano())
	water := Water{
		Number: rand.Intn(100) + 1,
		Status: "",
	}
	wind := Wind{
		Number: rand.Intn(100) + 1,
		Status: "",
	}
	if water.Number < 5 {
		water.Status = "Aman"
	} else if water.Number >= 5 && water.Number < 9 {
		water.Status = "Siaga"
	} else if water.Number > 8 {
		water.Status = "Bahaya"
	}
	if wind.Number < 6 {
		wind.Status = "Aman"
	} else if wind.Number >= 6 && wind.Number < 16 {
		wind.Status = "Siaga"
	} else if wind.Number > 15 {
		wind.Status = "Bahaya"
	}

	data := map[string]interface{}{
		"water": map[string]interface{}{
			"number": water.Number,
			"status": water.Status,
		},
		"wind": map[string]interface{}{
			"number": wind.Number,
			"status": wind.Status,
		},
	}

	template.Execute(rw, data)
}

package main

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	peopleHandler := func(w http.ResponseWriter, req *http.Request) {
		// Only allow the GET HTTP verb
		if req.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte("Method not allowed!\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		// Find all people if not query parameters are passed in the request
		if len(req.URL.Query()) == 0 {
			msg, err := json.Marshal(models.AllPeople())
			if err != nil {
				log.Fatal(err)
			}
			_, err = w.Write(append(msg, []byte("\r\n")...))
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	peopleIdHandler := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte("Method not allowed!\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		idParam := strings.TrimPrefix(req.URL.Path, "/people/")
		id, err := uuid.FromString(idParam)
		log.Println(idParam)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte("Not found!\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		
		person, err := models.FindPersonByID(id)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte("Not found!\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		body, err := person.ToJSON()
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(append([]byte(body), []byte("\r\n")...))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/people/", peopleIdHandler)
	// Bind to a TCP port and listen for incoming networking traffic
	// Use the non-privileged TCP port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

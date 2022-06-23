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
		// Find person by phone number
		if len(req.URL.Query()) == 1 {
			phoneNumber, ok := req.URL.Query()["phone_number"]
			if !ok {
				log.Println("missing phone_number parameter")
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("Content-Type", "appliation/json")
				_, err := w.Write([]byte("Bad request!\r\n"))
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			result := models.FindPeopleByPhoneNumber(phoneNumber[0])
			// if no one is found return a 404
			if len(result) == 0 {
				log.Println("person not found")
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "appliation/json")
				_, err := w.Write([]byte("Not found!\r\n"))
				if err != nil {
					log.Fatal(err)
				}
				return
			} else {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")

				_, err := w.Write([]byte("["))
				if err != nil {
					log.Fatal(err)
				}
				for index, person := range result {
					body, err := person.ToJSON()
					if err != nil {
						log.Fatal(err)
					}
					_, err = w.Write([]byte(body))
					if err != nil {
						log.Fatal(err)
					}
					if index < len(result)-1 {
						_, err = w.Write([]byte(","))
						if err != nil {
							log.Fatal(err)
						}
					}
				}
				_, err = w.Write([]byte("]\r\n"))
				if err != nil {
					log.Fatal(err)
				}
				return
			}

		}
		// Find person by First and last name
		firstName, ok := req.URL.Query()["first_name"]
		if !ok {
			log.Println("missing first_name parameter")
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "appliation/json")
			_, err := w.Write([]byte("Bad request!\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		lastName, ok := req.URL.Query()["last_name"]
		if !ok {
			log.Println("missing last_name parameter")
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "appliation/json")
			_, err := w.Write([]byte("Bad request!\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		result := models.FindPeopleByName(firstName[0], lastName[0])
		// if no one is found return a 404
		if len(result) == 0 {
			log.Println("person not found")
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "appliation/json")
			_, err := w.Write([]byte("Not found!\r\n"))
			if err != nil {
				log.Fatal(err)
			}
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")

			_, err := w.Write([]byte("["))
			if err != nil {
				log.Fatal(err)
			}
			for index, person := range result {
				body, err := person.ToJSON()
				if err != nil {
					log.Fatal(err)
				}
				_, err = w.Write([]byte(body))
				if err != nil {
					log.Fatal(err)
				}
				if index < len(result)-1 {
					_, err = w.Write([]byte(","))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			_, err = w.Write([]byte("]\r\n"))
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

	// A health check is needed for monitoring and observability of the application
	// this is often used by load balancers and container orchestrators
	healthCheck := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
	}

	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/people/", peopleIdHandler)
	http.HandleFunc("/health", healthCheck)
	// Bind to a TCP port and listen for incoming networking traffic
	// Use the non-privileged TCP port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

package profile

import (
	"encoding/json"
	"log"
	"net/http"
)

//
//
//
func Index(w http.ResponseWriter, r *http.Request) {

	log.Printf(r.Host)
	log.Printf("ProfileIndex")
	var profiles = GetAll()
	json.NewEncoder(w).Encode(profiles)
}

func Create(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.Host)
	log.Printf("ProfileCreate")

	var p Profile

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var id = Add(p)

	if err := json.NewEncoder(w).Encode(id); err != nil {
		panic(err)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {

	var emailAddress = r.URL.Query().Get("emailAddress")

	log.Println("GetOne " + emailAddress)

	var profile = GetOne(emailAddress)

	if err := json.NewEncoder(w).Encode(profile); err != nil {
		panic(err)
	}

}

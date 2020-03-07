package Profile

import (
	"encoding/json"
	_ "encoding/json"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	log.Printf(r.Host)
	log.Printf("ProfileIndex")
	var profiles = GetAll()
	json.NewEncoder(w).Encode(profiles)
}

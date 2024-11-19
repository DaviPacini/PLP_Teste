package main

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var herois Herois
	allHeroes := herois.ExibeInfosGerais()
	json.NewEncoder(w).Encode(allHeroes)

}

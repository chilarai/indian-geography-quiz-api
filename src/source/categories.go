package source

import (
	"config"
	"net/http"
)

func QuizCategories(w http.ResponseWriter, req *http.Request){
	
	config.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := config.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		Email string
		SessionKey string
		Name string
		PhotoLink string
	}
}
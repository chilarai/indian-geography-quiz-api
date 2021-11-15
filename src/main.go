package main

import (
	"fmt"
	"log"
	"net/http"
	"source"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Quiz server started..")

	http.HandleFunc("/login", source.Login)
	http.HandleFunc("/logout", source.Logout)
	http.HandleFunc("/leaderboard", source.Leaderboard)
	http.HandleFunc("/categories", source.QuizCategories)
	http.HandleFunc("/subcategories", source.QuizSubCategories)
	http.HandleFunc("/entries", source.QuizEntries)

	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Println(err.Error())
	}
}
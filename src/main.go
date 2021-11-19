package main

import (
	"common"
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
	http.HandleFunc("/checkname", source.CheckName)
	http.HandleFunc("/leaderboard", source.Leaderboard)
	http.HandleFunc("/currentscore", source.Currentscore)
	http.HandleFunc("/categories", source.QuizCategories)
	http.HandleFunc("/subcategories", source.QuizSubCategories)
	http.HandleFunc("/entries", source.QuizEntries)
	http.HandleFunc("/updatescore", source.UpdateScore)

	err := http.ListenAndServe(common.Appport, nil)
	if(err != nil){
		log.Println(err.Error())
	}
}
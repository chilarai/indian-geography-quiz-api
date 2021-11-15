package source

import (
	"config"
	"encoding/json"
	"log"
	"net/http"
)

// CategoriesOutputStruct return
type CategoriesOutputStruct struct {
	StatusOut   config.Status     `json:"status"`
	ListDataOut []*CategoriesOutputData `json:"data"`
}

// CategoriesOutputData struct
type CategoriesOutputData struct {
	CategoryID            int
	CategoryName          string
}

func QuizCategories(w http.ResponseWriter, req *http.Request){
	
	config.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := config.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		SessionKey string
	}

	var inputJSON input
	var CategoryName string
	var CategoryID int

	var status config.Status
	var outputStruct CategoriesOutputStruct
	var result []*CategoriesOutputData

	if(inputJSON.SessionKey == ""){
		status = config.Status{
			Code:    403,
			Message: config.InvalidInput,
		}
	} else {
		selCategories, errSel := mydb.Query("SELECT * FROM categories")

		if(errSel != nil){

			status = config.Status{
				Code:    403,
				Message: errSel.Error(),
			}

		} else{
			for selCategories.Next(){
				errSelScan := selCategories.Scan(&CategoryID, &CategoryName)

				if errSelScan != nil {
					log.Println(errSelScan)
				} else {

					elem := CategoriesOutputData{
						CategoryID: CategoryID,
						CategoryName: CategoryName,
					}

					result = append(result, &elem)
				}
			}

			status = config.Status{
				Code:    200,
				Message: config.SuccessMsg,
			}
		}
	}

	outputStruct = CategoriesOutputStruct{
		StatusOut: status,
		ListDataOut: result,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}
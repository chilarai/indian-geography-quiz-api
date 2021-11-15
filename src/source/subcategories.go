package source

import (
	"config"
	"encoding/json"
	"log"
	"net/http"
)

// SubCategoriesOutputStruct return
type SubCategoriesOutputStruct struct {
	StatusOut   config.Status     `json:"status"`
	ListDataOut []*SubCategoriesOutputData `json:"data"`
}

// SubCategoriesOutputData struct
type SubCategoriesOutputData struct {
	SubCategoryID            int
	SubCategoryName          string
}

func QuizSubCategories(w http.ResponseWriter, req *http.Request){
	
	config.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := config.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		SessionKey string
		CategoryID int
	}

	var inputJSON input
	var SubCategoryName string
	var SubCategoryID int

	var status config.Status
	var outputStruct SubCategoriesOutputStruct
	var result []*SubCategoriesOutputData

	if(inputJSON.SessionKey == "" || inputJSON.CategoryID == 0){
		status = config.Status{
			Code:    403,
			Message: config.InvalidInput,
		}
	} else {
		selSubCategories, errSel := mydb.Query("SELECT * FROM quizsubcategories WHERE quiz_id = ?", inputJSON.CategoryID)

		if(errSel != nil){

			status = config.Status{
				Code:    403,
				Message: errSel.Error(),
			}

		} else{
			for selSubCategories.Next(){
				errSelScan := selSubCategories.Scan(&SubCategoryID, &SubCategoryName)

				if errSelScan != nil {
					log.Println(errSelScan)
				} else {

					elem := SubCategoriesOutputData{
						SubCategoryID: SubCategoryID,
						SubCategoryName: SubCategoryName,
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

	outputStruct = SubCategoriesOutputStruct{
		StatusOut: status,
		ListDataOut: result,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}
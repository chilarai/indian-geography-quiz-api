package source

import (
	"common"
	"encoding/json"
	"log"
	"net/http"
)

// EntriesOutputStruct return
type EntriesOutputStruct struct {
	StatusOut   common.Status     `json:"status"`
	ListDataOut []*EntriesOutputData `json:"data"`
}

// EntriesOutputData struct
type EntriesOutputData struct {
	CategoryID            int
	CategoryName          string
}

func QuizEntries(w http.ResponseWriter, req *http.Request){
	
	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		SessionKey 		string
		CategoryID 		int
		SubCategoryID 	int
	}

	var inputJSON input
	var CategoryName string
	var CategoryID int

	var status common.Status
	var outputStruct CategoriesOutputStruct
	var result []*CategoriesOutputData

	if(inputJSON.SessionKey == ""){
		status = common.Status{
			Code:    403,
			Message: common.InvalidInput,
		}
	} else {
		selCategories, errSel := mydb.Query("SELECT * FROM categories")

		if(errSel != nil){

			status = common.Status{
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

			status = common.Status{
				Code:    200,
				Message: common.SuccessMsg,
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
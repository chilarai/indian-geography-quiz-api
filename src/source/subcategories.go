package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// SubCategoriesOutputStruct return
type SubCategoriesOutputStruct struct {
	StatusOut   common.Status     `json:"status"`
	ListDataOut []*SubCategoriesOutputData `json:"data"`
}

// SubCategoriesOutputData struct
type SubCategoriesOutputData struct {
	SubCategoryID            int
	SubCategoryName          string
}

func QuizSubCategories(w http.ResponseWriter, req *http.Request){
	
	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		SessionKey string
		CategoryID int
	}

	var inputJSON input
	var SubCategoryName string
	var SubCategoryID int

	var status common.Status
	var outputStruct SubCategoriesOutputStruct
	var result []*SubCategoriesOutputData

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {

		if(inputJSON.SessionKey == "" || inputJSON.CategoryID == 0){
			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			selSubCategories, errSel := mydb.Query("SELECT * FROM quizsubcategories WHERE quiz_id = ?", inputJSON.CategoryID)

			if(errSel != nil){

				status = common.Status{
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

				status = common.Status{
					Code:    200,
					Message: common.SuccessMsg,
				}
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
package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CategoriesOutputStruct return
type CategoriesOutputStruct struct {
	StatusOut   common.Status     `json:"status"`
	ListDataOut []*CategoriesOutputData `json:"data"`
}

// CategoriesOutputData struct
type CategoriesOutputData struct {
	CategoryID            int
	CategoryName          string
}

func QuizCategories(w http.ResponseWriter, req *http.Request){
	
	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		SessionKey string
	}

	var inputJSON input
	var CategoryName string
	var CategoryID int

	var status common.Status
	var outputStruct CategoriesOutputStruct
	var result []*CategoriesOutputData

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {

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
	}

	outputStruct = CategoriesOutputStruct{
		StatusOut: status,
		ListDataOut: result,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}
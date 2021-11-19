package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CheckNameOutputStruct return
type CheckNameOutputStruct struct {
	StatusOut           common.Status  `json:"status"`
}

func CheckName(w http.ResponseWriter, req *http.Request){

	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		Name string
	}

	var inputJSON input
	var status common.Status
	var outputStruct CheckNameOutputStruct
	var id int64

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println("Here", errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println("There", errParse)

	} else {

		if( inputJSON.Name == ""){

			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			errUser := mydb.QueryRow("SELECT id FROM users WHERE name = ?", inputJSON.Name).Scan(&id)

			log.Println(errUser, id)

			if(errUser != nil){
				status = common.Status{
					Code:    200,
					Message: common.SuccessMsg,
				}
			} else {
				status = common.Status{
					Code:    403,
					Message: common.FailedMsg,
				}
			}
		}
	}

	outputStruct = CheckNameOutputStruct{
		StatusOut: status,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}
package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CheckEmailOutputStruct return
type CheckEmailOutputStruct struct {
	StatusOut           common.Status  `json:"status"`
}

func CheckEmail(w http.ResponseWriter, req *http.Request){

	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		Email string
	}

	var inputJSON input
	var status common.Status
	var outputStruct CheckEmailOutputStruct
	var id int64

	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println("Here", errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println("There", errParse)

	} else {

		if( inputJSON.Email == ""){

			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			errUser := mydb.QueryRow("SELECT email FROM users WHERE email = ?", inputJSON.Email).Scan(&id)

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

	outputStruct = CheckEmailOutputStruct{
		StatusOut: status,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}
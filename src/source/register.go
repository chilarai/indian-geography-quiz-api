package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// RegisterOutputData struct
type RegisterOutputData struct {
	UserID    int64    `json:"userId"`
	Name    string `json:"name"`
	SessionToken string `json:"sessionToken"`
}

// RegisterOutputStruct return
type RegisterOutputStruct struct {
	StatusOut           common.Status  `json:"status"`
	RegisterDataOut RegisterOutputData `json:"data"`
}

func Register(w http.ResponseWriter, req *http.Request){

	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		Name string
		Email string
		Password string
	}

	var inputJSON input
	var status common.Status
	var outputStatus RegisterOutputData
	var outputStruct RegisterOutputStruct


	readData, errRead := ioutil.ReadAll(req.Body)
	if errRead != nil {
		log.Println(errRead)
	}

	errParse := json.Unmarshal(readData, &inputJSON)
	if errParse != nil {
		log.Println(errParse)

	} else {

		if( inputJSON.Password == "" || inputJSON.Email == "" || inputJSON.Name == ""){

			status = common.Status{
				Code:    403,
				Message: common.InvalidInput,
			}
		} else {
			passHash, _ := bcrypt.GenerateFromPassword([]byte(inputJSON.Password), 14)

			reg, _ := regexp.Compile(`\W`)
			finalSessionHash := reg.ReplaceAllString(string(passHash), "")[20:]

			insUser, errUser := mydb.Exec("INSERT INTO users (session_key, name, email, password) VALUES (?,?,?,?)",finalSessionHash,  inputJSON.Name, inputJSON.Email, passHash)

			if(errUser != nil){
				status = common.Status{
					Code:    403,
					Message: errUser.Error(),
				}
			} else {
				userId, _ := insUser.LastInsertId()

				outputStatus = RegisterOutputData {
					UserID: userId,
					Name: inputJSON.Name,
					SessionToken: finalSessionHash,
				}

				status = common.Status{
					Code:    200,
					Message: common.SuccessMsg,
				}
			}
		}
	}

	outputStruct = RegisterOutputStruct{
		StatusOut: status,
		RegisterDataOut: outputStatus,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}
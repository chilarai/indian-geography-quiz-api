package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

// EntriesOutputStruct return
type EntriesOutputStruct struct {
	StatusOut   common.Status     `json:"status"`
	QuizCategory string `json:"category"`
	QuizSubCateogry string `json:"sub-category"`
	ListDataOut []*EntriesOutputData `json:"data"`
}

// EntriesOutputData struct
type EntriesOutputData struct {
	RightOption      string
	Options          []string
	ImageLink		 string
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
	var Entry, ImageLink, Category, SubCategory string

	var status common.Status
	var outputStruct EntriesOutputStruct
	var result []*EntriesOutputData
	var answers []string
	var answersMap =  make(map[string]string)
	var answerOptions []string
	

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
			selEntries, errSel := mydb.Query("SELECT entry, image_link FROM entries WHERE quiz_id = ? AND quizsubcategory_id = ?", inputJSON.CategoryID, inputJSON.SubCategoryID)

			if(errSel != nil){

				status = common.Status{
					Code:    403,
					Message: errSel.Error(),
				}

			} else{
				for selEntries.Next(){
					errSelScan := selEntries.Scan(&Entry, &ImageLink)

					if errSelScan != nil {
						log.Println(errSelScan)
					} else {

						answers = append(answers, Entry)
						answersMap[Entry] = ImageLink
					}
				}

				answersLength := len(answers)
				for ans, image := range answersMap {

					answerOptions = nil

					if(answersLength >= 4){
						for len(answerOptions) <= 4 {
							CreateAnswerOptions(answers, &answerOptions)
						}
					} else if (answersLength >= 3) {

						for len(answerOptions) <= 3 {
							CreateAnswerOptions(answers, &answerOptions)
						}
					} else {

						for len(answerOptions) <= 2 {
							CreateAnswerOptions(answers, &answerOptions)
						}
					}

					elem := EntriesOutputData{
						RightOption : ans,
						Options : answerOptions,
						ImageLink : image,
					}

					result = append(result, &elem)
				}

				errSelQuiz := mydb.QueryRow("SELECT quizsubcategories.subcategory_name, quizes.quiz_name FROM quizsubcategories LEFT JOIN quizes ON quizes.id = quizsubcategories.quiz_id WHERE quizsubcategories.id = ?", inputJSON.SubCategoryID).Scan(&SubCategory, &Category)

				if errSelQuiz != nil{
					log.Println(errSelQuiz)
				}

				status = common.Status{
					Code:    200,
					Message: common.SuccessMsg,
				}
			}
		}
	}

	outputStruct = EntriesOutputStruct{
		StatusOut: status,
		QuizCategory: Category,
		QuizSubCateogry: SubCategory,
		ListDataOut: result,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}


func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func CreateAnswerOptions(answers []string, answerOptions *[]string){
	randomIndex := rand.Intn(len(answers))
	pick := answers[randomIndex]

	_, matches := Find(*answerOptions, pick)

	if !matches{
		*answerOptions = append(*answerOptions, pick)
	}
}
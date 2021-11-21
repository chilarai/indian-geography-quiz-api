package source

import (
	"common"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
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
	Title			 string
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
	var Entry, ImageLink, Category, SubCategory, Title string

	var status common.Status
	var outputStruct EntriesOutputStruct
	var result []*EntriesOutputData
	var answers []string
	var answersMap =  make(map[string]string)
	var answersTitleMap =  make(map[string]string)
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
			selEntries, errSel := mydb.Query("SELECT title, entry, image_link FROM entries WHERE quiz_id = ? AND quizsubcategory_id = ?", inputJSON.CategoryID, inputJSON.SubCategoryID)

			if(errSel != nil){

				status = common.Status{
					Code:    403,
					Message: errSel.Error(),
				}

			} else{
				for selEntries.Next(){
					errSelScan := selEntries.Scan(&Title, &Entry, &ImageLink)

					if errSelScan != nil {
						log.Println(errSelScan)
					} else {

						answers = append(answers, Entry)
						answersMap[Entry] = ImageLink
						answersTitleMap[Entry] = Title
					}
				}

				answersLength := len(answers)

				for ans, image := range answersMap {

					answerOptions = nil

					// Enter random options other than the answer
					if(answersLength > 3){
						for len(answerOptions) < 3 {
							CreateAnswerOptions(answers, &answerOptions, ans)
						}
					} else if (answersLength > 2) {

						for len(answerOptions) < 2 {
							CreateAnswerOptions(answers, &answerOptions, ans)
						}
					} else {

						for len(answerOptions) < 1 {
							CreateAnswerOptions(answers, &answerOptions, ans)
						}
					}

					// Append the right answer to the options
					answerOptions = append(answerOptions, ans)
					
					// Randomize the answer options
					rand.Seed(time.Now().UnixNano())
					rand.Shuffle(len(answerOptions), func(i, j int) { answerOptions[i], answerOptions[j] = answerOptions[j], answerOptions[i] })

					elem := EntriesOutputData{
						RightOption : ans,
						Options : answerOptions,
						ImageLink : image,
						Title : answersTitleMap[ans],
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

// check if the value exists in slice
func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

// function to insert random values in answer options
func CreateAnswerOptions(answers []string, answerOptions *[]string, actualAnswer string){
	randomIndex := rand.Intn(len(answers))
	pick := answers[randomIndex]

	if pick != actualAnswer {
		_, matches := Find(*answerOptions, pick)

		if !matches{
			*answerOptions = append(*answerOptions, pick)
		}
	}
}
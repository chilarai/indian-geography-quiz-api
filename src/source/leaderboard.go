package source

import (
	"common"
	"encoding/json"
	"log"
	"net/http"
)

// LeaderboardsOutputStruct return
type LeaderboardsOutputStruct struct {
	StatusOut   common.Status     `json:"status"`
	ListDataOut []*LeaderboardsOutputData `json:"data"`
	ListCurrentUserData CurrentUserOutputData `json:"currentUserData"`
}

// LeaderboardsOutputData struct
type LeaderboardsOutputData struct {
	UserID			int
	Score 			int
	CategoryID 		int
	SubCategoryID 	int
	Name            string
	PhotoLink       string
}

// CurrentUserOutputData struct
type CurrentUserOutputData struct {
	Score 			int
	CategoryID 		int
	SubCategoryID 	int
}

func Leaderboard(w http.ResponseWriter, req *http.Request){
	
	common.Cors(&w)
	w.Header().Set("Content-Type", "application/json")

	mydb := common.Mysqlconnect()
	defer mydb.Close()

	type input struct {
		UserID 			int
		CategoryID 		int
		SubCategoryID 	int
		SessionKey 		string
		QuizDate 		string
	}

	var inputJSON input
	var Name, PhotoLink string
	var UserID, Score, CategoryID, SubCategoryID, CurrentUserScore, CurrentUserCategoryID, CurrentUserSubCategoryID int

	var status common.Status
	var outputStruct LeaderboardsOutputStruct
	var result []*LeaderboardsOutputData
	var currentUserResult CurrentUserOutputData

	if(inputJSON.SessionKey == "" || inputJSON.QuizDate == ""){
		status = common.Status{
			Code:    403,
			Message: common.InvalidInput,
		}
	} else {
		selLeaderboards, errSel := mydb.Query("SELECT users.id, users.name, users.photo_link, leaderboards.score, leaderboards.quiz_id, leaderboards.quizsubcategory_id FROM leaderboards LEFT JOIN users ON users.id = leaderboards.user_id WHERE leaderboards.score_date = ? ORDER BY leaderboards.score DESC LIMIT 0, 20", inputJSON.QuizDate)

		if(errSel != nil){

			status = common.Status{
				Code:    403,
				Message: errSel.Error(),
			}

		} else{
			for selLeaderboards.Next(){
				errSelScan := selLeaderboards.Scan(&UserID, &Name, &PhotoLink, &Score, &CategoryID, &SubCategoryID)

				if errSelScan != nil {
					log.Println(errSelScan)
				} else {

					elem := LeaderboardsOutputData{
						UserID			:UserID,
						Score 			:Score,
						CategoryID 		:CategoryID,
						SubCategoryID 	:SubCategoryID,
						Name            :Name,
						PhotoLink       :PhotoLink,
					}

					result = append(result, &elem)
				}
			}

			errSelCurrUser := mydb.QueryRow("SELECT leaderboards.score, leaderboards.quiz_id, leaderboards.quizsubcategory_id FROM leaderboards WHERE leaderboards.score_date = ? AND leaderboards.user_id =? ", inputJSON.QuizDate, inputJSON.UserID).Scan(&CurrentUserScore, &CurrentUserCategoryID, &CurrentUserSubCategoryID)

			if(errSelCurrUser != nil){
				log.Println(errSelCurrUser)
			} else {
				currentUserResult = CurrentUserOutputData{
					Score 			:CurrentUserScore,
					CategoryID 		:CurrentUserCategoryID,
					SubCategoryID 	:CurrentUserSubCategoryID,
				}
			}

			status = common.Status{
				Code:    200,
				Message: common.SuccessMsg,
			}
		}

		
	}

	outputStruct = LeaderboardsOutputStruct{
		StatusOut: status,
		ListDataOut: result,
		ListCurrentUserData: currentUserResult,
	}

	output, _ := json.Marshal(outputStruct)

	w.Write(output)
}
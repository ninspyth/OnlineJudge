package submissionModel

type SubmissionRequest struct {
	UserId string				`json:"user_id"`
	ProblemId string				`json:"problem_id"`	
	UserCode string 		`json:"user_code"`
	Input string 				`json:"input"`
	Language string 		`json:"language"`
}


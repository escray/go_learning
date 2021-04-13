package go_learning

type BasicInfo struct {
	Name string `json:"name"`
	Age int `json:"age`
}

// easyjson:json
type BasicInfoList []BasicInfo

type JobInfo struct {
	Skills []string `json:"skills"`
}

type EmployeeJSON struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo JobInfo `json:"job_info"`
}

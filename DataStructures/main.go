package DataStructures

type Task struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	DateAdded      string `json:"dateAdded"`
	WorkingDateEst string `json:"workingPeriodEst"`
}

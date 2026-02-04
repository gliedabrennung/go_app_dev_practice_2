package models

type TaskResponse struct {
	ID    int64  `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

package model

type Todo struct {
	 ID string `json:"id" bson:"id"`
	 Title string `json:"title" bson:"title"`
}
package entity

type Person struct {
	name string `json:"name" binding:"required"`
	age  int `json:"age" binding:"required"`
}


type Video struct{
	Title 			string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	URL 				string `json:"url" binding:"required"`
	Author 			Person `json:"author" binding:"required"`
}
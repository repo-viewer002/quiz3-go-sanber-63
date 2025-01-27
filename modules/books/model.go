package books

import (
	"time"
)

type Book struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Image_Url    string    `json:"image_url"`
	Release_Year int       `json:"release_year"`
	Price        int       `json:"price"`
	Total_Page   int       `json:"total_page"`
	Thickness    string    `json:"thickness"`
	Category_Id  int       `json:"category_id"`
	Created_At   time.Time `json:"created_at"`
	Created_By   string    `json:"created_by"`
	Modified_At  time.Time `json:"modified_at"`
	Modified_By  string    `json:"modified_by"`
}

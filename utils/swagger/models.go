package swagger

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CategoryInput struct {
	Name string `json:"name"`
}

type BookInput struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image_Url    string `json:"image_url"`
	Release_Year int    `json:"release_year"`
	Price        int    `json:"price"`
	Total_Page   int    `json:"total_page"`
	Thickness    string `json:"thickness"`
	Category_Id  int    `json:"category_id"`
}

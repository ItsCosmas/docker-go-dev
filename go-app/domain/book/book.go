package book

// Book struct
type Book struct {

	Name string `json:"name" bson:"name"`
	Pages uint `json:"pages" bson:"page_count"`
}

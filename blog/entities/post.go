package entities

type Post struct {
	ID                int32
	Category          *Category
	Title, Text, Slug string
}

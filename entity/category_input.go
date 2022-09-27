package entity

type (
	CreateCategoryInput struct {
		Name string `json:"name" binding:"required"`
		Slug string `json:"slug"`
	}

	CategoryDetailInput struct {
		ID int `uri:"id" binding:"required"`
	}
)

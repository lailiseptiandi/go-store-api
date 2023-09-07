package entity

type (
	CreateProduct struct {
		CategoryID   int     `json:"category_id" binding:"required"`
		Name         string  `json:"name" binding:"required"`
		Slug         string  `json:"slug"`
		Descriptions string  `json:"descriptions" binding:"required"`
		Price        float64 `json:"price" binding:"required"`
		Qty          int     `json:"quantity" binding:"required"`
	}

	ProductDetail struct {
		ID int `uri:"id" binding:"required"`
	}
)

type (
	StoreImageByProduct struct {
		ImageName string `json:"image_name"`
		ProductID int    `json:"product_id"`
	}

	ImageByProduct struct {
		ID int `uri:"id" binding:"required"`
	}
)

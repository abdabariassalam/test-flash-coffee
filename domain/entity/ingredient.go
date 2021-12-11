package entity

type Ingredient struct {
	Id    *int   `gorm:"id"`
	Name  string `gorm:"name"`
	Color int    `gorm:"color"`
	Img   string `gorm:"img"`
}

func (Ingredient) TableName() string {
	return "ingredient"
}

type IngredientCategory struct {
	Id          *int   `gorm:"id"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
}

func (IngredientCategory) TableName() string {
	return "ingredient_category"
}

type IngredientCategoryIngredient struct {
	IngredientCategoryID int `gorm:"ingredient_category_id"`
	IngredientID         int `gorm:"ingredient_id"`
}

func (IngredientCategoryIngredient) TableName() string {
	return "ingredient_category_ingredient"
}

type CreateIngredientRequest struct {
	Name               string             `json:"name"`
	Color              int                `json:"color"`
	IngredientCategory IngredientCategory `json:"ingredient_category"`
	Image              File
}

type File struct {
	ImageContentType string
	ImageName        string
	ImageFile        []byte
	ImageSize        int64
}

type CreateIngredientResponse struct {
	Id                 *int
	Name               string
	Color              int
	IngredientCategory IngredientCategory
	Image              string
}

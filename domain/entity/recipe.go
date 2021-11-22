package entity

type Recipe struct {
	Id          *int   `gorm:"id"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
	AuthorID    *int   `gorm:"author_id"`
}

func (Recipe) TableName() string {
	return "recipe"
}

type RecipeCategory struct {
	Id   *int   `gorm:"id"`
	Name string `gorm:"name" json:"name"`
}

func (RecipeCategory) TableName() string {
	return "recipe_category"
}

type RecipeCategoryRecipe struct {
	RecipeCategoryID int `gorm:"recipe_category_id"`
	RecipeID         int `gorm:"recipe_id"`
}

func (RecipeCategoryRecipe) TableName() string {
	return "recipe_category_recipe"
}

type Recipes struct {
	Id          int
	Name        string
	Description string
	Author      User
}

type CreateRecipe struct {
	Id             *int
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	RecipeCategory RecipeCategory `json:"recipe_category"`
	User           User           `json:"user"`
}

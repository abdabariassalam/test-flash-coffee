package entity

type Step struct {
	Id          int    `gorm:"id"`
	RecipeId    int    `gorm:"recipe_id"`
	StepNumber  int    `gorm:"step_number"`
	Description string `gorm:"description"`
	Timer       int    `gorm:"timer"`
	Image       string `gorm:"image"`
}

func (Step) TableName() string {
	return "step"
}

type StepIngredient struct {
	RecipeId     int    `gorm:"recipe_id" json:"recipe_id"`
	IngredientId int    `gorm:"ingredient_id" json:"ingredient_id"`
	StepId       int    `gorm:"step_id" json:"step_id"`
	Amount       int    `gorm:"amount" json:"amount"`
	Unit         string `gorm:"unit" json:"unit"`
}

func (StepIngredient) TableName() string {
	return "step_ingredients"
}

type CreateStep struct {
	Id             *int           `gorm:"id"`
	StepNumber     int            `json:"step_number"`
	Description    string         `json:"description"`
	Timer          int            `json:"timer"`
	StepIngredient StepIngredient `json:"step_ingredient"`
	Image          File
}

type CreateStepResponse struct {
	Id             *int           `json:"id"`
	RecipeId       int            `json:"recipe_id"`
	StepNumber     int            `json:"step_number"`
	Description    string         `json:"description"`
	Timer          int            `json:"timer"`
	StepIngredient StepIngredient `json:"step_ingredient"`
	Image          string         `json:"image"`
}

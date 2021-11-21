package entity

type Recipe struct {
	Id          int    `gorm:"id"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
	Author      User
}

func (Recipe) TableName() string {
	return "recipe"
}

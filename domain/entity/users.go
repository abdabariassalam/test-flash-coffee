package entity

type User struct {
	Id   int    `gorm:"id"`
	Name string `gorm:"name"`
}

func (User) TableName() string {
	return "users"
}

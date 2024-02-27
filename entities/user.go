package entities

type UserDTO struct {
	Email      string `json:"email"`
	VerifyCode int    `json:"verify_code"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Grade      int64  `json:"grade"`
	Class      int64  `json:"class"`
	Number     int64  `json:"number"`
}

type User struct {
	Email    string `gorm:"primaryKey"`
	Session  string `gorm:"primaryKey"`
	Password string `gorm:"duplicate key"`
	Name     string `gorm:"duplicate key"`
	Grade    int64  `gorm:"duplicate key"`
	Class    int64  `gorm:"duplicate key"`
	Number   int64  `gorm:"duplicate key"`
}

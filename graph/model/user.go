package model

type User struct {
	ID            string  `json:"id"`
	Name          *string `json:"name"`
	Email         string  `json:"email"`
	GoogleId      *string `json:"google_id"`
	Image         *string `json:"image"`
	EmailVerified *string `json:"email_verified"`
	Tasks         []*Task `json:"tasks" gorm:"foreignKey:UserID"`
}

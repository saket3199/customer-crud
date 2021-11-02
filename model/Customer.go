package model

type Customer struct {
	Model
	Email    string  `gorm:" type:varchar(100)" json:"email"`
	UserPass string  `gorm:" type:varchar(100)" json:"userPass"`
	Fname    string  `gorm:" type:varchar(20)" json:"fName"`
	Lname    string  `gorm:" type:varchar(20)" json:"lName"`
	Age      int     `gorm:" type:int" json:"age"`
	IsMale   *bool   `gorm:" type:tinyint" json:"isMale"`
	Orders   []Order `json:"orders"`
}

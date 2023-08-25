package interfaces

type ADDRESS struct {
	LINE1   string `json:"line1" binding:"required,max=100" gorm:"type:varchar(100)"`
	LINE2   string `json:"line2" binding:"required,max=100" gorm:"type:varchar(100)"`
	CITY    string `json:"city" binding:"required,max=100" gorm:"type:varchar(100)"`
	STATE   string `json:"state" binding:"required,max=100" gorm:"type:varchar(100)"`
	PINCODE string `json:"pincode" binding:"max=6,min=6" gorm:"size:6"`
}

type User struct {
	// ID           uint64  `gorm:"primary_key;auto_increment" json:"id"`
	USERNAME     string  `json:"username" binding:"required,max=100,min=3" gorm:"type:varchar(50);UNIQUE"`
	NAME         string  `json:"name" binding:"required,max=100"`
	MOBILE       string  `json:"mobile" binding:"required,max=10,min=10"`
	PASSWORD     string  `json:"password" binding:"required,max=100"`
	EMAIL        string  `json:"email" binding:"required,max=100"`
	USER_ADDRESS ADDRESS `json:"address" gorm:"foreignkey:ADDRESS_ID" `
	ADDRESS_ID   uint64  `json:"-"`
}

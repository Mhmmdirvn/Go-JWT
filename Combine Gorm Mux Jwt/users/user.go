package users

type User struct {
	Id          int    `gorm:"primarykey" json:"id"`
	NamaLengkap string `gorm:"varchar(225)" json:"nama_lengkap"`
	Username    string `gorm:"varchar(225)" json:"username"`
	Password    string `gorm:"varchar(225)" json:"password"`
}

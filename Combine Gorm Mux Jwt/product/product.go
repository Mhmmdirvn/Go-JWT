package product

type Product struct {
	Id int `gorm:"primarykey" json:"id"`
	ProductName string `gorm:"varchar(255)" json:"product_name"`
	Price int `gorm:"int(200)" json:"price"`
}
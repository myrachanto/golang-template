package models

import (
	"time"
	"github.com/jinzhu/gorm"
)
type Customer struct {
	Name string `gorm:"not null"`
	Company string `gorm:"not null"`
	Phone string `gorm:"not null"`
	Address string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Invoice []Invoice `gorm:"foreignkey:UserRefer"`//has many invoices
	gorm.Model
}
type Invoice struct {
	CustomerID uint64 `gorm:"not null"`
	Customer Customer `gorm:"foreignKey:CustomerID; not null"`
	Title string `gorm:"not null"`
	Dated time.Time `gorm:"not null"`
	Due_date time.Time `gorm:"not null"`
	Discount float64 `gorm:"not null"`
	Sub_total float64 `gorm:"not null"`
	Total float64 `gorm:"not null"`
	InvoiceItem []InvoiceItem `gorm:"foreignkey:UserRefer"`//has many invoiceitems
	gorm.Model
}
type InvoiceItem struct {
	InvoiceID uint64 `gorm:"not null"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID; not null"`
	Description string `gorm:"not null"`
	Qty uint64 `gorm:"not null"`
	Unit_price float64 `gorm:"not null"`
	gorm.Model
}


/*

        'id','company', 'email', 'name', 'phone', 'address', 'created_at'

        'company', 'email', 'name', 'phone', 'address'
type Auth struct {
	User User `gorm:"foreignKey:UserID; not null"`
	UserID int `json:"userid"`
	Token string `gorm:"type:varchar(200); not null"`
	ModelBase
}*/
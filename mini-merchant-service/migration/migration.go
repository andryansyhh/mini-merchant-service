package migration

import "time"

type Users struct {
	UserID    string 	`gorm:"PrimaryKey"`
	FullName  string
	Email     string 	`gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Outlet    []Outlets `gorm:"ForeignKey:UserID"`
}

type Outlets struct {
	OutletID   string 	`gorm:"PrimaryKey"`
	OutletName string
	Picture    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    Products `gorm:"ForeignKey:OutletID"`
	UserID     string   `gorm:"index"`
}

type Products struct {
	ProductID    string 	  `gorm:"PrimaryKey"`
	ProductName  string
	Price        int64
	Sku          string
	Picture      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	OutletID     string       `gorm:"index"`
	DisplayImage ImageProduct `gorm:"ForeignKey:ProductID"`
}

type ImageProduct struct {
	ImageProductID string
	DisplayImage   string
	ProductID      string `gorm:"index"`
}

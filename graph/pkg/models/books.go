package models

// import (
//     "gorm.io/driver/postgres"
//     "gorm.io/gorm"
// )
type Book struct {
    Id     int    `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Desc   string `json:"desc"`
}
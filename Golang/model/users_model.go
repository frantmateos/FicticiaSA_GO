package model

type User struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	Nombre       string `gorm:"type:varchar(600);not null"`
	Password     string `gorm:"type:varchar(350);null"`
	Genero       string `gorm:"type:varchar(350);not null"`
	Atributos    string `gorm:"type:varchar(600);not null"`
	Maneja       bool   `gorm:"not null"`
	Lentes       bool   `gorm:"not null"`
	Diabetico    bool   `gorm:"not null"`
	Enfermedades string `gorm:"type:varchar(600);not null"`
	Admin        bool   `gorm:"not null"`
	Estado       bool   `gorm:"not null"`
}

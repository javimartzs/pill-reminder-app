package models

type Medicamento struct {
	NationalCode string `gorm:"unique;not null"`
	Ingredient   string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Status       string `gorm:"not null"`
	Generic      string `gorm:"not null"`
	Subgroup     string `gorm:"not null"`
	Receta       string `gorm:"not null"`
	Laboratorio  string `gorm:"not null"`
}

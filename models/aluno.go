package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

/*
- gorm.Model definition

o gorm.Model equivale a:

type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
*/

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero,regexp=^[A-Z a-z]*$"`
	CPF  string `json:"CPF" validate:"len=11,regexp=^[0-9]*$"`
	RG   string `json:"RG" validate:"len=9,regexp=^[0-9]*$"`
}

func Validate(a *Aluno) error {
	if err := validator.Validate(a); err != nil {
		return err
	}

	return nil
}

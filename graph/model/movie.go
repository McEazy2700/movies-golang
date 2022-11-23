package model

import (
	"gorm.io/gorm"
)
type Movie struct {
  gorm.Model
  Title string `json:"title"`
  ImdbCode string `json:"imdbcode"`
  Views uint `json:"views"`
}

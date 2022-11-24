package model

import "time"

type Movie struct {
  ID int `json:"id"` 
  Title string `json:"title"`
  ImdbCode string `json:"imdbCode"`
  Views int `json:"views"`
  UpdatedAt time.Time `json:"updatedAt"`
  CreatedAt time.Time `json:"createdAt"`
}

type MovieSuccess struct {
  Movie Movie `json:"movie"`
  Success bool `json:"success"`
}

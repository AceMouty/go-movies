package models

import "time"

type Movie struct {
  ID int `json:"id"`
  Title string `json:"title"`
  ReleaseDate time.Time `json:"releaseDate"`
  RunTime int `json:"runtime"`
  MPAARating string `json:"mpaaRating"`
  Description string `json:"description"`
  Image string `json:"image"`
  CreatedAt time.Time `json:"-"`
  UpdatedAt time.Time `json:"-"`
}

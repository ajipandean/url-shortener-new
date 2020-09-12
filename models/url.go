package models

import (
  "time"
  "math/rand"
  "github.com/Kamva/mgm/v3"
  "github.com/catinello/base62"
  h "github.com/ajipandean/url-shortener-new/helpers"
)

type URL struct {
  mgm.DefaultModel `bson:",inline"`
  Ref    int       `json:"ref"`
  Url    string    `json:"url"`
  Base   string    `json:"base"`
  Clicks uint      `json:"clicks"`
}

func NewURL(url string) *URL {
  return &URL{Url: url}
}

func (u *URL) Creating() error {
  rand.Seed(time.Now().UnixNano())

  randomNumber := h.GenerateRandomNumber(00001, 99999)
  u.Ref = randomNumber

  encodedRef := base62.Encode(u.Ref)
  u.Base = encodedRef

  return nil
}

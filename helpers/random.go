package helpers

import (
  "math/rand"
)

func GenerateRandomNumber(min, max int) int {
  return rand.Intn(max - min) + min
}

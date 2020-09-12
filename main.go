package main

import (
  "os"
  "fmt"
  "net/http"
  "github.com/joho/godotenv"
  "github.com/ajipandean/url-shortener-new/routes"
  "github.com/ajipandean/url-shortener-new/config/database"
)

func init() {
  if err := godotenv.Load(); err != nil {
    fmt.Println(err.Error())
  }

  database.Setup()
}

func main() {
  routes.Setup()

  port := os.Getenv("PORT")
  address := fmt.Sprintf(":%s", port)
  fmt.Printf("Server running on %s\n", address)
  srv := &http.Server{Addr: address}
  if err := srv.ListenAndServe(); err != nil {
    fmt.Println(err.Error())
  }
}

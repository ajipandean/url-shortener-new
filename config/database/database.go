package database

import (
  "os"
  "fmt"
  "time"
  "github.com/Kamva/mgm/v3"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() {
  databaseConfig := &mgm.Config{CtxTimeout: time.Second * 10}
  databaseName := os.Getenv("DATABASE_NAME")
  databaseUri := os.Getenv("DATABASE_URI")
  err := mgm.SetDefaultConfig(
    databaseConfig,
    databaseName,
    options.Client().ApplyURI(databaseUri),
  )
  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Println("Connected to database")
  }
}

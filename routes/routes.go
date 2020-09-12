package routes

import (
  "net/http"
  "encoding/json"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/gorilla/mux"
  "github.com/Kamva/mgm/v3"
  "github.com/catinello/base62"
  m "github.com/ajipandean/url-shortener-new/models"
)

type RequestBody struct {
  Url string `json:"url" form:"url"`
}

func Setup() {
  r := mux.NewRouter()
  r.HandleFunc("/urls", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
      var body RequestBody
      if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
      }

      url := m.NewURL(body.Url)
      if err := mgm.Coll(url).Create(url); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      res, err := json.MarshalIndent(url, "", "  ")
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      w.Header().Set("Content-Type", "application/json")
      w.Write(res)
    }
  })
  r.HandleFunc("/{base}", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
      vars := mux.Vars(r)
      encodedBase := vars["base"]
      baseRef, err := base62.Decode(encodedBase)
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      url := new(m.URL)
      if err := mgm.Coll(url).First(bson.M{"ref": baseRef}, url); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      http.Redirect(w, r, url.Url, http.StatusSeeOther)
    }
  })
  http.Handle("/", r)
}

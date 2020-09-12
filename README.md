# URL Shortener: New
New improved source code for url-shortener-server. Tech stack: Golang | Gorilla Mux | MongoDB | MGM

## Prerequisite
Make sure you have golang version 1.14 or later installed before clone and run this project

## Steps to run this project
#### Clone this project
```bash
git clone git@github.com:ajipandean/url-shortener-new.git
// or
git clone https://github.com/ajipandean/url-shortener-new.git
```
#### Move to the cloned directory (in this case /url-shortener-new)
```bash
cd url-shortener-new
```
#### Copy paste .env.example and rename to .env
```bash
cp .env.example .env
```
#### Install dependencies to your workspace
```bash
go mod vendor
```
#### Run the project
```bash
go run main.go
```

## What you can do here?
This project is intended to shorten the long URL to be shorter
#### Shortening URL
```
// request
POST /urls HTTP/1.1
Accept: application/json
Host: localhost:8080
Content-Type: application/json

{"url": URL_TO_SHORTEN}

// response
HTTP/1.1 200 OK
Content-Type: application/json

{
  "_id": objectId,
  "ref": int,
  "url": string,
  "base": string,
  "created_at": date,
  "updated_at": date
}
```
You simply make POST request to localhost:8080/urls with "url" as request body formatted in JSON
#### Redirect using shortened URL
```
// request
GET /{base} HTTP/1.1
Host: localhost:8080

// response
HTTP/1.1 303 See Other
Content-Type: text/html
Location: URL_TO_BE_REDIRECTED
```
You simply make GET request to localhost:8080/{base} with base (encoded ref) as path parameters

---

Thank you :)

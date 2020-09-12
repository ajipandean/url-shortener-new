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

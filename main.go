package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-macaron/binding"
	"github.com/jinzhu/gorm"
	"gopkg.in/macaron.v1"
)

var m *macaron.Macaron

var conf *Config

var db *gorm.DB

func main() {
	conf = initConfig()

	m = initMacaron()

	db = initDb()

	m.Get("/", newPageData, homeView)
	m.Get("/kontakt/", newPageData, kontaktView)

	// m.Post("/", binding.Bind(SignupForm{}), newPageData, signupView)
	m.Post("/kontakt/", binding.Bind(ContactForm{}), newPageData, kontaktViewPost)

	m.NotFound(view404)

	// m.Run()
	log.Println("Server is running...")
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", conf.Port), m)
}

package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type user struct {
	Email string
	Name  string
}

var database = []user{
	{Email: "alexander.davis@example.com", Name: "Alexander Davis"},
	{Email: "alexander.jackson@example.com", Name: "Alexander Jackson"},
	{Email: "avery.williams@example.com", Name: "Avery Williams"},
	{Email: "charlotte.smith@example.com", Name: "Charlotte Smith"},
	{Email: "daniel.miller@example.com", Name: "Daniel Miller"},
	{Email: "ella.smith@example.com", Name: "Ella Smith"},
	{Email: "jacob.white@example.com", Name: "Jacob White"},
	{Email: "james.martinez@example.com", Name: "James Martinez"},
	{Email: "james.miller@example.com", Name: "James Miller"},
	{Email: "jayden.jackson@example.com", Name: "Jayden Jackson"},
	{Email: "liam.robinson@example.com", Name: "Liam Robinson"},
	{Email: "mason.martin@example.com", Name: "Mason Martin"},
	{Email: "matthew.jackson@example.com", Name: "Matthew Jackson"},
	{Email: "mia.smith@example.com", Name: "Mia Smith"},
	{Email: "michael.white@example.com", Name: "Michael White"},
	{Email: "natalie.martin@example.com", Name: "Natalie Martin"},
	{Email: "sofia.garcia@example.com", Name: "Sofia Garcia"},
	{Email: "william.brown@example.com", Name: "William Brown"},
}

type worker struct {
	users []user
	ch    chan user
	name  string
}

func newWorker(users []user, ch chan user, name string) *worker {
	return &worker{users, ch, name}
}

func (w *worker) Find(email string) {
	for _, user := range w.users {
		if strings.Contains(user.Email, email) {
			log.Printf("Worker %s found owner of %s", w.name, user.Email)
			w.ch <- user
		}
	}
}

func main() {
	email := os.Args[1]
	ch := make(chan user)

	log.Printf("Looking for user with email: %s", email)

	go newWorker(database[:6], ch, "#1").Find(email)
	go newWorker(database[6:12], ch, "#2").Find(email)
	go newWorker(database[12:], ch, "#3").Find(email)

	for {
		select {
		case user := <-ch:
			log.Printf("Email address %s is owned by %s", user.Email, user.Name)
		case <-time.After(1 * time.Second):
			return
		}
	}
}

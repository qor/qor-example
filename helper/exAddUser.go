// +build ignore

package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

var (
	email      = flag.String("email", "", "email/login of the user")
	name       = flag.String("name", "", "name of the user")
	clearPassw = flag.String("pass", "", "password for the new user")
	role       = flag.String("role", "default", "the role the user should have")
	cost       = flag.Int("bcryptCost", bcrypt.MinCost, "cost factor for bcrypt")
)

func main() {
	flag.Parse()

	if *email == "" || *clearPassw == "" {
		log.Println("email and password can't be empty.")
		log.Fatalf("Usage: %s -email my@login.net -pass musterMaster -role admin", os.Args[0])
	}

	var u models.User
	u.Email = *email
	u.Name = *name
    u.Role = *role
	var err error
	u.Hashed, err = bcrypt.GenerateFromPassword([]byte(*clearPassw), *cost)
	check(err)
	check(db.DB.Save(&u).Error)
	log.Println("User Added")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

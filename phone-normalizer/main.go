package main

import (
	"fmt"
	"regexp"

	phoneDb "github.com/samueldaviddelacruz/golang-exercises/phone-normalizer/db"
)

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

/*
func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
*/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mypass"
	dbname   = "gophercises_phone"
)

func main() {
	psqlInf := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	must(phoneDb.Reset("postgres", psqlInf, dbname))

	psqlInf = fmt.Sprintf("%s dbname=%s", psqlInf, dbname)

	must(phoneDb.Migrate("postgres", psqlInf))
	db, err := phoneDb.Open("postgres", psqlInf)
	must(err)

	defer db.Close()

	if err := db.Seed(); err != nil {
		panic(err)
	}

	phones, err := db.AllPhones()
	must(err)
	for _, pn := range phones {
		fmt.Printf("Working on... %+v\n", pn)

		normalizedNumber := normalize(pn.Number)
		if normalizedNumber != pn.Number {
			fmt.Println("Update or removing...", normalizedNumber)
			foundPhone, err := db.FindPhone(normalizedNumber)
			must(err)

			if foundPhone != nil {
				//delete number
				must(db.DeletePhone(pn.ID))
			} else {
				pn.Number = normalizedNumber
				must(db.UpdatePhone(&pn))
			}

		} else {
			fmt.Println("No changes required")
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

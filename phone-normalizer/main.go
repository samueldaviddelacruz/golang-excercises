package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", psqlInf)
	must(err)
	err = resetDB(db, dbname)
	must(err)
	db.Close()
	phoneNumbers := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	psqlInf = fmt.Sprintf("%s dbname=%s", psqlInf, dbname)
	db, err = sql.Open("postgres", psqlInf)
	must(err)
	defer db.Close()

	must(createPhoneNumbersTable(db))
	for _, pn := range phoneNumbers {
		_, err := insertPhone(db, pn)

		must(err)

	}

	phones, err := allPhones(db)
	must(err)
	for _, pn := range phones {
		fmt.Printf("Working on... %+v\n", pn)

		normalizedNumber := normalize(pn.number)
		if normalizedNumber != pn.number {
			fmt.Println("Update or removing...", normalizedNumber)

			foundPhone, err := findPhone(db, normalizedNumber)
			must(err)
			if foundPhone != nil {
				//delete number
				must(deletePhone(db, pn.id))
			} else {
				pn.number = normalizedNumber
				must(updatePhone(db, pn))
			}

		} else {
			fmt.Println("No changes required")
		}
	}

	//fmt.Println("id=", id)
}

func updatePhone(db *sql.DB, p phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, p.id, p.number)
	return err
}
func deletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", id).Scan(&id, &number)
	if err != nil {
		return "", err
	}
	return number, nil
}
func findPhone(db *sql.DB, number string) (*phone, error) {
	var p phone
	err := db.QueryRow("SELECT * FROM phone_numbers WHERE value=$1", number).Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err

	}
	return &p, nil
}

type phone struct {
	id     int
	number string
}

func allPhones(db *sql.DB) ([]phone, error) {
	rows, err := db.Query("SELECT id,value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	var result []phone
	for rows.Next() {
		var p phone

		if err := rows.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	return result, nil
}
func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
			CREATE TABLE IF NOT EXISTS phone_numbers (
				id SERIAL,
				value VARCHAR(255)
			)`

	_, err := db.Exec(statement)
	return err
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}

	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}

	return nil
}

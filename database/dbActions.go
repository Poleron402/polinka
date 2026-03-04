package polinkadatabase

import (
	"database/sql"
	"log"
"fmt"
	_ "github.com/mattn/go-sqlite3"
)
var db *sql.DB

type Deck struct {
	ID int64
	Name string
}
func CreateDBIfNotExist() {
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		log.Fatal(err)}
	defer db.Close()

	sqlStatement := `create table if not exists polinka_deck (id integer not null primary key autoincrement,
	name text unique);`

	sqlStatement1 := `create table if not exists polinka_flashcard (id integer not null primary key autoincrement,
	question text,
	answer text,
	hint text);`

	// Execute and check the first statement
	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	// Execute and check the second statement
	_, err = db.Exec(sqlStatement1)
	if err != nil {
		log.Fatal(err)
	}
}
func CreateDeck(name string) error {
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	defer db.Close()
	_, err = db.Exec(`insert into polinka_deck (name) values (?);`, name) 
	if err != nil {
		return err
	}
	return nil
}

func ListDecks() ([]Deck, error) {
	var decks []Deck
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		log.Fatal(err)}
	defer db.Close()
	rows, err := db.Query(`select id, name from polinka_deck`)
	if err != nil {
		log.Fatal("There was an error fetching decks.")
		return decks, err
	}
	for rows.Next() {
		var deck Deck
		if err:=rows.Scan(&deck.ID, &deck.Name); err != nil {
			log.Fatal("There was an error parsing decks to an object.")
		}
		decks = append(decks, deck)
	}
	return decks, nil
}
func DeleteDeck(name string) error {
	fmt.Printf("%s", name)
	db, err := sql.Open("sqlite3", "./polinkadb")
	if err != nil {
		log.Fatal(err)}
	defer db.Close()
	_, err = db.Exec(`delete from polinka_deck where name=(?)`, name)
	if err != nil {
		return err
	}
	return nil
}
// func GetDeck(name string) (Deck, error) {
// 	db, err := sql.Open("sqlite3", "./polinkadb")
// 	if err != nil {
// 		log.Fatal(err)}
// 	defer db.Close()
// 	rows, err := db.Query(`select id, name from polinka_deck`)
// }
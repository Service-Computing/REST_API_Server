package database

import (
	"log"

	"github.com/boltdb/bolt"
)

type MyDB struct {
	db *bolt.DB
}

var myDB *MyDB

func GetDB() *MyDB {
	if myDB == nil {
		myDB = &MyDB{}
		myDB.db, _ = bolt.Open("./database/my.db", 0600, nil)
	}
	return myDB
}

func (myDB *MyDB) InsertUser(username string, password string) {
	db := myDB.db
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		err := b.Put([]byte(username), []byte(password))
		return err
	})

}

func (myDB *MyDB) CheckUserIsExist(username string) bool {
	db := myDB.db
	var length int
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		v := b.Get([]byte(username))
		length = len(v)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return (length != 0)
}

func (myDB *MyDB) CheckPassword(username string, password string) bool {
	db := myDB.db
	var password_saved string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		v := b.Get([]byte(username))
		password_saved = string(v[:])
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return password_saved == password
}

func (myDB *MyDB) QueryPeople(strId string) string {
	return queryString(myDB, strId, "People")

}

func (myDB *MyDB) QueryFilm(strId string) string {
	return queryString(myDB, strId, "Film")

}

func (myDB *MyDB) QueryPlanet(strId string) string {
	return queryString(myDB, strId, "Planet")
}

func (myDB *MyDB) QuerySpecies(strId string) string {
	return queryString(myDB, strId, "Species")
}

func (myDB *MyDB) QueryStarship(strId string) string {
	return queryString(myDB, strId, "Starship")
}

func (myDB *MyDB) QueryVehicle(strId string) string {
	return queryString(myDB, strId, "Vehicle")
}
func queryString(myDB *MyDB, strId string, model string) string {
	db := myDB.db
	var data string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(model))
		v := b.Get([]byte(strId))
		data = string(v)
		return nil
	})

	return data
}

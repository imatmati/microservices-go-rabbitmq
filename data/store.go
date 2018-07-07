package data

import (
	"encoding/json"
	"log"

	"github.com/tidwall/buntdb"
)

const (
	CLOSED = iota
	SUSPENDED
	OPENED
)

const (
	EUR = "EURO"
	USD = "DOLLAR"
)

type Account struct {
	Number   string
	Amount   int
	Currency string
	Status   int
}

var Db *buntdb.DB

func init() {
	var err error
	Db, err = buntdb.Open(":memory:")
	account := Account{"RI5TO9O", 568067, EUR, OPENED}
	if err != nil {
		log.Fatal(err)
	}
	//defer Db.Close()
	err = Db.Update(func(tx *buntdb.Tx) error {

		jsonAccount, err := json.Marshal(account)
		if err != nil {
			return err
		}
		_, _, err = tx.Set("RI5TO9O", string(jsonAccount), nil)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

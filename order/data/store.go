package data

import (
	"encoding/json"
	l "payment/utils/language"

	"github.com/tidwall/buntdb"
)

const (
	EUR = "EURO"
	USD = "DOLLAR"
)

type Account struct {
	Number   string
	Amount   int
	Currency string
}

var Db *buntdb.DB

func init() {
	var err error
	Db, err = buntdb.Open(":memory:")
	l.PanicIf(err)
	account := Account{"RI5TO9O", 568067, EUR}

	err = Db.Update(func(tx *buntdb.Tx) error {

		jsonAccount, err := json.Marshal(account)
		l.PanicIf(err)
		_, _, err = tx.Set("RI5TO9O", string(jsonAccount), nil)
		return err
	})
	l.PanicIf(err)
}

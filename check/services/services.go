package services

import (
	"encoding/json"
	"payment/check/data"
	l "payment/utils/language"

	"github.com/tidwall/buntdb"
)

func CheckAccount(number string) (check bool) {
	check = false
	data.Db.View(func(tx *buntdb.Tx) error {
		accountJson, err := tx.Get(number)
		l.PanicIf(err)
		account := data.Account{}

		json.Unmarshal([]byte(accountJson), &account)
		check = account.Status != data.CLOSED && account.Status != data.SUSPENDED
		return nil
	})
	return
}

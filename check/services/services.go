package services

import (
	"account/check/data"
	"encoding/json"
	"fmt"

	"github.com/tidwall/buntdb"
)

func CheckAccount(number string) (check bool) {
	check = false
	data.Db.View(func(tx *buntdb.Tx) error {
		accountJson, err := tx.Get(number)
		if err != nil {
			return fmt.Errorf("account %s not found", number)
		}

		account := data.Account{}

		json.Unmarshal([]byte(accountJson), &account)
		check = account.Status != data.CLOSED && account.Status != data.SUSPENDED
		return nil
	})
	return
}

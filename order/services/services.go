package services

import (
	"encoding/json"
	"fmt"
	"payment/logger"
	"payment/order/data"
	"payment/order/messaging"
	l "payment/utils/language"

	"github.com/tidwall/buntdb"
)

func CheckAccount(account string) bool {
	return messaging.CheckAccount(account)
}

func UpdateAccount(number string, amount int, currency string) error {
	return data.Db.Update(func(tx *buntdb.Tx) error {
		accountJson, err := tx.Get(number)
		logger.Logger.Printf("Check for account %v\n", accountJson)
		l.PanicIf(err)
		account := data.Account{}
		json.Unmarshal([]byte(accountJson), &account)
		if account.Amount < amount {
			return fmt.Errorf("Insufficient provision on number %s", number)
		}
		account.Amount -= amount
		jsonAccount, err := json.Marshal(account)
		l.PanicIf(err)
		_, _, err = tx.Set("RI5TO9O", string(jsonAccount), nil)
		return err
	})
}

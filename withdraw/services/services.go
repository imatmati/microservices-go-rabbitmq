package services

import (
	"account/logger"
	l "account/utils/language"
	"account/withdraw/data"
	"account/withdraw/messaging"
	"encoding/json"
	"fmt"

	"github.com/tidwall/buntdb"
)

func CheckAccount(account string) bool {
	return messaging.CheckAccount(account)
}

func UpdateAccount(number string, amount int, currency string) error {
	return data.Db.Update(func(tx *buntdb.Tx) error {
		accountJson, err := tx.Get(number)
		logger.Logger.Printf("Withdraw for account %v\n", accountJson)
		l.PanicIf(err)
		account := data.Account{}
		json.Unmarshal([]byte(accountJson), &account)
		if account.Currency != currency {
			return fmt.Errorf("Different currencies between account number %s and order : %s != %s", number, account.Currency, currency)
		}
		if account.Amount < amount {
			return fmt.Errorf("Insufficient provision on account number %s", number)
		}
		account.Amount -= amount
		jsonAccount, err := json.Marshal(account)
		if err != nil {
			return err
		}
		_, _, err = tx.Set("RI5TO9O", string(jsonAccount), nil)
		return err
	})
}

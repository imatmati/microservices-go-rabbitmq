package services

import (
	"account/logger"
	"account/withdraw/data"
	"account/withdraw/messaging"
	"encoding/json"
	"fmt"

	"github.com/tidwall/buntdb"
)

func CheckAccount(account string) bool {
	return messaging.CheckAccount(account)
}

func UpdateAccount(number string, amount int, currency string) (account_amount int, err error) {

	err = data.Db.Update(func(tx *buntdb.Tx) error {
		accountJson, err := tx.Get(number)
		logger.Logger.Printf("Withdraw for account %v\n", accountJson)
		if err != nil {
			return err
		}
		account := data.Account{}
		json.Unmarshal([]byte(accountJson), &account)
		if account.Currency != currency {
			return fmt.Errorf("Different currencies between account number %s and order : %s != %s", number, account.Currency, currency)
		}
		if account.Amount < amount {
			return fmt.Errorf("Insufficient provision on account number %s", number)
		}
		account.Amount -= amount
		account_amount = account.Amount
		jsonAccount, err := json.Marshal(account)
		if err != nil {
			return err
		}
		_, _, err = tx.Set(number, string(jsonAccount), nil)
		return err
	})

	return
}

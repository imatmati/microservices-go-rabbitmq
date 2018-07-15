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

// CheckAccount asks check account microservice about account availability ( true = go on, false = account closed, not found, blocked, etc.)
func CheckAccount(account string) bool {
	return messaging.CheckAccount(account)
}

// UpdateAccount withdraw amount of 1/100 from account and updates database.
func UpdateAccount(number string, amount int, currency string) int {
	account_amount := -1
	err := data.Db.Update(func(tx *buntdb.Tx) error {
		account_db, err := tx.Get(number)
		l.PanicIf(err, "Account ", number, " not found")
		logger.Logger.Printf("Withdraw of %d/100 of %s for account %v\n", amount, currency, account_db)

		account := data.Account{}
		json.Unmarshal([]byte(account_db), &account)

		if account.Amount < amount {
			return fmt.Errorf("Insufficient provision on account  %s", number)
		}
		if account.Currency != currency {
			return fmt.Errorf("Different currencies between account %s and order : %s != %s", number, account.Currency, currency)
		}
		account.Amount -= amount
		account_amount = account.Amount
		account_after_withdraw, err := json.Marshal(account)
		if err != nil {
			return err
		}
		_, _, err = tx.Set(number, string(account_after_withdraw), nil)
		logger.Logger.Printf("Account after withdraw %v\n", string(account_after_withdraw))
		return err
	})
	l.PanicIf(err)
	return account_amount
}

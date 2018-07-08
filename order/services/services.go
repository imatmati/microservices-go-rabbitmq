package services

import (
	"payment/order/messaging"
)

func CheckAccount(account string) bool {
	return messaging.CheckAccount(account)
}

package factory

import (
	"errors"
	"strconv"
)

const (
	Cash = iota+1
	DebitCard
)

// 支付方式
type PaymentMethod interface {
	Pay(amount int) string
}

type CashPM struct {

}

func (c *CashPM) Pay(amount int) string {
	return "you pay" + strconv.Itoa(amount) + "by CASH payment"
}

type DebitCardPM struct {

}

func (d *DebitCardPM) Pay(amount int) string {
	return "you pay" + strconv.Itoa(amount) + "by DEBIT CARD payment"
}

func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case Cash:
		return &CashPM{}, nil
	case DebitCard:
		return &DebitCardPM{}, nil
	default:
		return nil, errors.New("payment method not recognized")
	}
}

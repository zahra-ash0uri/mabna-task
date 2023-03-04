package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	ID           int             `json:"id"`
	InstrumentID int             `json:"instrument_id"`
	DateEn       time.Time       `json:"date_en"`
	Open         decimal.Decimal `json:"open"`
	High         decimal.Decimal `json:"high"`
	Low          decimal.Decimal `json:"low"`
	Close        decimal.Decimal `json:"close"`
}

package currency

import (
	"errors"
	"fmt"
)

func (c CurrencyStruct) ToString() string {
	strValue := fmt.Sprintf("%.2f", float64(c.ValueInCents)/100)
	strCurrency := ""

	switch c.CurrencyCode {
	case BRL:
		strCurrency = "R$"
	case USD:
		strCurrency = "$"
	case EUR:
		strCurrency = "€"
	default:
		strCurrency = "?"
	}

	return fmt.Sprintf("%s %s", strCurrency, strValue)
}

func (c CurrencyStruct) ParseTo(newCurrency CurrencyCode) (CurrencyStruct, error) {
	multiplier := 0.0

	currency := c.CurrencyCode

	if currency == newCurrency {
		return c, nil // No conversion needed
	}

	switch newCurrency {
	case BRL:
		if currency == USD {
			multiplier = 5.21 // Example conversion rate from USD to BRL
		}
		if currency == EUR {
			multiplier = 5.85 // Example conversion rate from EUR to BRL
		}
	case USD:
		if currency == BRL {
			multiplier = 0.19
		}
		if currency == EUR {
			multiplier = 1.12
		}
	case EUR:
		if currency == BRL {
			multiplier = 0.17
		}
		if currency == USD {
			multiplier = 0.89
		}
	default:
		return CurrencyStruct{}, errors.New("unsupported currency")
	}

	return CurrencyStruct{
		ValueInCents: int64(float64(c.ValueInCents) * multiplier),
		CurrencyCode: newCurrency,
	}, nil
}

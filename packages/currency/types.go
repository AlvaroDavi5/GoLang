package currency

type CurrencyCode string

const (
	USD CurrencyCode = "USD"
	EUR CurrencyCode = "EUR"
	BRL CurrencyCode = "BRL"
)

type CurrencyStruct struct {
	ValueInCents int64
	CurrencyCode CurrencyCode
}

type CurrencyInterface interface {
	ToString() string
	ParseTo(CurrencyCode) (CurrencyStruct, error)
}

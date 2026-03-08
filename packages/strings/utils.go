package strings

import (
	"rsc.io/quote"
)

func GetHelloWorld() string {
	helloWorld := quote.Hello()
	return helloWorld
}

package actions

import (
	"fmt"
	
	"rsc.io/quote"
)

func Quote() {
	fmt.Println(quote.Hello())
}

func DoAction() {
	fmt.Println("Action function called")
}

func DoOtherAction() {
	fmt.Println("Other action function called")
}

func MoreActions() {
	fmt.Println("More actions function called")
}
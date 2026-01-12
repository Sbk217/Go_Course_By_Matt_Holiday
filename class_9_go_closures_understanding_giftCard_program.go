package main

import (
	"fmt"
)

//in the below we see the concept of GO closures as closures are basically closures in the memory reserved by a enviorment pointer
//in this case (useGiftCard1 and useGiftCard2) the enviorment pointer are different but function pointer are the same

func activateGiftCard() func(int) (int, string) {
	//A function with no input parameters having return type as a function which takes 1 int parameter and return 2 int and string values
	amount := 100

	debitFunc := func(debitAmount int) (int, string) {
		amount -= debitAmount
		message := fmt.Sprintf("Amount after deduction: %d", amount)
		return amount, message
	}
	return debitFunc
}

func main() {
	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()
	fmt.Println(useGiftCard1(30))
	fmt.Println(useGiftCard2(34))
}

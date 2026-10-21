package main

import "fmt"

const (
	USD = 90.0  
	EUR = 100.0 
)

func main() {
	var rub float64

	fmt.Println("Введите сумму в рублях:")
	fmt.Scan(&rub)

	usd := rub / USD
	eur := rub / EUR
	usdToEur := usd * 90/ 100


	fmt.Println("Рубли в доллары:", int(usd), "USD")
	fmt.Println("Рубли в евро:", int(eur), "EUR")
	fmt.Println("Доллары в евро:", int(usdToEur), "EUR")
}
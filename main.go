package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	rates, err := LoadRates()
	if err != nil {
		log.Fatal(err)
	}

	var value float64
	var currency string

	fmt.Print("Digite o valor: ")
	if _, err := fmt.Scan(&value); err != nil {
		log.Fatal("valor inválido: ", err)
	}

	fmt.Print("Digite a moeda: ")
	if _, err := fmt.Scan(&currency); err != nil {
		log.Fatal("moeda inválida: ", err)
	}

	currency = strings.ToUpper(currency)

	result, err := Convert(value, currency, rates)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Valor convertido: %.4f %s\n", result, currency)
}

func Convert(value float64, currency string, rates RatesData) (float64, error) {
	rate, ok := rates.Rates[currency]
	if !ok {
		return 0, fmt.Errorf("moeda não encontrada: %s", currency)
	}

	return value * rate, nil
}

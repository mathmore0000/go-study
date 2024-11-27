package exchangeRateCli

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type exchangeRate struct {
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
	Country      string `json:"country"`
}

func main() {
	var exchangeRates []exchangeRate

	content, err := os.ReadFile("supported_currencies.json")
	if err != nil {
		panic(err)
	}

	// Fill the instance from the JSON file content
	err = json.Unmarshal(content, &exchangeRates)

	// Check if is there any error while filling the instance
	if err != nil {
		panic(err)
	}

	// fmt.Println(content)
	fmt.Printf("%+v\n", exchangeRates[0].CurrencyCode)

	// fmt.Println(getExchangeURL("USD"))

	//fmt.Println("Which currency do you wish do exchange?")

	//fmt.Println("How much of %v do you wish do exchange?", currency)

	//fmt.Println("To which currency do you wish do %v %v?", ammount, currency)

	//fmt.Println("%v %v = %v %v", currency, ammount, result.currency, result.ammount)
}

func getExchangeURL(currency string) (url string) {
	var API_KEY string = os.Getenv("API_KEY")

	return fmt.Sprintf("https://v6.exchangerate-api.com/v6/%v/latest/%v", API_KEY, currency)
}

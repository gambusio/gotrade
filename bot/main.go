package main

import (
    "flag"
    "fmt"
    "log"
    "github.com/gotrade/qtrade"
)

func main() {
    fiatCurrency := flag.String(
      "fiat", "USD", "The name of the fiat currency you would like to know the price of your crypto in",
    )

    nameOfCrypto := flag.String(
      "crypto", "ARMS", "Input the name of the CryptoCurrency you would like to know the price of",
    )
    flag.Parse()

    crypto, err := qtrade.FetchCrypto(*fiatCurrency, *nameOfCrypto)
    if err != nil {
        log.Println(err)
      }

  fmt.Println(crypto)
}

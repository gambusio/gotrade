package qtrade

import (
   "encoding/json"
   "fmt"
   "log"
   "net/http"
)

// Cryptoresponse is exported, it models the data we receive.
type Cryptoresponse []struct {
   Name              string    `json:"name"`
   Price             string    `json:"price"`
   Rank              string    `json:"rank"`
   High              string    `json:"high"`
   CirculatingSupply string    `json:"circulating_supply"`
}

//TextOutput is exported,it formats the data to plain text.
func (c Cryptoresponse) TextOutput() string {
p := fmt.Sprintf(
  "Name: %s\nPrice : %s\nRank: %s\nHigh: %s\nCirculatingSupply: %s\n",
  c[0].Name, c[0].Price, c[0].Rank, c[0].High, c[0].CirculatingSupply)
   return p
}

//Fetch is exported ...
func FetchCrypto(fiat string , crypto string) (string, error) {
//Build The URL string
   URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids="+crypto+"&convert="+fiat
//We make HTTP request using the Get function
   resp, err := http.Get(URL)
   if err != nil {
      log.Fatal("ooopsss an error occurred, please try again")
   }
   defer resp.Body.Close()
//Create a variable of the same type as our model
   var cResp Cryptoresponse
//Decode the data
   if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
      log.Fatal("ooopsss! an error occurred, please try again")
   }
//Invoke the text output function & return it with nil as the error value
   return cResp.TextOutput(), nil
}

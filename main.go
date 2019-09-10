package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	trakt "github.com/Iliyass/trakt/Trakt"
)

func main() {

	from := time.Now()
	ex := trakt.Tag{Name: "Exersicing", CreatedAt: time.Now()}
	run := trakt.Trakt{Text: "Run 5k on 29m", CreatedAt: time.Now(), Tags: []trakt.Tag{ex}}
	to := time.Now()

	data, _ := json.Marshal(run)
	fmt.Println(trakt.AddTrakt(data))
	fromStr := strconv.FormatInt(from.Unix(), 10)
	toStr := strconv.FormatInt(to.Unix(), 10)
	fmt.Println(string(trakt.GetTraktsByDate(fromStr, toStr)))
}

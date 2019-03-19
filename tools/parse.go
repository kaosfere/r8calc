package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"astuart.co/goq"
)

type test struct {
	Name []string `goquery:"table.table tbody tr h4"`
	Vals []string `goquery:"table.table tbody tr p"`
}

func main() {
	res, err := http.Get("https://www.thedepotserver.com/reference/rollingstock/index.php?f=all")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var tx test
	err = goq.NewDecoder(res.Body).Decode(&tx)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("weights := map[string]float64{")
	for index, raw := range tx.Vals {
		fmt.Printf("\"%s\": %s,\n", tx.Name[index], strings.Fields(strings.Split(raw, "\n")[2])[2])
	}
	fmt.Println("}")
}

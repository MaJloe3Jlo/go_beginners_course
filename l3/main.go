package main

import (
	"fmt"
)

func main() {
	var users [3]string = [3]string{"Dima", "Max", "Nick"}

	transaction := make(map[int]map[string]string)
	transaction[1] = map[string]string{
		"User":   "Dima",
		"Status": "paid",
		"Sum":    "100RUR",
	}
	transaction[2] = map[string]string{
		"User":   "Max",
		"Status": "cancel",
		"Sum":    "2000RUR",
	}
	transaction[3] = map[string]string{
		"User":   "Dima",
		"Status": "cancel",
		"Sum":    "10000000000RUR",
	}
	transaction[4] = map[string]string{
		"User":   "Nick",
		"Status": "paid",
		"Sum":    "100RUR",
	}
	transaction[5] = map[string]string{
		"User":   "Max",
		"Status": "paid",
		"Sum":    "4500RUR",
	}
	transaction[6] = map[string]string{
		"User":   "Nick",
		"Status": "cancel",
		"Sum":    "100000RUR",
	}

	fmt.Println(users)
	fmt.Printf("Unsorted trxs: %v \n", transaction)

	var sortedTrxs []map[string]string

	for _, value := range users {
		for i := 0; i <= len(transaction); i++ {
			if trx, ok := transaction[i]; ok {
				if trx["User"] == value {
					sortedTrxs = append(sortedTrxs, transaction[i])
				}
			}
		}
	}

	for _, value := range sortedTrxs {
		fmt.Printf("User: %s, Status: %s, Sum: %s \n", value["User"], value["Status"], value["Sum"])
	}
}

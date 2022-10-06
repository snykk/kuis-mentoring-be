package main

import "fmt"

func generateWords(num int) string {
	if num == 0 {
		return "nol"
	}

	numHelper := map[int]string{
		1: "satu",
		2: "dua",
		3: "tiga",
		4: "empat",
		5: "lima",
		6: "enam",
		7: "tujuh",
		8: "depalan",
		9: "sembilan",
	}

	if num < 10 {
		return numHelper[num]
	} else if num == 100 {
		return "seratus"
	} else if num%10 == 0 {
		return numHelper[num/10] + " puluh"
	} else if num < 20 {
		if num == 11 {
			return "sebelas"
		}

		return numHelper[num%10] + " belas"
	}

	return numHelper[num/10] + " puluh " + numHelper[num%10]

}

func main() {

	var input int = 30
	var result string

	for true {
		fmt.Scan(&input)
		result = generateWords(input)
		fmt.Println(result)
	}

	// result = generateWords(input)
	// fmt.Println(result)

}

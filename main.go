package main

import "fmt"

func main()  {
	var dia, mes int
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	signo := mes*mes*10 + dia
	switch {
	case signo <= 58:
		fmt.Println("acuario")
	case signo <= 110:
		fmt.Println("piscis")
	case signo <= 179:
		fmt.Println("aries")
	case signo <= 270:
		fmt.Println("tauro")
	case signo <= 380:
		fmt.Println("geminis")
	case signo <= 512:
		fmt.Println("cancer")
	case signo <= 662:
		fmt.Println("leo")
	case signo <= 832:
		fmt.Println("virgo")
	case signo <= 1022:
		fmt.Println("libra")
	case signo <= 1231:
		fmt.Println("escorpio")
	case signo <= 1461:
		fmt.Println("sagitario")
	case signo >= 1462 || signo <= 29:
		fmt.Println("capricornio")
	}
}
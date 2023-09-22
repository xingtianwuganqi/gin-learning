package handler

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func IpNetworking() {
	response, err := http.Get("http://www.ip3366.net/free/?stype=1")
	if err != nil {
		log.Println("Ip Network err")
		return
	}
	defer response.Body.Close()
	//body, err := io.ReadAll(response.Body)
	//if err != nil {
	//	log.Println("response body err")
	//	return
	//}
	if doc, err := goquery.NewDocumentFromReader(response.Body); err == nil {
		doc.Find("")
	}
}

func main() {
	IpNetworking()
}

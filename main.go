package main

import (
	"math/rand"
	"time"
	"strings"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/gocolly/colly"
)


func fetchAndSave(url string, dir string, filename string)  {

	response, e := http.Get(url)
	if e != nil {
			log.Fatal(e)
	}
	defer response.Body.Close()

	file, err := os.Create(dir + filename + ".webm")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
			log.Fatal(err)
	}
	fmt.Println(filename + " has been saved.")
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {

	c := colly.NewCollector()
	host := "https://2ch.hk"
	dir := "/Users/dimashulhin/2ch/webm/"

	//os.Mkdir(dir, os.FileMode(0522))

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, ".webm") {
			fmt.Println(link)
			fetchAndSave(host + link, dir, RandStringRunes(20))
	
		}// true
	
	})

	c.Visit("https://2ch.hk/b")


}
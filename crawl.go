package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func goGet() {
	var (
		row  []string
		rows [][]string
		err  error
	)

	fmt.Print("Enter Date String: ") //Print function is used to display output in same line
	var date string
	fmt.Scanln(&date)

	fmt.Print("Enter Month String: ") //Print function is used to display output in same line
	var month string
	fmt.Scanln(&month)

	fmt.Print("Enter Year String: ") //Print function is used to display output in same line
	var year string
	fmt.Scanln(&year)

	url := "https://xsmn.me/xsmb-" + date + "-" + month + "-" + year + "-ket-qua-xo-so-mien-bac-ngay-" + date + "-" + month + "-" + year + ".html"
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println(err)
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		log.Fatalln(err)
	}
	v := strings.NewReader(string(body))
	doc, err := goquery.NewDocumentFromReader(v)
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	doc.Find("table").Each(func(indextr int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
			})
			rows = append(rows, row)
			row = nil
		})
	})

	fmt.Println(rows[4:14])
}

type Prize struct {
	Prize  string
	Number int
}

//func GetClient() *mongo.Client {
//	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//	client, err := mongo.NewClient(clientOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = client.Connect(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return client
//}

func main() {
	goGet()


}

//1. git la gi
//2 quy trinh de git thanh cong commit,push
//3 branch
//4 tao 1 repository push 2 branch 1 dev 1 master
//5 xu li cac conflict
//output = tao dc github push code vao dev. master contain readme
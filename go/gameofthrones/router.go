package main

import (
	"fmt"
	"gameofthrones/models"
	"gameofthrones/resttemplate"
	"gameofthrones/urls"
)

func main() {
	root := models.Root{}
	resttemplate.MakeGetCall(urls.Root, &root)
	fmt.Printf("%+v\n", root)

	book := models.Book{}
	url := fmt.Sprintf("%s/%d", urls.Books, 1)
	resttemplate.MakeGetCall(url, &book)
	fmt.Println(book)
	//fmt.Printf("%+v\n", book)

	character := models.Character{}
	url = fmt.Sprintf("%s/%d", urls.Characters, 1)
	resttemplate.MakeGetCall(urls.Root, &character)
	//fmt.Printf("%+v\n", character)

	house := models.House{}
	url = fmt.Sprintf("%s/%d", urls.Houses, 1)
	resttemplate.MakeGetCall(urls.Root, &house)
	//fmt.Printf("%+v\n", house)
}

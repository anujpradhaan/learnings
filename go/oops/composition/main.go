package main

import "fmt"

/*
	This example we are going to leanr about
	composition of structs in golang. We should prefer composition over inheritance
*/

type author struct {
	fname string
	lname string
	age   int
}

func (a author) getFullname() string {
	return fmt.Sprintf("%s %s", a.fname, a.lname)
}

// Keeping a of composite structs inside another struct
type post struct {
	title   string
	content string
	author  author
}

func (p post) details() {
	fmt.Printf("Post title %s\n", p.title)
	fmt.Printf("Author Name %s\n", p.author.getFullname())
}

// Keeping a slice of composite structs
type website struct {
	posts []post
}

func (w website) websiteDetails() {
	for _, v := range w.posts {
		v.details()
	}
}

func main() {
	a1 := author{
		fname: "anuj",
		lname: "kumar",
	}

	a2 := author{
		fname: "ashvik",
		lname: "kumar",
	}

	p1 := post{
		title:   "How to write in golang",
		content: "Content of the blog post of golang",
		author:  a1,
	}

	p2 := post{
		title:   "How to write in java",
		content: "Content of the blog post of java",
		author:  a2,
	}
	w := website{
		posts: []post{p1, p2},
	}
	w.websiteDetails()
}

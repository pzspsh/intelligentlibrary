/*
@File   : main.go
@Author : pan
@Time   : 2024-01-08 12:22:25
*/
package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (p blogPost) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	blogPosts []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
}

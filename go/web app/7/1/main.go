package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	s1, err := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	//s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	//s1.ExecuteTemplate(os.Stdout, "footer", nil)
	//fmt.Println()
	s1.Execute(os.Stdout, nil)
}

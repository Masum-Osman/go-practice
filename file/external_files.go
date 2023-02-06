package file

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func WritingXML() {

	fmt.Println("WHERE?")
	note := notes{
		To:      "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    "Meeting at 5pm!",
	}

	file, _ := xml.MarshalIndent(note, "", " ")
	fmt.Println("IN HERE?")
	_ = ioutil.WriteFile("notes1.xml", file, 0644)
}

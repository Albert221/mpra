package domain

import "encoding/xml"

type Package struct {
	XMLName  xml.Name `xml:"opakowanie"`
	Size     string   `xml:"wielkosc,attr"`
	SizeUnit string   `xml:"jednostkaWielkosci,attr"`
	Ean      string   `xml:"kodEAN,attr"`
}

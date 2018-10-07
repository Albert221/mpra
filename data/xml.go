package data

import "encoding/xml"

type MedicalProducts struct {
	XMLName  xml.Name          `xml:"produktyLecznicze"`
	AsOfDay  string            `xml:"stanNaDzien,attr"`
	Children []*MedicalProduct `xml:"produktLeczniczy"`
}

type MedicalProduct struct {
	XMLName           xml.Name   `xml:"produktLeczniczy"`
	ProductName       string     `xml:"nazwaProduktu,attr"`
	Kind              string     `xml:"rodzajPreparatu,attr"`
	CommonProductName string     `xml:"nazwaPowszechnieStosowana,attr"`
	Strength          string     `xml:"moc,attr"`
	Form              string     `xml:"postac,attr"`
	ResponsibleEntity string     `xml:"podmiotOdpowiedzielny,attr"`
	ActiveSubstances  []string   `xml:"substancjeCzynne>substancjaCzynna"`
	Packages          []*Package `xml:"opakowania>opakowanie"`
}

type Package struct {
	XMLName  xml.Name `xml:"opakowanie"`
	Size     string   `xml:"wielkosc,attr"`
	SizeUnit string   `xml:"jednostkaWielkosci,attr"`
	Ean      string   `xml:"kodEAN,attr"`
}

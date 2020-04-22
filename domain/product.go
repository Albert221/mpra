package domain

import "encoding/xml"

type Product struct {
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

type root struct {
	XMLName  xml.Name   `xml:"produktyLecznicze"`
	AsOfDay  string     `xml:"stanNaDzien,attr"`
	Children []*Product `xml:"produktLeczniczy"`
}

func UnmarshallProducts(data []byte) ([]*Product, error) {
	var products root
	err := xml.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}

	return products.Children, err
}

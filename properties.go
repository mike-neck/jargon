package main

import "encoding/xml"

type Properties struct {
	Property map[string]string
}

func (p *Properties) Get(name string) string {
	return p.Property[name]
}

func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p.Property = make(map[string]string)
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch se := token.(type) {
		case xml.StartElement:
			var value string
			if err := d.DecodeElement(&value, &se); err != nil {
				return err
			}
			p.Property[se.Name.Local] = value
		case xml.EndElement:
			if se == start.End() {
				return nil
			}
		}
	}
}

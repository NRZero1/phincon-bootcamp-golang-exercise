package domain

import "encoding/xml"

type AddRequestJson struct {
	IntA int `json:"intA"`
	IntB int `json:"intB"`
}

type AddRequest struct {
    XMLName xml.Name `xml:"Add"`
    Xmlns   string   `xml:"xmlns,attr"`
    IntA    int      `xml:"intA"`
    IntB    int      `xml:"intB"`
}

type Body struct {
    XMLName xml.Name `xml:"soap:Body"`
    Add     AddRequest
}

type Envelope struct {
    XMLName xml.Name `xml:"soap:Envelope"`
    XmlnsSoap string `xml:"xmlns:soap,attr"`
    XmlnsXsi string `xml:"xmlns:xsi,attr"`
    XmlnsXsd string `xml:"xmlns:xsd,attr"`
    Body     Body
}

type EnvelopeRes struct {
	XMLName xml.Name `xml:"Envelope"`
    Body     struct {
		XMLName xml.Name `xml:"Body"`
    	Add     AddRes
	}
}

type AddRes struct {
    XMLName   xml.Name `xml:"AddResponse"`
    AddResult int      `xml:"AddResult"`
}

type AddResponseJson struct {
	Result int `json:"result"`
}
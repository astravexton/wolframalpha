package wolframalpha

import "encoding/xml"

// Queryresult ...
type Queryresult struct {
	XMLName       xml.Name `xml:"queryresult"`
	Text          string   `xml:",chardata"`
	Success       bool     `xml:"success,attr"`
	HasError      bool     `xml:"error,attr"`
	Error         string   `xml:"error>msg"`
	Numpods       string   `xml:"numpods,attr"`
	Datatypes     string   `xml:"datatypes,attr"`
	Timedout      string   `xml:"timedout,attr"`
	Timedoutpods  string   `xml:"timedoutpods,attr"`
	Timing        string   `xml:"timing,attr"`
	Parsetiming   string   `xml:"parsetiming,attr"`
	Parsetimedout string   `xml:"parsetimedout,attr"`
	Recalculate   string   `xml:"recalculate,attr"`
	ID            string   `xml:"id,attr"`
	Host          string   `xml:"host,attr"`
	Server        string   `xml:"server,attr"`
	Related       string   `xml:"related,attr"`
	Version       string   `xml:"version,attr"`
	Pods          []struct {
		Text       string `xml:",chardata"`
		Title      string `xml:"title,attr"`
		Scanner    string `xml:"scanner,attr"`
		ID         string `xml:"id,attr"`
		Position   string `xml:"position,attr"`
		Error      string `xml:"error,attr"`
		Numsubpods string `xml:"numsubpods,attr"`
		Primary    string `xml:"primary,attr"`
		SubPods    []struct {
			Text      string `xml:",chardata"`
			Title     string `xml:"title,attr"`
			Primary   string `xml:"primary,attr"`
			Plaintext string `xml:"plaintext"`
		} `xml:"subpod"`
		States struct {
			Text  string `xml:",chardata"`
			Count string `xml:"count,attr"`
			State struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Input string `xml:"input,attr"`
			} `xml:"state"`
		} `xml:"states"`
	} `xml:"pod"`
	Assumptions struct {
		Text       string `xml:",chardata"`
		Count      string `xml:"count,attr"`
		Assumption struct {
			Text     string `xml:",chardata"`
			Type     string `xml:"type,attr"`
			Word     string `xml:"word,attr"`
			Template string `xml:"template,attr"`
			Count    string `xml:"count,attr"`
			Value    []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Desc  string `xml:"desc,attr"`
				Input string `xml:"input,attr"`
			} `xml:"value"`
		} `xml:"assumption"`
	} `xml:"assumptions"`
}

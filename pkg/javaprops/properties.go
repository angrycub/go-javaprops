package javaprops

import "encoding/json"

// Entry holds a single property key/value combination
// <entry key="myKey">the value</entry>
type Entry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:",chardata"`
}

// Entries is a collection object since the properties
// file has a lot of them all beneath the root.
type Entries []Entry

// AsProps converts Entries ([]Entry) to Props (map[string]string)
func (e Entries) AsProps() Props {
	out := make(Props, len(e))
	for _, v := range e {
		out[v.Key] = v.Value
	}
	return out
}

// Props provide a place to add method receivers to a golang map[string]string
// for formatting
type Props map[string]string

// Renders the Props as a pretty JSON string
func (p Props) AsJSON() string { return p.s(json.MarshalIndent(p, "", "  ")) }

// s strips the error from the marshaling and makes the
// byte slice a string
func (_ Props) s(b []byte, _ error) string { return string(b) }

// JVMPropertyFile represents the top level of the properties XML document
type JVMPropertyFile struct {
	Comment string  `xml:"comment"`
	Entries Entries `xml:"entry"`
}

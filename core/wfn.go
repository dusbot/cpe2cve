package core

import (
	"fmt"
	"strings"
)

var KnownParts = map[string]string{
	"a": "application",
	"o": "operating system",
	"h": "hardware",
}

const (
	Any = ""
	NA  = "-"
)

const (
	uriPrefix = "cpe:/"
	fsbPrefix = "cpe:2.3:"
)

var parsers = map[string]func(s string) (*Attributes, error){
	uriPrefix: UnbindURI,
	fsbPrefix: UnbindFmtString,
}

func Parse(s string) (*Attributes, error) {
	for prefix, parserFunc := range parsers {
		if strings.HasPrefix(s, prefix) {
			return parserFunc(s)
		}
	}
	return nil, fmt.Errorf("wfn: unsupported format %q", s)
}

type Attributes struct {
	Part      string
	Vendor    string
	Product   string
	Version   string
	Update    string
	Edition   string
	SWEdition string
	TargetSW  string
	TargetHW  string
	Other     string
	Language  string
}

func NewAttributesWithNA() *Attributes {
	return newAttributes(NA)
}

func NewAttributesWithAny() *Attributes {
	return newAttributes(Any)
}

func newAttributes(defaultValue string) *Attributes {
	return &Attributes{
		Part:      defaultValue,
		Vendor:    defaultValue,
		Product:   defaultValue,
		Version:   defaultValue,
		Update:    defaultValue,
		Edition:   defaultValue,
		SWEdition: defaultValue,
		TargetSW:  defaultValue,
		TargetHW:  defaultValue,
		Other:     defaultValue,
		Language:  defaultValue,
	}
}

func WFNize(s string) (string, error) {
	const allowedPunct = "-!\"#$%&'()+,./:;<=>@[]^`{|}!~"

	in := strings.Replace(s, " ", "_", -1)
	buf := make([]byte, 0, len(in))

	for n, c := range in {
		c := byte(c)
		if c >= 'A' && c <= 'Z' ||
			c >= 'a' && c <= 'z' ||
			c >= '0' && c <= '9' ||
			c == '_' ||
			strings.IndexByte(allowedPunct, c) != -1 {
			buf = append(buf, c)
		}

		if c == '*' || c == '?' {
			if n == 0 || in[n-1] != '\\' {
				buf = append(buf, '\\')
			}
			buf = append(buf, c)
		}
	}

	s, _, err := addSlashesAt(string(buf), 0)
	return s, err
}

func (a Attributes) String() string {
	parts := make([]string, 0, 11)

	parts = append(parts, keyValueString("part", a.Part))
	parts = append(parts, keyValueString("vendor", a.Vendor))
	parts = append(parts, keyValueString("product", a.Product))
	parts = append(parts, keyValueString("version", a.Version))
	parts = append(parts, keyValueString("update", a.Update))
	parts = append(parts, keyValueString("edition", a.Edition))

	if a.SWEdition != Any || a.TargetHW != Any || a.TargetSW != Any || a.Other != Any {
		parts = append(parts, keyValueString("sw_edition", a.SWEdition))
		parts = append(parts, keyValueString("target_sw", a.TargetSW))
		parts = append(parts, keyValueString("target_hw", a.TargetHW))
		parts = append(parts, keyValueString("other", a.Other))
	}

	parts = append(parts, keyValueString("language", a.Language))
	return fmt.Sprintf("wfn:[%s]", strings.Join(parts, ","))
}

func keyValueString(k, v string) string {
	switch v {
	case Any:
		return fmt.Sprintf("%s=ANY", k)
	case NA:
		return fmt.Sprintf("%s=NA", k)
	default:
		return fmt.Sprintf("%s=\"%s\"", k, v)
	}
}

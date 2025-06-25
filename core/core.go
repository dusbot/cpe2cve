package core

import (
	"encoding/json"
	"fmt"
	"time"

	"resty.dev/v3"
)

const baseURL = "https://services.nvd.nist.gov/rest/json/cves/2.0"

func CPE2CVE(cpe_ string) (cves []string) {
	attr, err := Parse(cpe_)
	if err != nil {
		return
	}
	return cpe23_2cve(attr.ToCPE23String())
}

func cpe23_2cve(cpe23 string) (cves []string) {
	client := resty.New()
	defer client.Close()
	url := fmt.Sprintf("%s?cpeName=%s", baseURL, cpe23)
	resp, err := client.R().SetTimeout(30*time.Second).Get(url)
	if err != nil {
		return
	}
	var m map[string]interface{}
	if err := json.Unmarshal(resp.Bytes(), &m); err != nil {
		return
	}
	var walk func(interface{})
	walk = func(v interface{}) {
		switch vv := v.(type) {
		case map[string]interface{}:
			if cve, ok := vv["cve"]; ok {
				if cveMap, ok := cve.(map[string]interface{}); ok {
					if id, ok := cveMap["id"].(string); ok {
						cves = append(cves, id)
					}
				}
			}
			for _, v2 := range vv {
				walk(v2)
			}
		case []interface{}:
			for _, v2 := range vv {
				walk(v2)
			}
		}
	}
	walk(m)
	return
}

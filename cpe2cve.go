package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dusbot/cpe2cve/core"
)

func main() {
	cpeInput := flag.String("cpe", "", "CPE(s) to query, comma separated (support cpe2.2 and cpe2.3)")
	output := flag.String("o", "", "output file (optional)")
	flag.Parse()

	if *cpeInput == "" {
		fmt.Println("Usage: cpe2cve -cpe <cpe1>[,<cpe2>,...] [-o output.txt]")
		return
	}

	var out *os.File
	var err error
	if *output != "" {
		out, err = os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create file: %v\n", err)
			return
		}
		defer out.Close()
	} else {
		out = os.Stdout
	}

	cpes := splitAndTrim(*cpeInput)
	for _, cpe := range cpes {
		cves := core.CPE2CVE(cpe)
		fmt.Fprintf(out, "CPE: %s\nCVEs: %v\n", cpe, cves)
	}
}

func splitAndTrim(s string) []string {
	var res []string
	for _, v := range splitComma(s) {
		v = trimSpace(v)
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func splitComma(s string) []string {
	return strings.Split(s, ",")
}

func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

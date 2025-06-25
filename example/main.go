package main

import (
	"fmt"

	"github.com/dusbot/cpe2cve/core"
)

func main() {
	cves := core.CPE2CVE("cpe:/a:apache:tomcat:7.0.65")
	fmt.Println(cves)
}

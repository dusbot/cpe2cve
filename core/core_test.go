package core

import (
	"fmt"
	"testing"
)

func TestCPE2CVE(t *testing.T) {
	cves0 := CPE2CVE("cpe:2.3:a:apache:http_server:2.4.54")
	fmt.Println(cves0)
	cves1 := CPE2CVE("cpe:/a:apache:tomcat:7.0.65")
	fmt.Println(cves1)
}

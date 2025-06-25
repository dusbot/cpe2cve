# cpe2cve

## 使用

### 命令行
#### 编译
```bash
git clone https://github.com/dusbot/cpe2cve.git
cd cpe2cve
go mod tidy&&go build -trimpath -ldflags "-s -w"
```
#### 运行
```shell
root@localhost: ./cpe2cve -h
Usage of cpe2cve:
  -cpe string
        CPE(s) to query, comma separated (support cpe2.2 and cpe2.3)
  -o string
        output file (optional)

# 查询示例
./cpe2cve -cpe cpe:/a:apache:tomcat:7.0.65,cpe:2.3:a:apache:http_server:2.4.54
# 查询输出
CPE: cpe:/a:apache:tomcat:7.0.65
CVEs: [CVE-2012-5568 CVE-2015-5345 CVE-2015-5346 CVE-2015-5351 CVE-2016-0706 CVE-2016-0714 CVE-2016-0763 CVE-2016-3092 CVE-2016-5388 CVE-2016-6816 CVE-2016-8735 CVE-2017-5647 CVE-2017-5648 CVE-2017-5664 CVE-2016-0762 CVE-2016-5018 CVE-2016-6794 CVE-2016-6797 CVE-2016-8745 CVE-2016-6796 CVE-2017-7674 CVE-2014-9634 CVE-2014-9635 CVE-2017-12615 CVE-2017-12616 CVE-2017-12617 CVE-2018-1305 CVE-2018-1304 CVE-2018-8014 CVE-2018-8034 CVE-2018-1336 CVE-2018-11784 CVE-2019-0232 CVE-2019-2684 CVE-2019-0221 CVE-2019-17563 CVE-2019-12418 CVE-2020-1935 CVE-2020-1938 CVE-2020-9484 CVE-2020-8022 CVE-2020-13935 CVE-2021-24122 CVE-2021-25329 CVE-2021-30640]
CPE: cpe:2.3:a:apache:http_server:2.4.54
CVEs: [CVE-2007-4723 CVE-2009-0796 CVE-2009-2299 CVE-2011-1176 CVE-2011-2688 CVE-2012-3526 CVE-2012-4001 CVE-2012-4360 CVE-2013-0941 CVE-2013-0942 CVE-2013-2765 CVE-2013-4365 CVE-2006-20001 CVE-2022-36760 CVE-2022-37436 CVE-2023-25690 CVE-2023-27522 CVE-2023-31122 CVE-2023-45802 CVE-2024-27316 CVE-2024-38474 CVE-2024-38475 CVE-2024-38476 CVE-2024-38477 CVE-2024-40898]
```

### 代码调用

#### 安装

> go get github.com/dusbot/cpe2cve

#### 使用

```go
package main

import (
	"fmt"

	"github.com/dusbot/cpe2cve/core"
)

func main() {
	cves := core.CPE2CVE("cpe:/a:apache:tomcat:7.0.65")
	fmt.Println(cves)
}
```

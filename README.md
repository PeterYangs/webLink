# 获取网站所有链接

### 安装
```shell
go get github.com/PeterYangs/webLink
```


### 使用
```go
package main

import (
	"context"
	"github.com/PeterYangs/webLink"
	"time"
)

func main() {

	c, _ := context.WithTimeout(context.Background(), 30*time.Second)

	w := webLink.NewWebLink(c, "url.txt")

	w.Link("https://www.925g.com/")
    
	//正则表达式
	w.Regular("")
	
	w.Run()

}
```
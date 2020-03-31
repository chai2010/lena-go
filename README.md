- *赞助 BTC: 1Cbd6oGAUUyBi7X7MaR4np4nTmQZXVgkCW*
- *赞助 ETH: 0x623A3C3a72186A6336C79b18Ac1eD36e1c71A8a6*
- *Go语言付费QQ群: 1055927514*

----

# lena-go

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chai2010/lena-go"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, lena.Data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

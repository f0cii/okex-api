# okex-api
OKEx Open API V3 SDK (Golang Version)

-----

### 1.Downloads or updates OKEX code's dependencies, in your command line:

```
go get -u github.com/frankrap/okex-api
```
### 2.Write the go file. warm tips: test go file, must suffix *_test.go, eg: okex_open_api_v3_test.go
```
package gotest

import (
	"fmt"
	"github.com/frankrap/okex-api"
	"testing"
)

func TestOKExServerTime(t *testing.T) {
	serverTime, err := okex.NewOKExClient().GetServerTime()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("OKEx's server time: ", serverTime)
}

func NewOKExClient() *okex.Client {
	var config okex.Config
	config.Endpoint = "https://www.okex.com/"
	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = okex.ENGLISH

	client := okex.NewClient(config)
	return client
}
```
### 3. run test go:
```
go test -v -run TestOKExServerTime okex_open_api_v3_test.go
```

### go-sanbod

A Golang SDK for sanbod API.

All the REST APIs listed in sanbod API document are implemented.

For best compatibility, please use Go >= 1.22.

Make sure you have read sanbod API document before continuing.

### Installation

```shell
go get github.com/parparvaz/sanbod-sdk-golang
```

### REST API

#### Setup

Init client for API services. Get ApiKey from your sanbod account.

```golang
var (
    username = "your username"
    password = "your password"
)
client := sanbod.NewClient(username, password)
```

A service instance stands for a REST API endpoint and is initialized by client.NewXXXService function.

Simply call API in chain style. Call Do() in the end to send HTTP request.

If you have any questions, please refer to the specific reference definitions or usage methods

##### Proxy Client

```golang
proxyUrl := "http://127.0.0.1:7890" // Please replace it with your exact proxy URL.
client := sanbod.NewProxyClient(username, password, proxyUrl)
```


#### Match National Code With Card Number

```golang
res, err := client.NewMatchNationalCodeWithCardNumberService().
	MobileNumber("Mobile Number").
	NationalCode("National Code").
	CardNumber("Card Number").
	Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)

```

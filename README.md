# Pionex

golang sdk for [pionex open API](https://pionex-doc.gitbook.io/apidocs)

```
├── api  
│   ├── api.go
│   ├── singer.go
│   └── utils.go
├── account
│   └── account.go
├── market
│   └── market.go
└── order
    └── order.go
```

## usage

```shell
go get -u github.com/luanruisong/pionex
```

create signer by your ApiKey

```go
signer := api.NewSigner(
    "{Your API Key}",
    "{Your API Secret",
)
```

init api by singer and http client

```go
api.WithSinger(signer)
api.WithHttpClient(http.DefaultClient)
```

market interface

```go
ret, err := market.GetSymbols(&market.SymbolsReq{
    Symbols: "",
    Type:    "",
})
...
```

account interface

```go
ret, err := account.BalancesInfo()
...
```

order interface

```go
ret, err := order.NewOrder(&order.NewOrderReq{
    Side: "BUY",
    Type: "MARKET",
    Symbol: "PEPE_USDT",
    Amount: "0.009",
})
```

## TODO 
 - quotation websocket
 - private websocket
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
    "{Your API Secret}",
)
```

market interface

```go
mkt := market.NewMarket(http.DefaultClient)
ret,err := mkt.GetSymbols(&SymbolsReq{
	//todo
})
...
```

account interface

```go
acc := account.NewAccount(sign,http.DefaultClient)
ret, err := acc.BalancesInfo()
...
```

order interface

```go
trans := order.NewTrans(sign,http.DefaultClient)
ret, err := trans.NewOrder(&order.NewOrderReq{
    Side: "BUY",
    Type: "MARKET",
    Symbol: "PEPE_USDT",
    Amount: "0.009",
})
...
```

## TODO 
 - quotation websocket
 - private websocket
package main

import (
	"github.com/jinzhu/gorm"
)

// database table
type Account struct {
	gorm.Model
	NickName      string
	ExchangeName  string
	ApiKey        string
	ApiSecretKey  string
	ApiPassphrase string
	Assets        []Asset
}

type Asset struct {
	gorm.Model
	Btc      float64
	Usdt     float64
	Btc_Usdt float64
	Usdt_Usd float64
	Usd_Cny  float64
	Coins    []CoinAsset
}

type CoinAsset struct {
	gorm.Model
	AssetID      uint
	CoinName     string
	Amount       float64
	FrozenAmount float64
}

// configure
type Config struct {
	Freq  int    `toml:"freq"`
	Proxy string `toml:proxy`
	User  User   `toml:"user"`
}

type User struct {
	UserName string
	Password string
}

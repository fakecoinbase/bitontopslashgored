package test

import (
	"log"
	"testing"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/exchange/okex"
	"github.com/bitontop/gored/pair"
	"github.com/bitontop/gored/test/conf"
	"github.com/bitontop/gored/utils"
	// "../../exchange/okex"
	// "../conf"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Okex(t *testing.T) {
	e := InitOkex()

	pair := pair.GetPairByKey("BTC|ETH")

	// Test_Coins(e)
	// Test_Pairs(e)
	Test_Pair(e, pair)
	// Test_Orderbook(e, pair)
	// Test_ConstraintFetch(e, pair)
	// Test_Constraint(e, pair)

	// new interface methods
	Test_DoWithdraw(e, pair.Target, "1", "0x37E0Fc27C6cDB5035B2a3d0682B4E7C05A4e6C46", "tag")
	Test_DoTransfer(e, pair.Target, "10", exchange.AssetWallet, exchange.SpotWallet)
	Test_CheckBalance(e, pair.Target, exchange.AssetWallet)
	Test_CheckAllBalance(e, exchange.AssetWallet)

	// okex.Socket(pair)
	// Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.00000001, 100)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")
}

func Test_OKEX_TradeHistory(t *testing.T) {
	e := InitOkex()
	p := pair.GetPairByKey("USDT|LTC")

	opTradeHistory := &exchange.PublicOperation{
		Type:      exchange.TradeHistory,
		EX:        e.GetName(),
		Pair:      p,
		DebugMode: true,
	}

	err := e.LoadPublicData(opTradeHistory)
	if err != nil {
		log.Printf("%v", err)
	}

	log.Printf("TradeHistory: %s::%s", opTradeHistory.EX, opTradeHistory.Pair.Name)

	for _, d := range opTradeHistory.TradeHistory {
		log.Printf(">> %+v ", d)
	}
}

func InitOkex() exchange.Exchange {
	coin.Init()
	pair.Init()

	utils.GetCommonDataFromJSON("https://raw.githubusercontent.com/bitontop/gored/master/data")

	config := &exchange.Config{}
	config.Source = exchange.JSON_FILE
	config.SourceURI = "https://raw.githubusercontent.com/bitontop/gored/master/data"

	conf.Exchange(exchange.OKEX, config)

	ex := okex.CreateOkex(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}

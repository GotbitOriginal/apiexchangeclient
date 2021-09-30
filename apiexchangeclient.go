package main

import (
	"gitlab.com/gotbitbot/apiclient"

	"gitlab.com/gotbitbot/alterdiceapi"
	"gitlab.com/gotbitbot/altillyapi"
	"gitlab.com/gotbitbot/altsbitapi"
	azbitapi "gitlab.com/gotbitbot/azbit"
	"gitlab.com/gotbitbot/biboxapi"
	"gitlab.com/gotbitbot/bideskapi"
	"gitlab.com/gotbitbot/bigoneapi"
	"gitlab.com/gotbitbot/bikiapi"
	"gitlab.com/gotbitbot/bilaxyapi"
	binanceapi "gitlab.com/gotbitbot/binanceapi"
	"gitlab.com/gotbitbot/binancedexapi"
	"gitlab.com/gotbitbot/bitcoinapi"
	"gitlab.com/gotbitbot/bitforexapi"
	"gitlab.com/gotbitbot/bithumbapi"
	"gitlab.com/gotbitbot/bitmartapi"
	"gitlab.com/gotbitbot/bitmaxapi"
	"gitlab.com/gotbitbot/bitrueapi"
	"gitlab.com/gotbitbot/bittrexapi"
	"gitlab.com/gotbitbot/bkexapi"
	"gitlab.com/gotbitbot/btcalphaapi"
	"gitlab.com/gotbitbot/bwapi"
	c3api "gitlab.com/gotbitbot/c3"
	"gitlab.com/gotbitbot/chainxapi"
	"gitlab.com/gotbitbot/citexapi"
	"gitlab.com/gotbitbot/coinallapi"
	"gitlab.com/gotbitbot/coindcxapi"
	"gitlab.com/gotbitbot/coinealapi"
	"gitlab.com/gotbitbot/coinfieldapi"
	"gitlab.com/gotbitbot/coinsbitapi"
	"gitlab.com/gotbitbot/coinsuperapi"
	"gitlab.com/gotbitbot/cointigerapi"
	cointradersapi "gitlab.com/gotbitbot/cointradeapi"
	"gitlab.com/gotbitbot/crex24api"
	"gitlab.com/gotbitbot/dcoinapi"
	"gitlab.com/gotbitbot/dexfinapi"
	"gitlab.com/gotbitbot/digifinexapi"
	"gitlab.com/gotbitbot/dobiexchangeapi"
	exmarketapi "gitlab.com/gotbitbot/exmarketsapi"
	"gitlab.com/gotbitbot/exratesapi"
	"gitlab.com/gotbitbot/folgoryapi"
	"gitlab.com/gotbitbot/gateioapi"
	"gitlab.com/gotbitbot/graviexapi"
	"gitlab.com/gotbitbot/hitbtcapi"
	"gitlab.com/gotbitbot/hooapi"
	"gitlab.com/gotbitbot/hotbitapi"
	"gitlab.com/gotbitbot/huobiapi"
	"gitlab.com/gotbitbot/idaxapi"
	idexapi "gitlab.com/gotbitbot/idex"
	"gitlab.com/gotbitbot/kucoinapi"
	"gitlab.com/gotbitbot/latokenapi"
	"gitlab.com/gotbitbot/lbankapi"
	"gitlab.com/gotbitbot/livecoinapi"
	"gitlab.com/gotbitbot/mxcapi"
	"gitlab.com/gotbitbot/oceanexapi"
	"gitlab.com/gotbitbot/p2pb2bapi"
	"gitlab.com/gotbitbot/poloniexapi"
	"gitlab.com/gotbitbot/probitapi"
	"gitlab.com/gotbitbot/rivermountapi"
	"gitlab.com/gotbitbot/satoapi"
	"gitlab.com/gotbitbot/southxchangeapi"
	"gitlab.com/gotbitbot/stexapi"
	"gitlab.com/gotbitbot/timexapi"
	"gitlab.com/gotbitbot/tradeogreapi"
	"gitlab.com/gotbitbot/ubixapi"
	unitTestapi "gitlab.com/gotbitbot/unit_testapi"
	"gitlab.com/gotbitbot/unitapi"
	"gitlab.com/gotbitbot/vccapi"
	virgoapi "gitlab.com/gotbitbot/virgoapi"
	"gitlab.com/gotbitbot/whitebitapi"
	"gitlab.com/gotbitbot/xtapi"
	"gitlab.com/gotbitbot/zbapi"
)

type ExchangesNames string

const EmptyExchange ExchangesNames = ""

const (
	hotbitExchange       ExchangesNames = "hotbit"
	binancedexExchange   ExchangesNames = "binancedex"
	binanceExchange      ExchangesNames = "binance"
	eterbaseExchange     ExchangesNames = "eterbase"
	latokenExchange      ExchangesNames = "latoken"
	bitforexExchange     ExchangesNames = "bitforex"
	p2pb2bExchange       ExchangesNames = "p2pb2b"
	probitExchange       ExchangesNames = "probit"
	idaxExchange         ExchangesNames = "idax"
	altsbitExchange      ExchangesNames = "altsbit"
	bitforexapiExchange  ExchangesNames = "bitforex"
	latokenapiExchange   ExchangesNames = "latoken"
	kucoinapiExchange    ExchangesNames = "kucoin"
	coinsbitExchange     ExchangesNames = "coinsbit"
	hitbtcExchange       ExchangesNames = "hitbtc"
	exmarketsExchange    ExchangesNames = "exmarkets"
	dobiexchangeExchange ExchangesNames = "dobiexchange"
	chainxExchange       ExchangesNames = "chainx"
	bitcoinExchange      ExchangesNames = "bitcoin"
	bithumbExchange      ExchangesNames = "bithumb"
	bitmartExchange      ExchangesNames = "bitmart"
	biboxExchange        ExchangesNames = "bibox"
	exratesExchange      ExchangesNames = "exrates"
	whitebitExchange     ExchangesNames = "whitebit"
	btcalphaExchange     ExchangesNames = "btcalpha"
	idexExchange         ExchangesNames = "idex"
	livecoinExchange     ExchangesNames = "livecoin"
	gateioExchange       ExchangesNames = "gateio"
	citexExchange        ExchangesNames = "citex"
	stexExchange         ExchangesNames = "stex"
	crex24Exchange       ExchangesNames = "crex24"
	digifinexExchange    ExchangesNames = "digifinex"
	cointradersExchange  ExchangesNames = "cointraders"
	bilaxyExchange       ExchangesNames = "bilaxy"
	bittrexExchange      ExchangesNames = "bittrex"
	huobiExchange        ExchangesNames = "huobi"
	poloniexExchange     ExchangesNames = "poloniex"
	coinleanExchange     ExchangesNames = "coineal"
	unitExchange         ExchangesNames = "unit"
	hooExchange          ExchangesNames = "hoo"
	graviexExchange      ExchangesNames = "graviex"
	alterdiceExchange    ExchangesNames = "alterdic"
	bitrueExchange       ExchangesNames = "bitrue"
	coindcxExchange      ExchangesNames = "coindcx"
	bitmaxExchange       ExchangesNames = "bitmax"
	altillyExchange      ExchangesNames = "altilly"
	folgoryExchange      ExchangesNames = "folgory"
	dcoinExchange        ExchangesNames = "dcoin"
	southxchangeExchange ExchangesNames = "southxchange"
	coinallExchange      ExchangesNames = "coinall"
	lbankExchange        ExchangesNames = "lbank"
	tradeogreExchange    ExchangesNames = "tradeogre"
	unitTestExchange     ExchangesNames = "unit_test"
	rivermountExchange   ExchangesNames = "rivermount"
	bwExchange           ExchangesNames = "bw"
	vccExchange          ExchangesNames = "vcc"
	cointigerExchange    ExchangesNames = "cointiger"
	coinfieldExchange    ExchangesNames = "coinfield"
	mxcExchange          ExchangesNames = "mxc"
	coinsuperExchange    ExchangesNames = "coinsuper"
	xtExchange           ExchangesNames = "xt"
	ubixExchange         ExchangesNames = "ubix"
	satoExchange         ExchangesNames = "sato"
	timexExchange        ExchangesNames = "timex"
	oceanexExchange      ExchangesNames = "oceanex"
	dexfinExchange       ExchangesNames = "dexfin"
	bideskExchange       ExchangesNames = "bidesk"
	c3Exchange           ExchangesNames = "c3"
	bikiExchange         ExchangesNames = "biki"
	bigoneExchange       ExchangesNames = "bigone"
	bkexExchange         ExchangesNames = "bkex"
	zbExchange           ExchangesNames = "zb"
	virgoExchange        ExchangesNames = "virgo"
	azbitExchange        ExchangesNames = "azbit"
)

func GetAPIClientByExchange(exchange ExchangesNames) apiclient.APIClient {
	switch exchange {
	case hotbitExchange:
		return &hotbitapi.HotbitApi{}
	case binancedexExchange:
		return &binancedexapi.BinanceDexApi{}
	case p2pb2bExchange:
		return &p2pb2bapi.P2PB2BApi{}
	case probitExchange:
		return &probitapi.ProbitApi{}
	case idaxExchange:
		return &idaxapi.IDAXApi{}
	case altsbitExchange:
		return &altsbitapi.AltsBitApi{}
	case bitforexapiExchange:
		return &bitforexapi.BitforexAPI{}
	case latokenapiExchange:
		return &latokenapi.LatokenApi{}
	case kucoinapiExchange:
		return &kucoinapi.KucoinApi{}
	case coinsbitExchange:
		return &coinsbitapi.CoinsbitApi{}
	case hitbtcExchange:
		return &hitbtcapi.HitBTCApi{}
	case exmarketsExchange:
		return &exmarketapi.ExmarketApi{}
	case dobiexchangeExchange:
		return &dobiexchangeapi.DobiExchangeApi{}
	case chainxExchange:
		return &chainxapi.ChainXApi{}
	case bitcoinExchange:
		return &bitcoinapi.BitcoinApi{}
	case bithumbExchange:
		return &bithumbapi.BitHumbApi{}
	case bitmartExchange:
		return &bitmartapi.BitmartApi{}
	case biboxExchange:
		return &biboxapi.BiBoxApi{}
	case exratesExchange:
		return &exratesapi.ExRatesApi{}
	case whitebitExchange:
		return &whitebitapi.WhiteBitApi{}
	case btcalphaExchange:
		return &btcalphaapi.BtcAlphaApi{}
	case idexExchange:
		return &idexapi.IdexApi{}
	case livecoinExchange:
		return &livecoinapi.LiveCoinApi{}
	case gateioExchange:
		return &gateioapi.GateIoApi{}
	case citexExchange:
		return &citexapi.CitexApi{}
	case stexExchange:
		return &stexapi.StexApi{}
	case crex24Exchange:
		return &crex24api.Crex24Api{}
	case digifinexExchange:
		return &digifinexapi.DigifinexApi{}
	case cointradersExchange:
		return &cointradersapi.CoinTradersApi{}
	case bilaxyExchange:
		return &bilaxyapi.BilaxyApi{}
	case bittrexExchange:
		return &bittrexapi.BittrexApi{}
	case huobiExchange:
		return &huobiapi.HuobiApi{}
	case poloniexExchange:
		return &poloniexapi.PoloniexApi{}
	case coinleanExchange:
		return &coinealapi.CoinealApi{}
	case unitExchange:
		return &unitapi.UnitApi{}
	case hooExchange:
		return &hooapi.HooApi{}
	case graviexExchange:
		return &graviexapi.GraviexApi{}
	case alterdiceExchange:
		return &alterdiceapi.AlterdiceApi{}
	case bitrueExchange:
		return &bitrueapi.BitrueApi{}
	case coindcxExchange:
		return &coindcxapi.CoinDCXApi{}
	case bitmaxExchange:
		return &bitmaxapi.BitMaxApi{}
	case altillyExchange:
		return &altillyapi.AltillyApi{}
	case folgoryExchange:
		return &folgoryapi.FolgoryApi{}
	case dcoinExchange:
		return &dcoinapi.DCoinApi{}
	case southxchangeExchange:
		return &southxchangeapi.SouthXchangeApi{}
	case coinallExchange:
		return &coinallapi.CoinAllApi{}
	case lbankExchange:
		return &lbankapi.LbankApi{}
	case tradeogreExchange:
		return &tradeogreapi.TradeOgreApi{}
	case unitTestExchange:
		return &unitTestapi.UnitApi{}
	case rivermountExchange:
		return &rivermountapi.RiverMountApi{}
	case bwExchange:
		return &bwapi.BWApi{}
	case vccExchange:
		return &vccapi.VCCApi{}
	case cointigerExchange:
		return &cointigerapi.CoinTigerApi{}
	case binanceExchange:
		return &binanceapi.BinanceAPI{}
	case coinfieldExchange:
		return &coinfieldapi.CoinFieldApi{}
	case mxcExchange:
		return &mxcapi.MXCApi{}
	case coinsuperExchange:
		return &coinsuperapi.CoinsuperApi{}
	case xtExchange:
		return &xtapi.XTApi{}
	case ubixExchange:
		return &ubixapi.UbixApi{}
	case satoExchange:
		return &satoapi.SatoApi{}
	case timexExchange:
		return &timexapi.TimeXApi{}
	case oceanexExchange:
		return &oceanexapi.OceanExApi{}
	case dexfinExchange:
		return &dexfinapi.DexfinApi{}
	case bideskExchange:
		return &bideskapi.BideskApi{}
	case c3Exchange:
		return &c3api.C3Api{}
	case bikiExchange:
		return &bikiapi.BikiApi{}
	case bigoneExchange:
		return &bigoneapi.BigOneApi{}
	case bkexExchange:
		return &bkexapi.BkexApi{}
	case zbExchange:
		return &zbapi.ZBApi{}
	case virgoExchange:
		return &virgoapi.VirgoApi{}
	case azbitExchange:
		return &azbitapi.AzbitApi{}
	default:
		return nil
	}
}

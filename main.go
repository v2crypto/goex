package goex

import (
	"github.com/v2crypto/goex/v2/binance"
	"github.com/v2crypto/goex/v2/httpcli"
	"github.com/v2crypto/goex/v2/huobi"
	"github.com/v2crypto/goex/v2/logger"
	"github.com/v2crypto/goex/v2/okx"
	"reflect"
)

var (
	DefaultHttpCli = httpcli.Cli
)

var (
	OKx     = okx.New()
	Binance = binance.New()
	HuoBi   = huobi.New()
)

func SetDefaultHttpCli(cli httpcli.IHttpClient) {
	logger.Infof("use new http client implement: %s", reflect.TypeOf(cli).Elem().String())
	httpcli.Cli = cli
}

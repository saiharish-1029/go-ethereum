package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/internal/flags"
	"github.com/ethereum/go-ethereum/internal/web3ext"
	"github.com/urfave/cli/v2"
)

var Modules = web3ext.Modules

func GlobalBig(ctx *cli.Context, name string) *big.Int {
	return flags.GlobalBig(ctx, name)
}

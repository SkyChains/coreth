// (c) 2019-2020, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/skychains/chain/utils/logging"
	"github.com/skychains/chain/utils/ulimit"
	"github.com/skychains/chain/vms/rpcchainvm"

	"github.com/skychains/coreth/plugin/evm"
)

func main() {
	version, err := PrintVersion()
	if err != nil {
		fmt.Printf("couldn't get config: %s", err)
		os.Exit(1)
	}
	if version {
		fmt.Println(evm.Version)
		os.Exit(0)
	}
	if err := ulimit.Set(ulimit.DefaultFDLimit, logging.NoLog{}); err != nil {
		fmt.Printf("failed to set fd limit correctly due to: %s", err)
		os.Exit(1)
	}

	rpcchainvm.Serve(context.Background(), &evm.VM{IsPlugin: true})
}

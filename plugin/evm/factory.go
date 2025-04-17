// (c) 2019-2020, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/skychains/chain/ids"
	"github.com/skychains/chain/utils/logging"
	"github.com/skychains/chain/vms"
)

var (
	// ID this VM should be referenced by
	ID = ids.ID{'e', 'v', 'm'}

	_ vms.Factory = &Factory{}
)

type Factory struct{}

func (*Factory) New(logging.Logger) (interface{}, error) {
	return &VM{}, nil
}

// (c) 2019-2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"context"
	"fmt"

	"github.com/skychains/chain/codec"
	"github.com/skychains/chain/ids"
)

// Request represents a Network request type
type Request interface {
	// Requests should implement String() for logging.
	fmt.Stringer

	// Handle allows `Request` to call respective methods on handler to handle
	// this particular request type
	Handle(ctx context.Context, nodeID ids.NodeID, requestID uint32, handler RequestHandler) ([]byte, error)
}

// BytesToRequest unmarshals the given requestBytes into Request object
func BytesToRequest(codec codec.Manager, requestBytes []byte) (Request, error) {
	var request Request
	if _, err := codec.Unmarshal(requestBytes, &request); err != nil {
		return nil, err
	}
	return request, nil
}

// RequestToBytes marshals the given request object into bytes
func RequestToBytes(codec codec.Manager, request Request) ([]byte, error) {
	return codec.Marshal(Version, &request)
}

// CrossChainRequest represents the interface a cross chain request should implement
type CrossChainRequest interface {
	// CrossChainRequest should implement String() for logging.
	fmt.Stringer

	// Handle allows [CrossChainRequest] to call respective methods on handler to handle
	// this particular request type
	Handle(ctx context.Context, requestingChainID ids.ID, requestID uint32, handler CrossChainRequestHandler) ([]byte, error)
}

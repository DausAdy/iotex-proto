// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package golang

import (
	"errors"

	"google.golang.org/protobuf/proto"

	"github.com/iotexproject/iotex-proto/golang/iotexrpc"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
	"github.com/iotexproject/iotex-proto/golang/testingpb"
)

// GetTypeFromRPCMsg retrieves the proto message type
func GetTypeFromRPCMsg(msg proto.Message) (iotexrpc.MessageType, error) {
	switch msg.(type) {
	case *iotextypes.Block:
		return iotexrpc.MessageType_BLOCK, nil
	case *iotexrpc.BlockSync:
		return iotexrpc.MessageType_BLOCK_REQUEST, nil
	case *iotextypes.Action:
		return iotexrpc.MessageType_ACTION, nil
	case *iotextypes.Actions:
		return iotexrpc.MessageType_ACTIONS, nil
	case *iotextypes.ActionHash:
		return iotexrpc.MessageType_ACTION_HASH, nil
	case *iotexrpc.ActionSync:
		return iotexrpc.MessageType_ACTION_REQUEST, nil
	case *iotextypes.ConsensusMessage:
		return iotexrpc.MessageType_CONSENSUS, nil
	case *iotextypes.NodeInfoRequest:
		return iotexrpc.MessageType_NODE_INFO_REQUEST, nil
	case *iotextypes.NodeInfo:
		return iotexrpc.MessageType_NODE_INFO, nil
	case *testingpb.TestPayload:
		return iotexrpc.MessageType_TEST, nil
	default:
		return iotexrpc.MessageType_UNKNOWN, errors.New("unknown RPC message type")
	}
}

// TypifyRPCMsg unmarshal a proto message based on the given MessageType
func TypifyRPCMsg(t iotexrpc.MessageType, msg []byte) (proto.Message, error) {
	var m proto.Message
	switch t {
	case iotexrpc.MessageType_BLOCK:
		m = &iotextypes.Block{}
	case iotexrpc.MessageType_CONSENSUS:
		m = &iotextypes.ConsensusMessage{}
	case iotexrpc.MessageType_BLOCK_REQUEST:
		m = &iotexrpc.BlockSync{}
	case iotexrpc.MessageType_ACTION:
		m = &iotextypes.Action{}
	case iotexrpc.MessageType_ACTIONS:
		m = &iotextypes.Actions{}
	case iotexrpc.MessageType_ACTION_HASH:
		m = &iotextypes.ActionHash{}
	case iotexrpc.MessageType_ACTION_REQUEST:
		m = &iotexrpc.ActionSync{}
	case iotexrpc.MessageType_NODE_INFO_REQUEST:
		m = &iotextypes.NodeInfoRequest{}
	case iotexrpc.MessageType_NODE_INFO:
		m = &iotextypes.NodeInfo{}
	case iotexrpc.MessageType_TEST:
		m = &testingpb.TestPayload{}
	default:
		return nil, errors.New("unknown RPC message type")
	}

	err := proto.Unmarshal(msg, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

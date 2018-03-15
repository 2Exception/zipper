// Copyright (C) 2017, Zipper Team.  All rights reserved.
//
// This file is part of zipper
//
// The zipper is free software: you can use, copy, modify,
// and distribute this software for any purpose with or
// without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// The zipper is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// ISC License for more details.
//
// You should have received a copy of the ISC License
// along with this program.  If not, see <https://opensource.org/licenses/isc>.


package protoManager

import (
	"fmt"
	"sync"

	"github.com/zipper-project/zipper/common/mpool"
	"github.com/zipper-project/zipper/peer"
	"github.com/zipper-project/zipper/peer/proto"
	mproto "github.com/zipper-project/zipper/proto"
	"github.com/zipper-project/zipper/types"
)

type ProtoManager struct {
	sync.Mutex
	wm map[mproto.ProtoID]*mpool.VirtualMachine
}

func NewProtoManager() *ProtoManager {
	return &ProtoManager{
		wm: make(map[mproto.ProtoID]*mpool.VirtualMachine),
	}
}

func (pm *ProtoManager) RegisterWorker(protocalID mproto.ProtoID, workers []mpool.VmWorker) error {
	pm.Lock()
	defer pm.Unlock()

	if _, ok := pm.wm[protocalID]; ok {
		return fmt.Errorf("wm: %s have beed registered", protocalID)
	}

	pm.wm[protocalID] = mpool.CreateCustomVM(workers)
	return nil
}

func (pm *ProtoManager) Handle(sendPeer *peer.Peer, msg *proto.Message) error {
	pm.Lock()
	defer pm.Unlock()

	err := pm.wm[mproto.ProtoID(msg.Header.ProtoID)].SendWorkCleanAsync(types.NewWorkerData(sendPeer, msg))
	return err
}

func (pm *ProtoManager) InitAndRegisterWorker() {
	//pm.RegisterWorker(params.BlockSyncIdx, blocksync.NewSyncWorker())

}

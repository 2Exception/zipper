// Copyright (C) 2017, Beijing Bochen Technology Co.,Ltd.  All rights reserved.
//
// This file is part of L0
//
// The L0 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The L0 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/zipper-project/zipper/peer"
	"github.com/zipper-project/zipper/proto"
)

var (
	srv *peer.Server
)

func init() {
	option := peer.NewDefaultOption()
	option.ListenAddress = ":20066"
	option.BootstrapNodes = []string{"encode://303030315f616263@127.0.0.1:20166&1"}
	option.PeerID = []byte("0005_abc")
	srv = peer.NewServer(option, nil)
}

func Relay(tx *proto.Transaction) {
	//TODO
	//msg := p2p.NewMessage(nil, nil)
	//srv.Broadcast(msg, peer.VP)
}

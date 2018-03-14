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

package rpc

import (
	"bytes"
	"encoding/base64"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"

	"github.com/zipper-project/zipper/common/log"
	"github.com/zipper-project/zipper/blockchain"
	"github.com/zipper-project/zipper/config"
)


type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func NewHttConn(in io.Reader, out io.Writer) *HttpConn {
	return &HttpConn{
		in:  in,
		out: out,
	}
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }

func NewServer(bc *blockchain.Blockchain) *rpc.Server {

	server := rpc.NewServer()

	server.Register(NewRPCTransaction(bc))
	server.Register(NewRPCLedger(bc))
	return server
}

// StartServer with Test instance as a service
func StartServer(server *rpc.Server, option *config.RPCOption) {
	if option.Enabled == false {
		return
	}

	listener, err := net.Listen("tcp", ":"+option.Port)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer listener.Close()

	http.Serve(listener, http.HandlerFunc(BasicAuth(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := server.ServeRequest(serverCodec)
			if err != nil {
				log.Printf("Error while serving JSON request: %v", err)
				http.Error(w, "Error while serving JSON request, details have been logged.", 500)
				return
			}
		}
	}, option.User, option.PassWord)))
}

// ViewFunc defines view method
type ViewFunc func(http.ResponseWriter, *http.Request)

// BasicAuth handles basic authtication
func BasicAuth(f ViewFunc, user, passwd string) ViewFunc {
	if user == "" || passwd == "" {
		return f
	}
	return func(w http.ResponseWriter, r *http.Request) {
		basicAuthPrefix := "Basic "
		// get request header
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, basicAuthPrefix) {
			payload, err := base64.StdEncoding.DecodeString(
				auth[len(basicAuthPrefix):],
			)
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 && bytes.Equal(pair[0], []byte(user)) &&
					bytes.Equal(pair[1], []byte(passwd)) {
					f(w, r)
					return
				}
			}
		}
		http.Error(w, "Error user or password.", http.StatusUnauthorized)
	}
}


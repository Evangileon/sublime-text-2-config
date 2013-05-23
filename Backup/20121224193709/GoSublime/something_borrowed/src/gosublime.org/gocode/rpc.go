// WARNING! Autogenerated by goremote, don't touch.

package gocode

import (
	"net/rpc"
)

type RPC struct {
}

// wrapper for: server_auto_complete

type Args_auto_complete struct {
	Arg0 []byte
	Arg1 string
	Arg2 int
}
type Reply_auto_complete struct {
	Arg0 []candidate
	Arg1 int
}

func (r *RPC) RPC_auto_complete(args *Args_auto_complete, reply *Reply_auto_complete) error {
	reply.Arg0, reply.Arg1 = server_auto_complete(args.Arg0, args.Arg1, args.Arg2)
	return nil
}
func client_auto_complete(cli *rpc.Client, Arg0 []byte, Arg1 string, Arg2 int) (c []candidate, d int) {
	var args Args_auto_complete
	var reply Reply_auto_complete
	args.Arg0 = Arg0
	args.Arg1 = Arg1
	args.Arg2 = Arg2
	err := cli.Call("RPC.RPC_auto_complete", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0, reply.Arg1
}

// wrapper for: server_cursor_type_pkg

type Args_cursor_type_pkg struct {
	Arg0 []byte
	Arg1 string
	Arg2 int
}
type Reply_cursor_type_pkg struct {
	Arg0, Arg1 string
}

func (r *RPC) RPC_cursor_type_pkg(args *Args_cursor_type_pkg, reply *Reply_cursor_type_pkg) error {
	reply.Arg0, reply.Arg1 = server_cursor_type_pkg(args.Arg0, args.Arg1, args.Arg2)
	return nil
}
func client_cursor_type_pkg(cli *rpc.Client, Arg0 []byte, Arg1 string, Arg2 int) (typ, pkg string) {
	var args Args_cursor_type_pkg
	var reply Reply_cursor_type_pkg
	args.Arg0 = Arg0
	args.Arg1 = Arg1
	args.Arg2 = Arg2
	err := cli.Call("RPC.RPC_cursor_type_pkg", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0, reply.Arg1
}

// wrapper for: server_close

type Args_close struct {
	Arg0 int
}
type Reply_close struct {
	Arg0 int
}

func (r *RPC) RPC_close(args *Args_close, reply *Reply_close) error {
	reply.Arg0 = server_close(args.Arg0)
	return nil
}
func client_close(cli *rpc.Client, Arg0 int) int {
	var args Args_close
	var reply Reply_close
	args.Arg0 = Arg0
	err := cli.Call("RPC.RPC_close", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0
}

// wrapper for: server_status

type Args_status struct {
	Arg0 int
}
type Reply_status struct {
	Arg0 string
}

func (r *RPC) RPC_status(args *Args_status, reply *Reply_status) error {
	reply.Arg0 = server_status(args.Arg0)
	return nil
}
func client_status(cli *rpc.Client, Arg0 int) string {
	var args Args_status
	var reply Reply_status
	args.Arg0 = Arg0
	err := cli.Call("RPC.RPC_status", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0
}

// wrapper for: server_drop_cache

type Args_drop_cache struct {
	Arg0 int
}
type Reply_drop_cache struct {
	Arg0 int
}

func (r *RPC) RPC_drop_cache(args *Args_drop_cache, reply *Reply_drop_cache) error {
	reply.Arg0 = server_drop_cache(args.Arg0)
	return nil
}
func client_drop_cache(cli *rpc.Client, Arg0 int) int {
	var args Args_drop_cache
	var reply Reply_drop_cache
	args.Arg0 = Arg0
	err := cli.Call("RPC.RPC_drop_cache", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0
}

// wrapper for: server_set

type Args_set struct {
	Arg0, Arg1 string
}
type Reply_set struct {
	Arg0 string
}

func (r *RPC) RPC_set(args *Args_set, reply *Reply_set) error {
	reply.Arg0 = server_set(args.Arg0, args.Arg1)
	return nil
}
func client_set(cli *rpc.Client, Arg0, Arg1 string) string {
	var args Args_set
	var reply Reply_set
	args.Arg0 = Arg0
	args.Arg1 = Arg1
	err := cli.Call("RPC.RPC_set", &args, &reply)
	if err != nil {
		panic(err)
	}
	return reply.Arg0
}

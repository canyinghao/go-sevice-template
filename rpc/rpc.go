package rpc

import "net/rpc"

type Arith int

type Args struct {
	A, B int
}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// 注册rpc
func Register() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

}

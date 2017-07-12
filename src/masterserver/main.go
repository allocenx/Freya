package main

import (
    "share/logger"
    "masterserver/def"
    "share/lib/rpc2"
    "fmt"
    "net"
    "share/rpc/models/server"
)

var log = logger.Instance()

var g_ServerConfig   = def.ServerConfig
var g_ServerSettings = def.ServerSettings

type Args struct{ A, B int }
type Reply int

func main() {
    log.Info("MasterServer init")

    // read config
    //g_ServerConfig.Read()

    srv := rpc2.NewServer()
    srv.Handle("add", func(client *rpc2.Client, args *Args, reply *Reply) error {
        // Reversed call (server to client)
        var rep Reply
        client.Call("mult", Args{2, 3}, &rep)
        fmt.Println("mult result:", rep)

        *reply = Reply(args.A + args.B)
        return nil
    })
    srv.Handle("ServerRegister",
        func(
            client *rpc2.Client,
            request server.RegisterRequest,
            reply *server.RegisterResponse) error {

            log.Info(request)
        *reply = server.RegisterResponse{true}
        return nil
    })

    lis, _ := net.Listen("tcp", "127.0.0.1:9001")
    srv.Accept(lis)
}
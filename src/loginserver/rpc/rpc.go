package rpc

import (
    "share/logger"
    "share/rpc"
    "loginserver/def"
)

var log = logger.Instance()

var g_ServerConfig   = def.ServerConfig
var g_ServerSettings = def.ServerSettings
var g_NetworkManager = def.NetworkManager
var g_PacketHandler  = def.PacketHandler
var g_RPCHandler     = def.RPCHandler

func RegisterCalls() {
    g_RPCHandler.Register(rpc.UserVerify, UserVerify)
}
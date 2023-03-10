package handler

import (
	"log"

	"github.com/wztzt/gameserver/message"
)

func handle_LoginRequest(req *message.LoginRequest) {
	log.Printf("Hello User %v, Password %v\n", req.GetAccount(), req.GetPassword())
}

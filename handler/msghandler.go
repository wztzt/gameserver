package handler

import (
	"fmt"

	"github.com/wztzt/gameserver/message"
)

func handle_LoginRequest(req *message.LoginRequest) {
	fmt.Printf("Hello User %v, Password %v\n", req.GetAccount(), req.GetPassword())
}

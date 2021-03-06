package verify

import (
	"fmt"
	"github.com/davyxu/cellmesh/demo/proto"
	"github.com/davyxu/cellmesh/demo/svc/agent/api"
	"github.com/davyxu/cellmesh/service"
	"github.com/davyxu/cellnet"
)

func init() {

	proto.Handle_Game_VerifyREQ = agentapi.HandleBackendMessage(func(ev cellnet.Event, cid proto.ClientID) {

		msg := ev.Message().(*proto.VerifyREQ)

		fmt.Printf("verfiy: %+v \n", msg.GameToken)

		service.Reply(ev, &proto.VerifyACK{})
	})
}

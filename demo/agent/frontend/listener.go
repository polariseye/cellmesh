package frontend

import (
	"github.com/davyxu/cellmesh/demo/agent/model"
	"github.com/davyxu/cellmesh/discovery"
	"github.com/davyxu/cellmesh/svcfx/model"
	"github.com/davyxu/cellmesh/util"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"time"
)

func Start(addr string) {

	clientListener := peer.NewGenericPeer("tcp.Acceptor", model.FrontendName, addr, nil)

	proc.BindProcessorHandler(clientListener, "tcp.frontend", nil)

	socketOpt := clientListener.(cellnet.TCPSocketOption)

	// 无延迟设置缓冲
	socketOpt.SetSocketBuffer(2048, 2048, true)

	// 40秒无读，20秒无写断开
	socketOpt.SetSocketDeadline(time.Second*40, time.Second*20)

	clientListener.Start()
	model.FrontendSessionManager = clientListener.(peer.SessionManager)

	// 保存端口
	listenPort := clientListener.(cellnet.TCPAcceptor).Port()

	host := util.GetLocalIP()

	sd := &discovery.ServiceDesc{
		Host: host,
		Port: listenPort,
		ID:   fxmodel.GetSvcID(model.FrontendName),
		Name: model.FrontendName,
		Tags: []string{fxmodel.Node},
	}

	model.AgentSvcID = sd.ID

	// 服务发现注册服务
	discovery.Default.Register(sd)
}

func Stop() {

	if model.FrontendSessionManager != nil {
		model.FrontendSessionManager.(cellnet.Peer).Stop()

		discovery.Default.Deregister(model.AgentSvcID)
	}

}

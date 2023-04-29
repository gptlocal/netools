package core_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/gptlocal/netools/app/dispatcher"
	"github.com/gptlocal/netools/app/proxyman"
	"github.com/gptlocal/netools/common"
	"github.com/gptlocal/netools/common/net"
	"github.com/gptlocal/netools/common/protocol"
	"github.com/gptlocal/netools/common/serial"
	"github.com/gptlocal/netools/common/uuid"
	. "github.com/gptlocal/netools/core"
	"github.com/gptlocal/netools/features/dns"
	"github.com/gptlocal/netools/features/dns/localdns"
	_ "github.com/gptlocal/netools/main/distro/all"
	"github.com/gptlocal/netools/proxy/dokodemo"
	"github.com/gptlocal/netools/proxy/vmess"
	"github.com/gptlocal/netools/proxy/vmess/outbound"
	"github.com/gptlocal/netools/testing/servers/tcp"
)

func TestNetoolDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestNetoolClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}

package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/gptlocal/netools/app/dispatcher"
	_ "github.com/gptlocal/netools/app/proxyman/inbound"
	//_ "github.com/gptlocal/netools/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	//_ "github.com/gptlocal/netools/app/commander"
	//_ "github.com/gptlocal/netools/app/log/command"
	//_ "github.com/gptlocal/netools/app/proxyman/command"
	//_ "github.com/gptlocal/netools/app/stats/command"

	// Developer preview services
	//_ "github.com/gptlocal/netools/app/observatory/command"

	// Other optional features.
	//_ "github.com/gptlocal/netools/app/dns"
	//_ "github.com/gptlocal/netools/app/dns/fakedns"
	//_ "github.com/gptlocal/netools/app/log"
	//_ "github.com/gptlocal/netools/app/metrics"
	//_ "github.com/gptlocal/netools/app/policy"
	//_ "github.com/gptlocal/netools/app/reverse"
	//_ "github.com/gptlocal/netools/app/router"
	//_ "github.com/gptlocal/netools/app/stats"

	// Fix dependency cycle caused by core import in internet package
	//_ "github.com/gptlocal/netools/transport/internet/tagged/taggedimpl"

	// Developer preview features
	//_ "github.com/gptlocal/netools/app/observatory"

	// Inbound and outbound proxies.
	//_ "github.com/gptlocal/netools/proxy/blackhole"
	//_ "github.com/gptlocal/netools/proxy/dns"
	//_ "github.com/gptlocal/netools/proxy/dokodemo"
	//_ "github.com/gptlocal/netools/proxy/freedom"
	//_ "github.com/gptlocal/netools/proxy/http"
	//_ "github.com/gptlocal/netools/proxy/loopback"
	//_ "github.com/gptlocal/netools/proxy/mtproto"
	//_ "github.com/gptlocal/netools/proxy/shadowsocks"
	//_ "github.com/gptlocal/netools/proxy/socks"
	//_ "github.com/gptlocal/netools/proxy/trojan"
	//_ "github.com/gptlocal/netools/proxy/vless/inbound"
	//_ "github.com/gptlocal/netools/proxy/vless/outbound"
	//_ "github.com/gptlocal/netools/proxy/vmess/inbound"
	//_ "github.com/gptlocal/netools/proxy/vmess/outbound"
	//_ "github.com/gptlocal/netools/proxy/wireguard"

	// Transports
	//_ "github.com/gptlocal/netools/transport/internet/domainsocket"
	//_ "github.com/gptlocal/netools/transport/internet/grpc"
	//_ "github.com/gptlocal/netools/transport/internet/http"
	//_ "github.com/gptlocal/netools/transport/internet/kcp"
	//_ "github.com/gptlocal/netools/transport/internet/quic"
	//_ "github.com/gptlocal/netools/transport/internet/reality"
	//_ "github.com/gptlocal/netools/transport/internet/tcp"
	//_ "github.com/gptlocal/netools/transport/internet/tls"
	//_ "github.com/gptlocal/netools/transport/internet/udp"
	//_ "github.com/gptlocal/netools/transport/internet/websocket"

	// Transport headers
	//_ "github.com/gptlocal/netools/transport/internet/headers/http"
	//_ "github.com/gptlocal/netools/transport/internet/headers/noop"
	//_ "github.com/gptlocal/netools/transport/internet/headers/srtp"
	//_ "github.com/gptlocal/netools/transport/internet/headers/tls"
	//_ "github.com/gptlocal/netools/transport/internet/headers/utp"
	//_ "github.com/gptlocal/netools/transport/internet/headers/wechat"
	//_ "github.com/gptlocal/netools/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	//_ "github.com/gptlocal/netools/main/json"
	//_ "github.com/gptlocal/netools/main/toml"
	//_ "github.com/gptlocal/netools/main/yaml"

	// Load config from file or http(s)
	//_ "github.com/gptlocal/netools/main/confloader/external"

	// Commands
	_ "github.com/gptlocal/netools/main/commands/all"
)

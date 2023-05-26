package service

import (
	"fmt"
	zapPlugin "github.com/easysoft/zentaoatf/internal/pkg/plugin/zap/plugin"
	zapService "github.com/easysoft/zentaoatf/internal/pkg/plugin/zap/service"
	"github.com/easysoft/zentaoatf/internal/pkg/plugin/zap/shared"
	shellUtils "github.com/easysoft/zentaoatf/pkg/lib/shell"
	"github.com/hashicorp/go-plugin"
)

const (
	ZapPath = "/Users/aaron/rd/project/zentao/go/ztf/internal/pkg/plugin/zap-plugin"
)

type PluginService struct {
	zapClient    *plugin.Client
	zapRpcClient plugin.ClientProtocol
}

func (s *PluginService) Start() (err error) {
	s.zapClient = plugin.NewClient(&plugin.ClientConfig{
		Plugins: map[string]plugin.Plugin{
			zapShared.PluginNameZap: &zapPlugin.ZapPlugin{},
		},
		Cmd:              shellUtils.GetCmd(ZapPath),
		HandshakeConfig:  zapShared.Handshake,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})

	s.zapRpcClient, err = s.zapClient.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	// Request the plugin
	raw, err := s.zapRpcClient.Dispense(zapShared.PluginNameZap)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	zapService := raw.(zapService.ZapInterface)

	err = zapService.Put("key", []byte("Set Msg"))
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	result, err := zapService.Get("key")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println(string(result))

	return
}

func (s *PluginService) Stop() (err error) {
	s.zapClient.Kill()
	return
}

func (s *PluginService) Install() (err error) {

	return
}

func (s *PluginService) Uninstall() (err error) {

	return
}

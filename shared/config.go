package shared

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

// config options for the chain connection
type HypercocoaConfig struct {
	MspID         string
	CryptoPath    string
	CertPath      string
	KeyPath       string
	TlsCertPath   string
	PeerEndpoint  string
	GatewayPeer   string
	ChannelName   string
	ChaincodeName string
	// Config options for gateway mode
	BaseURL     string
	GatewayHost string
	GatewayPort string
}

var Hyperconfig *HypercocoaConfig

// InitConfig reades and fixes the config
func InitConfig() {
	viper.SetConfigName("hypercocoa")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("~/.hypercocoa")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Can not read config: %w", err))
	}
	err = viper.Unmarshal(&Hyperconfig)
	if err != nil {
		panic(fmt.Errorf("Can not compile config: %w", err))
	}
	Hyperconfig.CertPath = filepath.Join(Hyperconfig.CryptoPath, Hyperconfig.CertPath)
	Hyperconfig.KeyPath = filepath.Join(Hyperconfig.CryptoPath, Hyperconfig.KeyPath)
	Hyperconfig.TlsCertPath = filepath.Join(Hyperconfig.CryptoPath, Hyperconfig.TlsCertPath)
}

// PrintConfig prints the full config to stdout
func PrintConfig() {
	j, err := json.MarshalIndent(Hyperconfig, "", "\t")
	if err != nil {
		fmt.Println("Error printing config: %w", err)
	}
	fmt.Println(string(j))
}

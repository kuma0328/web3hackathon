package config

import "github.com/kuma0328/web3hackathon/utils/helper"

// Portはサーバーのport番号を返します
func Port() string {
	return helper.GetEnvOrDefault("PORT", "8080")
}

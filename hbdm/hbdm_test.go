package hbdm

import (
	"github.com/spf13/viper"
	"log"
)

func newTestClient() *Client {
	viper.SetConfigName("test_config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	accessKey := viper.GetString("access_key")
	secretKey := viper.GetString("secret_key")

	baseURL := "https://api.btcgateway.pro"
	apiParams := &ApiParameter{
		Debug:              true,
		AccessKey:          accessKey,
		SecretKey:          secretKey,
		EnablePrivateSign:  false,
		Url:                baseURL,
		PrivateKeyPrime256: "",
	}
	c := NewClient(apiParams)
	return c
}
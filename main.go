package main

import (
	"log"

	"github.com/haron1996/mpesa-go/funcs"
	"github.com/haron1996/mpesa-go/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Printf("could not load config: %v", err)
		return
	}

	accessToken, err := funcs.RequestAccessToken(config)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := funcs.RequestPayment(accessToken.AccessToken, config.PassKey, config.Callback)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(response)

}

package main

import (
	"log"

	"github.com/bronze1man/radius"
)

type radiusService struct{}

var radiusServer *radius.Server

func (p radiusService) RadiusHandle(request *radius.Packet) *radius.Packet {
	log.Printf("[auth] %s\n", request.String())
	npac := request.Reply()
	switch request.Code {
	case radius.AccessRequest:
		if checkCredentials(request.GetUsername(), request.GetPassword()) {
			npac.Code = radius.AccessAccept
			return npac
		}
		npac.Code = radius.AccessReject
		npac.AVPs = append(npac.AVPs, radius.AVP{Type: radius.ReplyMessage, Value: []byte("Go away!")})
		return npac
	case radius.AccountingRequest:
		npac.Code = radius.AccountingResponse
		return npac
	default:
		npac.Code = radius.AccessReject
		return npac
	}
}

func initRadius() {
	radiusServer = radius.NewServer(config.Radius.Listen, config.Radius.Secret, radiusService{})
	/* or you can convert it to a server that accept request from some hosts with different secrets
	cls := radius.NewClientList([]radius.Client{
		radius.NewClient("127.0.0.1", "secret1"),
		radius.NewClient("10.10.10.10", "secret2"),
	})
	s.WithClientList(cls)
	*/
}

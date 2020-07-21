package rest

import (
	"github.com/gorilla/mux"
	"github.com/julwil/bazo-client/client"
	"github.com/julwil/bazo-client/network"
	"github.com/julwil/bazo-miner/protocol"
	"math/big"
	"net/http"
)

func GetAccountEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	param := params["id"]

	logger.Printf("Incoming acc request for account: %v", param)

	var address [64]byte
	var addressHash [32]byte

	pubKeyInt, _ := new(big.Int).SetString(param, 16)

	if len(param) == 64 {
		copy(addressHash[:], pubKeyInt.Bytes())

		network.AccReq(false, addressHash)

		accI, _ := network.Fetch(network.AccChan)
		acc := accI.(*protocol.Account)

		address = acc.Address
	} else if len(param) == 128 {
		copy(address[:], pubKeyInt.Bytes())
		addressHash = protocol.SerializeHashContent(address)
	}

	acc, lastTenTx, err := client.GetAccount(address)
	if err != nil {
		SendJsonResponse(w, JsonResponse{http.StatusInternalServerError, err.Error(), nil})
	} else {
		var content []Content
		content = append(content, Content{"account", acc})

		for _, tx := range lastTenTx {
			if tx != nil {
				content = append(content, Content{"inbound", tx})
			}
		}

		SendJsonResponse(w, JsonResponse{http.StatusOK, "", content})
	}
}

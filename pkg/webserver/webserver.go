package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.zx2c4.com/wireguard/wgctrl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Generate a function serve http with specific port.
func Serve(addr, token, wgInterface, desiredPeerPubKey string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Get parameters from http request
		vcode := r.FormValue("t")
		if vcode == "" {
			vcode = r.FormValue("token")
		}
		if vcode != token {
			//response 403
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}

		ep, err := getWatchPeer(wgInterface, desiredPeerPubKey)
		if err != nil {
			fmt.Fprintf(w, "Error: %q", err.Error())
			return
		}
		fmt.Fprint(w, ep)
	})
	fmt.Printf("serving registry http at %v \n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getWatchPeer(wgInterface string, desiredPeerPubKey string) (endpoint string, err error) {
	endpoint = "not found"
	if wgInterface == "" {
		return "", status.Errorf(codes.Internal, "service interface is not set")
	}
	if desiredPeerPubKey == "" {
		return "", status.Errorf(codes.Internal, "peer public key is not set")
	}

	wgClient, err := wgctrl.New()
	if err != nil {
		log.Printf("error constructing Wireguard control client: %v", err)
		return "", status.Errorf(codes.Internal, err.Error())
	}

	dev, err := wgClient.Device(wgInterface)
	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}

	for _, p := range dev.Peers {
		if p.PublicKey.String() != strings.TrimSpace(desiredPeerPubKey) {
			continue
		}

		if p.Endpoint != nil {
			endpoint = p.Endpoint.IP.String()
			if p.Endpoint.Port > 0 {
				endpoint = endpoint + ":" + strconv.Itoa(p.Endpoint.Port)
			}
		}

	}

	return endpoint, nil

}

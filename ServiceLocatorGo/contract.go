package ServiceLocatorGo

import (
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

type ServiceLocatorGocontract interface {
	HelloWorld() (string, error)
}

type ServiceLocatorGo struct{}

// AddServices sets up and starts the service.
func AddServices() {
	ctx := context.Background()
	svc := ServiceLocatorGo{}

	ServiceLocatorGoHandler := httptransport.NewServer(
		ctx,
		makeHelloWorldEndpoint(svc),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/", ServiceLocatorGoHandler)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

package web

import (
	"fmt"
	"net/http"
)

// Query handles chaincode query requests.
func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Received Query request")
	queryParams := r.URL.Query()
	chainCodeName := queryParams.Get("chaincodeid")
	channelID := queryParams.Get("channelid")
	function := queryParams.Get("function")
	args := r.URL.Query()["args"]
	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	evaluateResponse, err := contract.EvaluateTransaction(function, args...)
	if err != nil {
		// fmt.Fprintf(w, "Error: %s", err)
		fmt.Println("error querying..")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error querying..."))
		return
	}
	fmt.Fprintf(w, "Response: %s", evaluateResponse)
}

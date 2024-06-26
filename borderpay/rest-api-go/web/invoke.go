package web

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// Invoke handles chaincode invoke requests.
func (setup *OrgSetup) Invoke(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Received Invoke request")
	if err := r.ParseForm(); err != nil {
		// fmt.Fprintf(w, "ParseForm() err: %s", err)
		fmt.Println("ParseForm() err: %s", err)
		w.Write([]byte("Error parsing form"))
		return
	}
	chainCodeName := r.FormValue("chaincodeid")
	channelID := r.FormValue("channelid")
	function := r.FormValue("function")
	args := r.Form["args"]
	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	txn_proposal, err := contract.NewProposal(function, client.WithArguments(args...))
	if err != nil {
		// fmt.Fprintf(w, "Error creating txn proposal: %s", err)
		fmt.Println("Error creating txn proposal")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating txn proposal"))
		return
	}
	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		// fmt.Fprintf(w, "Error endorsing txn: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error endorsing txn "))
		return
	}
	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		// fmt.Fprintf(w, "Error submitting transaction: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error submitting txn"))
		return
	}
	fmt.Fprintf(w, "Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())
}

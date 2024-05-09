package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var last_id string = "0"

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type Account struct {
	Account_No string  `json:"Account_No"`
	Bank_Name  string  `json:"Bank_Name"`
	Balance    float64 `json:"Balance"`
	Currency   string  `json:"Currency"`
}

type Bank struct {
	Bank_Name   string             `json:"Bank_Name"`
	Account_Map map[string]Account `json:"Account_Map"`
	Tax         float64            `json:"Tax"`
	Bank_Type   string             `json:"Bank_Type"`
	Currency    string             `json:"Currency"`
}

type Transaction struct {
	Transaction_ID  string  `json:"Transaction_ID"`
	Sender_Bank     string  `json:"Sender_Bank"`
	Sender_Acc_No   string  `json:"Sender_Acc_No"`
	Receiver_Bank   string  `json:"Receiver_Bank"`
	Receiver_Acc_No string  `json:"Receiver_Acc_No"`
	Amount          float64 `json:"Amount"`
	Currency        string  `json:"Currency"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// Define the list of banks with their initial account details
	banks := []Bank{
		{
			Bank_Name: "PNB",
			Account_Map: map[string]Account{
				"1": {Account_No: "1", Bank_Name: "PNB", Balance: 0, Currency: "INR"},
			},
			Tax:       0.05,
			Bank_Type: "Local",
			Currency:  "INR",
		},
		{
			Bank_Name: "SBI",
			Account_Map: map[string]Account{
				"1": {Account_No: "1", Bank_Name: "SBI", Balance: 0, Currency: "INR"},
			},
			Tax:       0.05,
			Bank_Type: "Local",
			Currency:  "INR",
		},
		{
			Bank_Name: "RBI",
			Account_Map: map[string]Account{
				"1": {Account_No: "1", Bank_Name: "RBI", Balance: 0, Currency: "INR"},
			},
			Tax:       0.05,
			Bank_Type: "National",
			Currency:  "INR",
		},
		{
			Bank_Name: "Exchange",
			Account_Map: map[string]Account{
				"1": {Account_No: "1", Bank_Name: "Exchange", Balance: 0, Currency: "USD"},
			},
			Tax:       0.05,
			Bank_Type: "Exchange",
			Currency:  "USD",
		},
		{
			Bank_Name: "BOA",
			Account_Map: map[string]Account{
				"1": {Account_No: "1", Bank_Name: "BOA", Balance: 0, Currency: "USD"},
			},
			Tax:       0.05,
			Bank_Type: "National",
			Currency:  "USD",
		},
		{
			Bank_Name: "BNY",
			Account_Map: map[string]Account{
				"1": {Account_No: "1", Bank_Name: "BNY", Balance: 0, Currency: "USD"},
			},
			Tax:       0.05,
			Bank_Type: "Local",
			Currency:  "USD",
		},
		{
			Bank_Name: "Barclays",
			Account_Map: map[string]Account{
				"1": {Account_No: "1", Bank_Name: "Barclays", Balance: 0, Currency: "USD"},
			},
			Tax:       0.05,
			Bank_Type: "Local",
			Currency:  "USD",
		},
	}

	// Iterate over each bank and its initial account details
	for _, bank := range banks {
		// Marshal bank details to JSON
		bankJSON, err := json.Marshal(bank)
		if err != nil {
			return err
		}

		// Store bank details in private data collection
		err = ctx.GetStub().PutPrivateData("Bank_Collections", bank.Bank_Name, bankJSON)
		if err != nil {
			return fmt.Errorf("failed to add initial banks in world state. %v", err)
		}
	}

	return nil
}

func (s *SmartContract) GetAllBanks(ctx contractapi.TransactionContextInterface) ([]Bank, error) {
	// Get an iterator over all the Banks in the ledger
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange("Bank_Collections", "", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// Create an empty slice to store the contracts
	var allB []Bank

	// Iterate over the contracts and add them to the slice
	for resultsIterator.HasNext() {
		bankJSON, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var b Bank
		err = json.Unmarshal(bankJSON.Value, &b)
		if err != nil {
			return nil, err
		}

		allB = append(allB, b)
	}

	return allB, nil
}

func (s *SmartContract) GetAllTransactions(ctx contractapi.TransactionContextInterface) ([]Transaction, error) {
	// Get an iterator over all the Banks in the ledger
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange("Transaction_Collections", "", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// Create an empty slice to store the contracts
	var allT []Transaction

	// Iterate over the contracts and add them to the slice
	for resultsIterator.HasNext() {
		transJSON, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var T Transaction
		err = json.Unmarshal(transJSON.Value, &T)
		if err != nil {
			return nil, err
		}

		allT = append(allT, T)
	}

	return allT, nil
}

func (s *SmartContract) CreateAccount(ctx contractapi.TransactionContextInterface, bank_name string, currency string, account_no string) (string, error) {
	bankJSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", bank_name)
	if err != nil {
		return "", fmt.Errorf("failed to get bank from world state. %v", err)
	}
	if bankJSON == nil {
		return "", fmt.Errorf("bank %s does not exist", bank_name)
	}
	var bank Bank
	err = json.Unmarshal(bankJSON, &bank)
	if err != nil {
		return "", err
	}
	account := Account{
		Account_No: account_no,
		Bank_Name:  bank_name,
		Balance:    1000000, //Initial Balance
		Currency:   currency,
	}
	bank.Account_Map[account.Account_No] = account
	bankJSON, err = json.Marshal(bank)
	if err != nil {
		return "", fmt.Errorf("failed to marshal bank: %v", err)
	}
	err = ctx.GetStub().PutPrivateData("Bank_Collections", bank_name, bankJSON)
	if err != nil {
		return "", fmt.Errorf("failed to update bank %s is world state: %v", bank_name, err)
	}
	return "Account added on Hyperledger", nil
}

func execute(ctx contractapi.TransactionContextInterface, src Bank, src_bank_acc_no string, receiver Bank, receiver_bank_acc_no string, amount float64, rate float64 /*conversion b/w curr*/) (string, error) { //executes transactions after all checks are done
	//Adding amount to receiver's Account
	fmt.Printf("The src bank is %s\n", src.Bank_Name)
	fmt.Printf("The src bank is %s\n", receiver.Bank_Name)
	account_rec := receiver.Account_Map[receiver_bank_acc_no]
	account_rec.Balance += (float64(amount) * rate)
	receiver.Account_Map[receiver_bank_acc_no] = account_rec
	//Subtracting amount from sender's Accountf
	account_send := src.Account_Map[src_bank_acc_no]
	account_send.Balance -= amount
	src.Account_Map[src_bank_acc_no] = account_send

	fmt.Printf("The value of rc %f\n", src.Account_Map[src_bank_acc_no].Balance)
	//Updating changes in world state
	srcJSON, err := json.Marshal(src)
	if err != nil {
		return "", fmt.Errorf("failed to marshal src bank: %v", err)
	}
	err = ctx.GetStub().PutPrivateData("Bank_Collections", src.Bank_Name, srcJSON)
	if err != nil {
		return "", fmt.Errorf("failed to update src bank %s is world state: %v", src.Bank_Name, err)
	}
	receiverJSON, err := json.Marshal(receiver)
	if err != nil {
		return "", fmt.Errorf("failed to marshal receiver bank: %v", err)
	}
	err = ctx.GetStub().PutPrivateData("Bank_Collections", receiver.Bank_Name, receiverJSON)
	if err != nil {
		return "", fmt.Errorf("failed to update receiver bank %s is world state: %v", receiver.Bank_Name, err)
	}
	fmt.Printf("New_amt is:4\n")

	//Creating Transaction ID
	transaction_id, err := strconv.Atoi(string(last_id))
	var li int
	li, err = strconv.Atoi(string(last_id))
	li += 1
	last_id = strconv.Itoa(li)
	if err != nil {
		return "", fmt.Errorf("failed to convert last transaction id to int: %v", err)
	}
	transaction_id += 1
	//Creating Transaction
	transaction := Transaction{
		Transaction_ID:  strconv.Itoa(transaction_id),
		Sender_Bank:     src.Bank_Name,
		Sender_Acc_No:   src_bank_acc_no,
		Receiver_Bank:   receiver.Bank_Name,
		Receiver_Acc_No: receiver_bank_acc_no,
		Amount:          amount,
		Currency:        src.Currency,
	}
	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return "", fmt.Errorf("failed to marshal transaction: %v", err)
	}
	err = ctx.GetStub().PutPrivateData("Transaction_Collections", strconv.Itoa(transaction_id), transactionJSON)
	if err != nil {
		return "", fmt.Errorf("failed to update transaction %s is world state: %v", strconv.Itoa(transaction_id), err)
	}

	//Updating last transaction id
	err = ctx.GetStub().PutPrivateData("Transaction_Collections", "last_id", []byte(strconv.Itoa(transaction_id)))
	if err != nil {
		return "", fmt.Errorf("failed to update last transaction id in world state but transaction pushed: %v", err)
	}
	return strconv.Itoa(transaction_id), nil
}

func (s *SmartContract) MakePayment(ctx contractapi.TransactionContextInterface, src_bank string, src_bank_acc_no string, receiver_bank string, receiver_bank_acc_no string, amount float64) (string, error) {
	// checking sender's account details
	if amount < 0 {
		return "", fmt.Errorf("-ve transactions not allowed")
	}
	src_bank_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", src_bank)
	if err != nil {
		return "", fmt.Errorf("failed to connect to sender's bank.%s: %v", src_bank, err)
	}
	if src_bank_JSON == nil {
		return "", fmt.Errorf("Bank %s does not exist", src_bank)
	}
	var src Bank
	err = json.Unmarshal(src_bank_JSON, &src)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal sendor bank %s details : %v", src_bank, err)
	}
	// checking if sender's account is present
	if _, ok := src.Account_Map[src_bank_acc_no]; !ok {
		// create sender's account
		_, err = s.CreateAccount(ctx, src.Bank_Name, src.Currency, src_bank_acc_no)
		// fmt.Println("\nAccount Created\n")
		if err != nil {
			return "", fmt.Errorf("failed to create sender's account: %v", err)
		}
	}

	//to be modified

	var curr_src string = src.Currency

	//checking receiver's account details
	receiver_bank_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", receiver_bank)
	if err != nil {
		return "", fmt.Errorf("failed to connect to receiver's bank %s: %v", receiver_bank, err)
	}
	if receiver_bank_JSON == nil {
		return "", fmt.Errorf("bank %s does not exist", receiver_bank)
	}
	var receiver Bank
	err = json.Unmarshal(receiver_bank_JSON, &receiver)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal receiver bank %s details: %v", receiver_bank, err)
	}

	// checking if receiver's account is present
	if _, ok := receiver.Account_Map[receiver_bank_acc_no]; !ok {
		// create receiver's account
		_, err = s.CreateAccount(ctx, receiver.Bank_Name, receiver.Currency, receiver_bank_acc_no)
		if err != nil {
			return "", fmt.Errorf("failed to create receiver's account: %v", err)
		}
	}

	// var curr_receiver string = receiver.Currency
	var balance = src.Account_Map[src_bank_acc_no].Balance
	if src.Currency == receiver.Currency {
		if amount < 0 {
			return "", fmt.Errorf("-ve transactions not allowed")
		}
		if balance < amount {
			return "", fmt.Errorf("insufficient Balance in Sender's Account")
		} else {
			var trans string
			//first I am sending from src bank to central bank all amt.
			if curr_src == "USD" {
				boa_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", "BOA")
				if err != nil {
					return "", fmt.Errorf("failed to connect to america's central bank.%s: %v", "BOA", err)
				}
				var boa Bank //getting BOA data
				err = json.Unmarshal(boa_JSON, &boa)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", "boa", err)
				}

				_, err = execute(ctx, src, src_bank_acc_no, boa, "1", amount, 1)
				if err != nil {
					return "", fmt.Errorf("failed to execute transaction from %s to boa : %v", src_bank, err)
				}
				var new_amt float64 = amount * (1 - .1)
				trans, err = execute(ctx, boa, "1", receiver, receiver_bank_acc_no, new_amt, 1)
				if err != nil {
					return "", fmt.Errorf("failed to execute transaction from boa to %s : %v", receiver_bank, err)
				}
				return trans, nil
			} else {
				rbi_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", "RBI")
				if err != nil {
					return "", fmt.Errorf("failed to connect to india's central bank.%s: %v", "rbi", err)
				}
				var rbi Bank //getting rbi data
				err = json.Unmarshal(rbi_JSON, &rbi)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", "rbi", err)
				}

				_, err = execute(ctx, src, src_bank_acc_no, rbi, "1", amount, 1)

				if err != nil {
					return "", fmt.Errorf("failed to execute transaction from %s to rbi : %v", src_bank, err)
				}
				fmt.Printf("New_amt is:1\n")

				var new_amt float64 = amount * (1.0 - .1)
				fmt.Printf("New_amt is %f\n", new_amt)
				trans, err = execute(ctx, rbi, "1", receiver, receiver_bank_acc_no, new_amt, 1)

				if err != nil {
					return "", fmt.Errorf("failed to execute transaction from boa to %s : %v", receiver_bank, err)
				}
				fmt.Printf("New_amt is:3\n")
				return trans, nil
			}
		}
	} else {
		if curr_src == "USD" {
			if amount < 0 {
				return "", fmt.Errorf("-ve transactions not allowed")
			}
			if balance < amount {
				return "", fmt.Errorf("insufficient Balance in Sender's Account")
			} else {
				boa_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", "BOA")
				if err != nil {
					return "", fmt.Errorf("failed to connect to america's central bank.%s: %v", "BOA", err)
				}
				var boa Bank //getting BOA data
				err = json.Unmarshal(boa_JSON, &boa)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", "boa", err)
				}

				_, err = execute(ctx, src, src_bank_acc_no, boa, "1", amount, 1)
				if err != nil {
					return "", err
				}

				exchange_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", "Exchange")
				if err != nil {
					return "", fmt.Errorf("failed to connect to foreign exchange.%s: %v", "BOA", err)
				}
				var exchange Bank //getting Exchange data
				err = json.Unmarshal(exchange_JSON, &exchange)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", "exchange", err)
				}
				_, err = execute(ctx, boa, "1", exchange, "1", amount, 1)
				if err != nil {
					return "", err
				}
				amount = amount * (1 - .01)
				rbi_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", "RBI")
				if err != nil {
					return "", fmt.Errorf("failed to connect to india's central bank.%s: %v", "RBI", err)
				}
				var rbi Bank //getting RBI data
				err = json.Unmarshal(rbi_JSON, &rbi)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", "rbi", err)
				}
				_, err = execute(ctx, exchange, "1", rbi, "1", amount, 80)
				if err != nil {
					return "", err
				}
				amount = amount * 80
				amount = amount * (1 - .1)
				rec_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", receiver_bank)
				if err != nil {
					return "", fmt.Errorf("failed to connect to .%s: %v", receiver_bank, err)
				}
				var temp Bank //getting final reciever data
				err = json.Unmarshal(rec_JSON, &temp)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", receiver_bank, err)
				}
				var trans_4 string

				trans_4, err = execute(ctx, rbi, "1", temp, receiver_bank_acc_no, amount, 1)
				if err != nil {
					return "", err
				}
				return trans_4, nil
			}
		}

		if curr_src == "INR" {
			if balance < amount {
				return "", fmt.Errorf("insufficient Balance in Sender's Account")
			} else {
				rbi_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", "RBI")
				if err != nil {
					return "", fmt.Errorf("failed to connect to india's central bank.%s: %v", "RBI", err)
				}
				var rbi Bank //getting RBI data
				err = json.Unmarshal(rbi_JSON, &rbi)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", "rbi", err)
				}
				_, err = execute(ctx, src, src_bank_acc_no, rbi, "1", amount, 1)
				if err != nil {
					return "", err
				}

				exchange_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", "Exchange")
				if err != nil {
					return "", fmt.Errorf("failed to connect to foreign exchange.%s: %v", "Exchange", err)
				}
				var exchange Bank //getting Exchange data
				err = json.Unmarshal(exchange_JSON, &exchange)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", "exchange", err)
				}
				_, err = execute(ctx, rbi, "1", exchange, "1", amount, 1.0/float64(80))
				if err != nil {
					return "", err
				}
				amount /= 80
				boa_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", "BOA")
				if err != nil {
					return "", fmt.Errorf("failed to connect to america's central bank.%s: %v", "BOA", err)
				}
				var boa Bank //getting BOA data
				err = json.Unmarshal(boa_JSON, &boa)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", "boa", err)
				}
				amount = amount * (1 - .01)
				_, err = execute(ctx, exchange, "1", boa, "1", amount, 1)
				if err != nil {
					return "", err
				}

				rec_JSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", receiver_bank)
				if err != nil {
					return "", fmt.Errorf("failed to connect to .%s: %v", receiver_bank, err)
				}
				var temp Bank //getting final receiver data
				err = json.Unmarshal(rec_JSON, &temp)
				if err != nil {
					return "", fmt.Errorf("failed to unmarshal bank %s details : %v", receiver_bank, err)
				}
				var trans_8 string
				amount = amount * (1 - .1)
				trans_8, err = execute(ctx, boa, "1", temp, receiver_bank_acc_no, amount, 1)
				if err != nil {
					return "", err
				}
				return trans_8, nil
			}
		}
	}
	return "", nil
}

func (s *SmartContract) ViewBalance(ctx contractapi.TransactionContextInterface, bank_name string, acc_no string) (float64, error) {
	bankJSON, err := ctx.GetStub().GetPrivateData("Bank_Collections", bank_name)
	if err != nil {
		return 0, fmt.Errorf("failed to get bank details: %v", err)
	}
	if bankJSON == nil {
		return 0, fmt.Errorf("bank %s does not exist", bank_name)
	}

	var bank Bank
	err = json.Unmarshal(bankJSON, &bank)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal bank details: %v", err)
	}

	account, ok := bank.Account_Map[acc_no]
	if !ok {
		return 0, fmt.Errorf("account %s does not exist in bank %s", acc_no, bank_name)
	}

	return account.Balance, nil
}

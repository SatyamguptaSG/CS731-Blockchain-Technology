package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var idx string = "1"

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

//Create user,[ID,PASSWORD,USER TYPE,ACCOUNT BANK NAME,ACCOUNT NUMBER,]   //Signup

//Login[ID,PASSWORD,USER TYPE]

//Struct Contract[CONTRACT ID,EMPLOYER ID,EMPLOYEE ID,BANK NAME,ACC NUMBER(BOTH),PAYMENT,DURATION,STATUS]

//Struct User[ID,PASSWORD,USER TYPE,ACCOUNT BANK NAME,ACCOUNT NUMBER,]

//Create contract,[EMPLOYEE ID,PAYMENT,DURATION,EMPLOYER ID]

//Approve contract[CONTRACT ID,EMPLOYEE ID,]
//WILL CHANGE STAUS

//Revkoe Contract[CONTRACT ID,USERtype,UserId]//Status rejected

//APPROVE PAYMENT[EMPLOYER ID,CONTRACT ID,]

//View Contracts [CONTRACT ID]

//Fetch all Contracts [USER ID]

// Asset describes basic details of what makes up a simple asset
//Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically

type transaction struct {
	Transaction_ID       string `json:"Transaction_ID"`
	Src_Bank             string `json:"Source_Bank"`
	Src_Bank_Account_No  string `json:"Source_Bank_Account_No"`
	Dest_Bank            string `json:"Receiver_Bank"`
	Dest_Bank_Account_No string `json:"Receiver_Bank_Account_No"`
	Amount               int    `json:"Amount"`
	Origin_Time_Stamp    string `json:"Origin_Time_Stamp"`
	Status               string `json:"Status"`
	Update_Time_Stamp    string `json:"Update_Time_Stamp"`
}

type Contract struct {
	ContractID                 string        `json:"ContractID"`
	EmployerId                 string        `json:"EmployerId"`
	EmployeeId                 string        `json:"EmployeeId"`
	Payment                    int           `json:"Payment"`
	Duration                   int           `json:"Duration"`
	BankName_Employer          string        `json:"BankName_Employer"`
	BankAccountNumber_Employer string        `json:"BankAccountNumber_Employer"`
	BankName_Employee          string        `json:"BankName_Employee"`
	BankAccountNumber_Employee string        `json:"BankAccountNumber_Employee"`
	STATUS                     string        `json:"STATUS"`
	Last_Payment_Date          string        `json:"Last_Payment_Date"`
	All_Transactions           []transaction `json:"All_Transactions"`
}

type User struct {
	UserID            string `json:"UserID"`
	Password          string `json:"Password"`
	UserType          int    `json:"UserType"`
	BankName          string `json:"BankName"`
	BankAccountNumber string `json:"BankAccountNumber"`
}

// Usertype 0 => Employee
// Usertype 1 => Employer

//Functions related to User

func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, id string, psw string, bname string, baccNo string, t int) error {
	// Check if user with provided id already exists
	existingUserJSON, err := ctx.GetStub().GetPrivateData("UsersCollection", id)
	if err != nil {
		return err
	}
	if existingUserJSON != nil {
		return fmt.Errorf("user with id %s already exists, register with some different id", id)
	}

	asset := User{
		UserID:            id,
		Password:          psw,
		UserType:          t,
		BankName:          bname,
		BankAccountNumber: baccNo,
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutPrivateData("UsersCollection", id, assetJSON)
}

// Login Function
func (s *SmartContract) LoginFunc(ctx contractapi.TransactionContextInterface, id string, psw string, Utype int) (bool, error) {
	userJson, err := ctx.GetStub().GetPrivateData("UsersCollection", id)

	if err != nil {
		return false, err
	}
	if userJson == nil {
		return false, fmt.Errorf("no such user present")
	}
	var us User
	err = json.Unmarshal(userJson, &us)
	if err != nil {
		return false, err
	}
	if psw != us.Password {
		return false, fmt.Errorf("incorrect password")
	} else if Utype != us.UserType {
		return false, fmt.Errorf("usertype mistaked")
	}
	return true, nil
}

func (s *SmartContract) GetAllUsers(ctx contractapi.TransactionContextInterface) ([]User, error) { //api to be eddited
	// Create an empty slice to store the users
	var users []User

	// Get an iterator over all the users in the ledger
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange("UsersCollection", "", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// Iterate over the users and add them to the slice
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var user User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

//Functions related to Contracts

func (s *SmartContract) CreateContract(ctx contractapi.TransactionContextInterface, employerId string, employeeId string, payment int, duration int) error {
	// Fetch the rest of the parameters using employeeId and employerId
	fmt.Println("Inside CreateContract")
	employee, err := ctx.GetStub().GetPrivateData("UsersCollection", employeeId)
	if err != nil {
		fmt.Println("Error in fetching employee data")
		return err
	}
	if employee == nil {
		fmt.Println("No such employee with given employeeId present")
		return fmt.Errorf("no such employee with given employeeId present") // No such employee present
	}
	var t1 User
	err = json.Unmarshal(employee, &t1)
	if err != nil {
		fmt.Println("Error in unmarshalling employee data")
		return err
	}
	if t1.UserType == 1 {
		fmt.Println("Given employeeId is not of an employee, rather of a employer")
		return fmt.Errorf("given employeeid is not of an employee, rather of a employer")
	}
	employer, err := ctx.GetStub().GetPrivateData("UsersCollection", employerId)
	fmt.Println("Employer data fetched")
	if err != nil {
		fmt.Println("Error in fetching employer data")
		return err
	}
	if employer == nil {
		fmt.Println("No such employer with given employerid present")
		return fmt.Errorf("no such employer with given employerid present") // No such employer present
	}
	fmt.Println("Employer data fetched")
	var t2 User
	err = json.Unmarshal(employer, &t2)
	if err != nil {
		return err
	}

	ct := Contract{
		ContractID:                 idx,
		EmployerId:                 employerId,
		EmployeeId:                 employeeId,
		Payment:                    payment,
		Duration:                   duration,
		BankName_Employer:          t2.BankName,
		BankAccountNumber_Employer: t2.BankAccountNumber,
		BankName_Employee:          t1.BankName,
		BankAccountNumber_Employee: t1.BankAccountNumber,
		STATUS:                     "Pending",
		Last_Payment_Date:          "",
		All_Transactions:           []transaction{},
	}

	ctJson, err := json.Marshal(ct)
	if err != nil {
		fmt.Println("Error in marshalling contract data")
		return err
	}

	var new_idx string = idx
	val, _ := strconv.Atoi(new_idx)
	val = val + 1
	idx = strconv.Itoa(val)
	fmt.Println("Contract data marshalled")
	// infoln("Contract data marshalleddsnkjdsn cj 32889");
	return ctx.GetStub().PutPrivateData("ContractsCollection", new_idx, ctJson)
}

func (s *SmartContract) AcceptContract(ctx contractapi.TransactionContextInterface, contractID string) error {
	fmt.Println("Inside AcceptContract")
	contractJSON, err := ctx.GetStub().GetPrivateData("ContractsCollection", contractID)
	if err != nil {
		fmt.Println("Error in getting contract JSON:", err)
		return fmt.Errorf("failed to read contract from world state: %v", err)
	}
	if contractJSON == nil {
		fmt.Println("Contract JSON is nil")
		return fmt.Errorf("the contract with contract id %s does not exist", contractID)
	}

	var contract Contract
	err = json.Unmarshal(contractJSON, &contract)
	if err != nil {
		fmt.Println("Error in unmarshalling contract JSON:", err)
		return err
	}
	currStatus := contract.STATUS
	if currStatus == "Revoked" {
		fmt.Println("Cannot accept a revoked contract")
		return fmt.Errorf("cannot accept an already revoked contract") // Fix: Change error string to lowercase
	} else if currStatus == "Accepted" {
		fmt.Println("Cannot accept an already accepted contract")
		return fmt.Errorf("already accepted") // Fix: Change error string to lowercase
	} else {
		contract.STATUS = "Accepted"
		ctJson, err := json.Marshal(contract)
		if err != nil {
			fmt.Println("Error in marshalling contract JSON:", err)
			return err
		}
		fmt.Println("Contract JSON marshalled")
		return ctx.GetStub().PutPrivateData("ContractsCollection", contractID, ctJson)
	}
}

func (s *SmartContract) RevokeContract(ctx contractapi.TransactionContextInterface, id_of_Contract string) error {
	fmt.Println("Inside RevokeContract")
	ctJson, err := ctx.GetStub().GetPrivateData("ContractsCollection", id_of_Contract)
	if err != nil {
		fmt.Println("Error in getting contract JSON:", err)
		return err
	}
	if ctJson == nil {
		fmt.Println("Contract JSON is nil")
		return fmt.Errorf("contract not found")
	}
	var ct Contract
	err = json.Unmarshal(ctJson, &ct)
	if err != nil {
		fmt.Println("Error in unmarshalling contract JSON:", err)
		return err
	}
	currStatus := ct.STATUS
	if currStatus == "Revoked" {
		fmt.Println("Cannot revoke an already revoked contract")
		return fmt.Errorf("cannot revoke an already revoked contract") // Fix: Change error string to lowercase
	} else if currStatus == "Accepted" || currStatus == "Pending" {
		ct.STATUS = "Revoked"
		ctJson, err := json.Marshal(ct)
		if err != nil {
			fmt.Println("Error in marshalling contract JSON:", err)
			return err
		}
		return ctx.GetStub().PutPrivateData("ContractsCollection", id_of_Contract, ctJson)
	}
	return nil
}

// func (s *SmartContract) RevokeContract(ctx contractapi.TransactionContextInterface, id_of_Contract string) error {
// 	ctJson, err := ctx.GetStub().GetPrivateData("ContractsCollection", id_of_Contract)
// 	if err != nil {
// 		return err
// 	}
// 	if ctJson == nil {
// 		return err
// 	}
// 	var ct Contract
// 	err = json.Unmarshal(ctJson, &ct)
// 	if err != nil {
// 		return err
// 	}
// 	currStatus := ct.STATUS
// 	if currStatus == "Revoked" {
// 		return fmt.Errorf("cannot revoke an already revoked contract") // Fix: Change error string to lowercase
// 	} else if currStatus == "Accepted" || currStatus == "Pending" {
// 		ct.STATUS = "Revoked"
// 		ctJson, err := json.Marshal(ct)
// 		if err != nil {
// 			return err
// 		}
// 		return ctx.GetStub().PutPrivateData("ContractsCollection", id_of_Contract, ctJson)
// 	}

// 	return nil
// }

func (s *SmartContract) ViewAllContracts(ctx contractapi.TransactionContextInterface) ([]Contract, error) {
	// Get an iterator over all the contracts in the ledger
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange("ContractsCollection", "", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// Create an empty slice to store the contracts
	var contracts []Contract

	// Iterate over the contracts and add them to the slice
	for resultsIterator.HasNext() {
		contractJSON, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var contract Contract
		err = json.Unmarshal(contractJSON.Value, &contract)
		if err != nil {
			return nil, err
		}

		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (s *SmartContract) FetchAllContractsbyUserID(ctx contractapi.TransactionContextInterface, UserID string) ([]Contract, error) {

	userJson, err := ctx.GetStub().GetPrivateData("UsersCollection", UserID)
	if err != nil {
		return nil, err
	}
	if userJson == nil {
		return nil, fmt.Errorf("user not found")
	}
	var user User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		return nil, err
	}
	UserType := user.UserType
	var user_contracts []Contract

	if UserType == 0 {
		// Query to fetch all contracts with given UserID as EmployeeID
		queryString := fmt.Sprintf(`{"selector":{"EmployeeId":"%s"}}`, UserID)
		resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("ContractsCollection", queryString)
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()

		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return nil, err
			}
			var ct Contract
			err = json.Unmarshal(queryResponse.Value, &ct)
			if err != nil {
				return nil, err
			}
			user_contracts = append(user_contracts, ct)
		}
	} else if UserType == 1 {
		// Query to fetch all contracts with given UserID as EmployerId
		queryString := fmt.Sprintf(`{"selector":{"EmployerId":"%s"}}`, UserID)
		resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("ContractsCollection", queryString)
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()

		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return nil, err
			}
			var ct Contract
			err = json.Unmarshal(queryResponse.Value, &ct)
			if err != nil {
				return nil, err
			}
			user_contracts = append(user_contracts, ct)
		}
	} else {
		return nil, fmt.Errorf("invalid UserType")
	}
	return user_contracts, nil
}

func (s *SmartContract) UpdateContract(ctx contractapi.TransactionContextInterface, transac transaction, con Contract) error {
	//To be tested when we do make payment contract
	contractJson, err := ctx.GetStub().GetPrivateData("ContractsCollection", con.ContractID)

	if err != nil {
		return err
	}
	if contractJson == nil {
		return err
	}
	var coJson []byte

	var co Contract
	id_of_Contract := con.ContractID

	err = json.Unmarshal(contractJson, &co)
	if err != nil {
		return err
	}

	co.All_Transactions = append(co.All_Transactions, transac)
	coJson, err = json.Marshal(co)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutPrivateData("ContractsCollection", id_of_Contract, coJson)
}

// Given user id this function returns the bank details of the user
func (s *SmartContract) GetBankDetails(ctx contractapi.TransactionContextInterface, id string) ([]User, error) {
	fmt.Println("Inside GetBankDetails")
	var users []User
	userJSON, err := ctx.GetStub().GetPrivateData("UsersCollection", id)
	if err != nil {
		return nil, err
	}
	if userJSON == nil {
		fmt.Println("No such user present")
		return nil, fmt.Errorf("no such user present")
	}
	var user User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		fmt.Println("Error in unmarshalling user data")
		return nil, err
	}
	users = append(users, user)
	return users, nil
}

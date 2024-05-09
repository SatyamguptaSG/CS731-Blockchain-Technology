# BorderPay.Io

## Pre-requisites for Hyperledger Fabric

Before setting up this project, ensure you have the following pre-requisites installed:

- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go Programming Language (version >= 1.20)](https://golang.org/doc/install)
- [Node.js (version >= 12.x)](https://nodejs.org/en/download/)
- [npm](https://www.npmjs.com/get-npm)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

For detailed installation instructions of HyperLedger -Fabric, refer to the [Hyperledger Fabric documentation](https://hyperledger-fabric.readthedocs.io/en/latest/getting_started.html).

## Description

This repository contains the following folder:

1. `Blockchain_frontend` - Contains frontend part of the application
2. `borderpay` - it contains chaincodes and API-server in `rest-api-go` folder
3. `CS731_Project_Report` - It contains Project Report 

## Frontend Installation

To install and run the frontend, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/havibohra/BorderPay.Io.git
   ```

2. Navigate to the frontend directory:

   ```bash
   cd frontend
   ```

3. Install dependencies:

   ```bash
   npm install
   ```

4. Start the frontend server:

   ```bash
   npm start
   ```
5.Access the application in your browser at `http://localhost:3000`.
## Setup Details

Follow these steps to set up and run the project:

1. Copy the `borderpay` folder from the repo inside the folder `fabric-samples` of Hyperledger-Fabric
2. 
```bash
#To avoid any permissions problem:
   sudo bash
# navigate to test-network folder in fabric-samples

# Start with creating channel named mychannel and setting Peer state database to deploy: couchdb with ca

./network.sh up  createChannel -c mychannel -ca -s couchdb

# deploying chaincodes to peers 

./network.sh deployCC -ccn basictest  -ccp ../borderpay/chaincode-go  -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"  -cccg '../borderpay/chaincode-go/collections_config.json' -ccep "OR('Org1MSP.peer','Org2MSP.peer')"

./network.sh deployCC -ccn paytest  -ccp ../borderpay/chaincode2-go  -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"  -cccg '../borderpay/chaincode2-go/collections2_config.json' -ccep "OR('Org1MSP.peer','Org2MSP.peer')"

# required environment variables setup

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

# Environment variables for Org1, login as peer0.org1

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n paytest --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

#Start a new terminal

# navigate to rest-api-go folder(run this command if you're in test-network currently)
cd ../borderpay/rest-api-go
# start api-server (It will start api-sever on port:3002)
go run main.go
   ```

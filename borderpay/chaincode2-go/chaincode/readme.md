# navigate to test-network folder

# Start with creating channel named mychannel and setting Peer state database to deploy: couchdb

sudo ./network.sh up  createChannel -c mychannel -ca -s couchdb

# deploying companysidecode to peers 

sudo ./network.sh deployCC -ccn basictest  -ccp ../borderpay/chaincode-go  -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"  -cccg '../borderpay/chaincode-go/collections_config.json' -ccep "OR('Org1MSP.peer','Org2MSP.peer')"


# required environment variables setup
sudo bash

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

# Environment variables for Org1, login as peer0.org1

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# Sample Query
peer chaincode query -C mychannel -n basictest -c '{"Args":["GetAllUsers"]}'

# Sample invoke

# create user
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basictest --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"CreateUser","Args":["Havi","abc123","SBI","123456","0"]}'

# create contract
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basictest --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"CreateContract","Args":["Haviboss","Havi","5000","12"]}'

https://youtu.be/aFfAQHXsYIs?si=mSdjrCuIDACzswqi


# navigate to api-folder 
cd ../borderpay/rest-api-go
# start api-server (port =3002 currently in my computer)
go run main.go


# Invoke request (POST request)

# CreateUser
http://localhost:3002/invoke?channelid=mychannel&chaincodeid=basictest&function=CreateUser&args=Havi&args=abc123&args=SBI&args=123456&args=0

http://localhost:3002/invoke?channelid=mychannel&chaincodeid=basictest&function=CreateUser&args=Haviboss&args=def123&args=ICICI&args=456&args=1

# CreateContract
http://localhost:3002/invoke?channelid=mychannel&chancodeid=basictest&function=CreateContract&args=Haviboss&args=Havi&args=200&args=1

# Query (GET request)

# GetAllUsers 
http://localhost:3002/query?channelid=mychannel&chaincodeid=basictest&function=GetAllUsers

# LoginFunc
http://localhost:3002/invoke?channelid=mychannel&chaincodeid=basictest&function=LoginFunc&args=Havi&args=abc123&args=0


peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basictest --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"LoginFunc","Args":["Havi","abc123","0"]}'
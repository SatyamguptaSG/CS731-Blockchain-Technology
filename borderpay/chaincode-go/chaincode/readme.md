# navigate to test-network folder
sudo bash

# Start with creating channel named mychannel and setting Peer state database to deploy: couchdb

./network.sh up  createChannel -c mychannel -ca -s couchdb

# deploying companysidecode to peers 

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

# Environment variables for Org2, login as peer0.org2

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

# Sample Query
peer chaincode query -C mychannel -n basictest -c '{"Args":["GetAllUsers"]}'

peer chaincode query -C mychannel -n paytest -c '{"Args":["GetAllBanks"]}'

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
http://localhost:3002/invoke?channelid=mychannel&chaincodeid=basictest&function=CreateContract&args=Haviboss&args=Havi&args=200&args=1

# Query (GET request)

# GetAllUsers 
http://localhost:3002/query?channelid=mychannel&chaincodeid=basictest&function=GetAllUsers

# ViewAllContracts
http://localhost:3002/query?channelid=mychannel&chaincodeid=basictest&function=ViewAllContracts

# LoginFunc
http://localhost:3002/invoke?channelid=mychannel&chaincodeid=basictest&function=LoginFunc&args=Havi&args=abc123&args=0


peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basictest --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"LoginFunc","Args":["Havi","abc123","0"]}'




# payment chaincode

# InitLedger (maybe we can includmple.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n paytest --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles e this in deploying time) (PNB,RBI,Exchange,BOA,BNY)
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.exa"${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

# Make Payment
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n paytest --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"MakePayment","Args":["PNB","456","RBI","]}'

const response = await axios.get('http://localhost:3002/query', {
          params: {
          'channelid': 'mychannel',
          'chaincodeid': 'basictest',
          'function': 'FetchAllContractsbyUserID',
          'args': 'Havi'
          }
        });

const response = await axios.post(
        'http://localhost:3002/invoke',
        new URLSearchParams([
          ['', ''],
          ['channelid', 'mychannel'],
          ['chaincodeid', 'basictest'],
          ['function', 'RevokeContract'],
          ['args', contract_id],
        ])
      );
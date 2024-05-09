# BorderPay.Io

#Navigate to test-network
'sudo bash'

#Clone this repo

#deploying companyside to peer
./network.sh deployCC -ccn basictest  -ccp path/chaincode-go  -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"  -cccg '../borderpay/chaincode-go/collections_config.json' -ccep "OR('Org1MSP.peer','Org2MSP.peer')"

package main

import (
	"github.com/hyperledger/myfabric/blockchian"
	"log"
)
var (
	peer0Org1 = "peer0.org1.example.com"
	peer0Org2 = "peer0.org2.example.com"
)
func main() {
	org1Client := blockchian.New("./config/org1-config.yaml", "Org1", "Admin", "User1")
	org2Client := blockchian.New("./config/org2-config.yaml", "Org2", "Admin", "User1")
	defer org1Client.SDK.Close()
	defer org2Client.SDK.Close()
	// Install, instantiate, invoke, query
	//Phase1(org1Client, org2Client)
	//// Install, upgrade, invoke, query
	//Phase2(org1Client, org2Client)
	//org1Client.InstantiateCC("v1", peer0Org1)
	//org1Client.InvokeCC([]string{"peer0.org1.example.com",
	//	"peer0.org2.example.com"})
	//org1Client.InvokeCC([]string{peer0Org1})
	org1Client.QueryCC("peer0.org2.example.com", "a")
}

func Phase1(cli1, cli2 *blockchian.Client) {
	log.Println("=================== Phase 1 begin ===================")
	defer log.Println("=================== Phase 1 end ===================")

	if err := cli1.InstallCC("v1", peer0Org1); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org1's peers")

	if err := cli2.InstallCC("v1", peer0Org2); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org2's peers")

	// InstantiateCC chaincode only need once for each channel
	if _, err := cli1.InstantiateCC("v1", peer0Org1); err != nil {
		log.Panicf("Instantiated chaincode error: %v", err)
	}
	log.Println("Chaincode has been instantiated")

	if _, err := cli1.InvokeCC([]string{peer0Org1}); err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	if err := cli1.QueryCC("peer0.org1.example.com", "a"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org1")
}

func Phase2(cli1, cli2 *blockchian.Client) {
	log.Println("=================== Phase 2 begin ===================")
	defer log.Println("=================== Phase 2 end ===================")

	v := "v2"

	// Install new version chaincode
	if err := cli1.InstallCC(v, peer0Org1); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org1's peers")

	if err := cli2.InstallCC(v, peer0Org2); err != nil {
		log.Panicf("Intall chaincode error: %v", err)
	}
	log.Println("Chaincode has been installed on org2's peers")

	// Upgrade chaincode only need once for each channel
	if err := cli1.UpgradeCC(v, peer0Org2); err != nil {
		log.Panicf("Upgrade chaincode error: %v", err)
	}
	log.Println("Upgrade chaincode success for channel")

	if _, err := cli1.InvokeCC([]string{"peer0.org1.example.com",
		"peer0.org2.example.com"}); err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
	}
	log.Println("Invoke chaincode success")

	if err := cli1.QueryCC("peer0.org2.example.com", "a"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	log.Println("Query chaincode success on peer0.org2")
}


package blockchian

import (
	"log"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	_ "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type Client struct {
	// Fabric network information
	ConfigPath string
	OrgName    string
	OrgAdmin   string
	OrgUser    string

	// sdk clients
	SDK *fabsdk.FabricSDK
	resMgtClient  *resmgmt.Client
	channelClient  *channel.Client

	// Same for each peer
	ChannelID string
	CCID      string // chaincode ID, eq name
	CCPath    string // chaincode source path, 是GOPATH下的某个目录
	CCGoPath  string // GOPATH used for chaincode
}

func New(cfg, org, admin, user string) *Client {
	c := &Client{
		ConfigPath: cfg,
		OrgName:    org,
		OrgAdmin:   admin,
		OrgUser:    user,

		CCID:      "example2",
		CCPath:    "github.com/hyperledger/fabric-samples/chaincode/chaincode_example02/go/", // 相对路径是从GOPAHT/src开始的
		CCGoPath:  os.Getenv("GOPATH"),
		ChannelID: "mychannel",
	}

	// create sdk
	sdk, err := fabsdk.New(config.FromFile(c.ConfigPath))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}
	c.SDK = sdk
	log.Println("Initialized fabric sdk")

	rcp := sdk.Context(fabsdk.WithUser(c.OrgAdmin), fabsdk.WithOrg(c.OrgName))
	c.resMgtClient, err = resmgmt.New(rcp)
	if err != nil {
		log.Panicf("failed to create resource client: %s", err)
	}
	log.Println("Initialized resource client")
	ccp := sdk.ChannelContext(c.ChannelID, fabsdk.WithUser(c.OrgUser))
	c.channelClient, err = channel.New(ccp)
	if err != nil {
		log.Panicf("failed to create channel client: %s", err)
	}
	log.Println("Initialized channel client")

	return c
}

package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"peers_crawler/model"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// SeedURL seed url of all network
var SeedURL = map[string][]string{
	"main": {
		"http://main1.seed.starcoin.org:9850",
		"http://main2.seed.starcoin.org:9850",
		"http://main3.seed.starcoin.org:9850",
		"http://main4.seed.starcoin.org:9850",
		"http://main5.seed.starcoin.org:9850",
		"http://main6.seed.starcoin.org:9850",
		"http://main7.seed.starcoin.org:9850",
		"http://main8.seed.starcoin.org:9850",
		"http://main9.seed.starcoin.org:9850",
	},
	"barnard": {
		"http://barnard4.seed.starcoin.org:9850",
		"http://barnard5.seed.starcoin.org:9850",
		"http://barnard6.seed.starcoin.org:9850",
	},
	"proxima": {
		"http://proxima1.seed.starcoin.org:9850",
		"http://proxima2.seed.starcoin.org:9850",
		"http://proxima3.seed.starcoin.org:9850",
	},
}

type resultUnit struct {
	PeerID        string `json:"peer_id"`
	VersionString string `json:"version_string"`
}

// NodePeersResult result from node peers
type NodePeersResult struct {
	JSONRPC string       `json:"jsonrpc"`
	Result  []resultUnit `json:"result"`
	ID      int          `json:"id"`
}

// GetMainPeers get peers from main seed
func GetMainPeers() error {
	net := "main"
	urls := SeedURL[net]
	return getPeers(urls, net)
}

// GetBarnardPeers get peers from barnard seed
func GetBarnardPeers() error {
	net := "barnard"
	urls := SeedURL[net]
	return getPeers(urls, net)
}

// GetProximaPeers get peers from proxima seed
func GetProximaPeers() error {
	net := "proxima"
	urls := SeedURL[net]
	return getPeers(urls, net)
}

func getPeers(urls []string, net string) error {

	modelPeers := make(map[string]model.Peer)

	for _, url := range urls {
		logs.Info(fmt.Sprintf("now get peers info from url: %s", url))
		peers, err := getPeersFromURL(url, net)
		if err != nil {
			logs.Error(fmt.Sprintf("request url %s, got error: %s", url, err.Error()))
			return err
		}
		for _, peer := range peers {
			modelPeers[peer.HashID] = peer
		}
	}

	updatePeers(modelPeers)

	return nil
}

// GetPeersFromURL get peers from specified url
func getPeersFromURL(url string, net string) ([]model.Peer, error) {

	payload := strings.NewReader(`{
		"id":200,
		"jsonrpc":"2.0",
		"method":"node.peers",
		"params":[]
	}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))
	var nodePeersRes NodePeersResult
	err = json.Unmarshal(body, &nodePeersRes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var modelPeers []model.Peer
	for _, peerInfo := range nodePeersRes.Result {

		r, _ := regexp.Compile("0x([A-Za-z0-9]+)")
		walletaddr := r.FindString(peerInfo.VersionString)

		modelPeers = append(modelPeers, model.Peer{HashID: peerInfo.PeerID, Network: net, Address: walletaddr})
	}

	return modelPeers, nil
}

func updatePeers(peers map[string]model.Peer) {
	o := orm.NewOrm()

	for _, p := range peers {
		r := o.Raw("insert into peer(hash_id,address,network) values(?,?,?)  on duplicate key update online_duration=online_duration+300,address=values(address)", p.HashID, p.Address, p.Network)
		_, err := r.Exec()
		if err != nil {
			logs.Error(fmt.Sprintf("updatePeers raw sql, got error: %s", err.Error()))
		}
	}
}

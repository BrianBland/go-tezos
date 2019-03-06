package goTezos

import (
	"testing"
)

func Test_GetSnapShot(t *testing.T) {

	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting snapshot 15 from the network")

	snapshot, err := gt.GetSnapShot(15)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(snapshot))
}

func Test_GetAllCurrentSnapShots(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all current snapshots")
	snapshots, err := gt.GetAllCurrentSnapShots()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(snapshots))
}

func Test_GetChainHead(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all current snapshots")
	head, err := gt.GetChainHead()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(head))
}

func Test_GetNetworkConstants(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting network constantss")
	netConts, err := gt.GetNetworkConstants()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(netConts))
}

func Test_GetNetworkVersions(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting network versions")
	netVers, err := gt.GetNetworkVersions()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(netVers))
}

func Test_GetBranchProtocol(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting branch protocol")
	brProto, err := gt.GetNetworkVersions()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(brProto))
}

func Test_GetBranchHash(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting branch hash")
	brHash, err := gt.GetBranchHash()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(brHash))
}

func Test_GetBlockLevelHead(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("mainnet-node.tzscan.io", "80")
	gt.AddNewClient(client)

	t.Log("Getting branch hash")
	levelHead, levelHash, err := gt.GetBlockLevelHead()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(levelHead))
	t.Log(PrettyReport(levelHash))
}

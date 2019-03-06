package goTezos

import (
	"testing"
)

func Test_NewRPCClient(t *testing.T) {

	t.Log("RPC client connecting to a Tezos node over localhost, port 8732")

	gtClient := NewTezosRPCClient("localhost", "8732")
	gt := NewGoTezos()
	gt.AddNewClient(gtClient)

	if !gtClient.Healthcheck() {
		t.Errorf("Unable to query RPC on 'localhost:8732'. Check that a node is accessible.")
	}
}

func Test_NewWebClient(t *testing.T) {

	t.Log("Web-based RPC client using https://rpc.tzbeta.net")

	gtClient := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gtClient.IsWebClient(true)

	gt := NewGoTezos()
	gt.AddNewClient(gtClient)

	if !gtClient.Healthcheck() {
		t.Errorf("Unable to query RPC at 'https://rpc.tzbeta.net'.")
	}
}

func Test_CreateWalletWithMnemonic(t *testing.T) {

	gt := NewGoTezos()

	t.Log("Create new wallet using Alphanet faucet account")

	mnemonic := "normal dash crumble neutral reflect parrot know stairs culture fault check whale flock dog scout"
	password := "PYh8nXDQLB"
	email := "vksbjweo.qsrgfvbw@tezos.example.org"

	// These values were gathered after manually importing above mnemonic into CLI wallet
	pkh := "tz1Qny7jVMGiwRrP9FikRK95jTNbJcffTpx1"
	pk := "edpkvEoAbkdaGALxi2FfeefB8hUkMZ4J1UVwkzyumx2GvbVpkYUHnm"
	sk := "edskRxB2DmoyZSyvhsqaJmw5CK6zYT7dbkUfEVSiQeWU1gw3ZMnC99QMMXru3imsbUrLhvuHktrymvNqhMxkhz7Y4LJAtevW5V"

	// Alphanet 'password' is email & password concatenated together
	myWallet, err := gt.CreateWallet(mnemonic, email+password)
	if err != nil {
		t.Errorf("Unable to create wallet from Mnemonic: %s", err)
	}

	if myWallet.Address != pkh || myWallet.Pk != pk || myWallet.Sk != sk {
		t.Errorf("Created wallet values do not match known answers")
	}
}

func Test_ImportWalletFullSk(t *testing.T) {

	gt := NewGoTezos()

	t.Log("Import existing wallet using complete secret key")

	pkh := "tz1fYvVTsSQWkt63P5V8nMjW764cSTrKoQKK"
	pk := "edpkvH3h91QHjKtuR45X9BJRWJJmK7s8rWxiEPnNXmHK67EJYZF75G"
	sk := "edskSA4oADtx6DTT6eXdBc6Pv5MoVBGXUzy8bBryi6D96RQNQYcRfVEXd2nuE2ZZPxs4YLZeM7KazUULFT1SfMDNyKFCUgk6vR"

	myWallet, err := gt.ImportWallet(pkh, pk, sk)
	if err != nil {
		t.Errorf("%s", err)
	}

	if myWallet.Address != pkh || myWallet.Pk != pk || myWallet.Sk != sk {
		t.Errorf("Created wallet values do not match known answers")
	}
}

func Test_ImportWalletSeedSk(t *testing.T) {

	gt := NewGoTezos()

	t.Log("Import existing wallet using seed-secret key")

	pkh := "tz1U8sXoQWGUMQrfZeAYwAzMZUvWwy7mfpPQ"
	pk := "edpkunwa7a3Y5vDr9eoKy4E21pzonuhqvNjscT9XG27aQV4gXq4dNm"
	sks := "edsk362Ypv3qLgbnGvZK7JwqNbwiLGe18XhTMFQY4gUonqnaCPiT6X"
	sk := "edskRjBSseEx9bSRSJJpbypJe5ZXucTtApb6qjechMB1BzEYwcEZyfLooo22Nwk33mPPJ3xZniFoa3o8Js7nNXDdqK9nNjFDi7"

	myWallet, err := gt.ImportWallet(pkh, pk, sks)
	if err != nil {
		t.Errorf("%s", err)
	}

	if myWallet.Address != pkh || myWallet.Pk != pk || myWallet.Sk != sk {
		t.Errorf("Created wallet values do not match known answers")
	}
}

func Test_ImportEncryptedSecret(t *testing.T) {

	gt := NewGoTezos()

	t.Log("Import wallet using password and encrypted key")

	pw := "password12345##"
	sk := "edesk1fddn27MaLcQVEdZpAYiyGQNm6UjtWiBfNP2ZenTy3CFsoSVJgeHM9pP9cvLJ2r5Xp2quQ5mYexW1LRKee2"

	// known answers for testing
	pk := "edpkuHMDkMz46HdRXYwom3xRwqk3zQ5ihWX4j8dwo2R2h8o4gPcbN5"
	pkh := "tz1L8fUQLuwRuywTZUP5JUw9LL3kJa8LMfoo"

	myWallet, err := gt.ImportEncryptedWallet(pw, sk)
	if err != nil {
		t.Errorf("%s", err)
	}

	if myWallet.Address != pkh || myWallet.Pk != pk {
		t.Errorf("Imported encrypted wallet does not match known answers")
	}
}

func TestGetChainHead(t *testing.T) {
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

func TestGetNetworkConstants(t *testing.T) {
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

func TestGetNetworkVersions(t *testing.T) {
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

func TestGetBranchProtocol(t *testing.T) {
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

func TestGetBranchHash(t *testing.T) {
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

func TestGetBlockLevelHead(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting branch hash")
	levelHead, levelHash, err := gt.GetBlockLevelHead()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(levelHead))
	t.Log(PrettyReport(levelHash))
}

func TestGetBlockAtLevel(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting block at level 100,000")
	block, err := gt.GetBlockAtLevel(100000)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(block))
}

func TestGetBlockByHash(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting block at BLz6yCE4BUL4ppo1zsEWdK9FRCt15WAY7ECQcuK9RtWg4xeEVL7")
	block, err := gt.GetBlockByHash("BLz6yCE4BUL4ppo1zsEWdK9FRCt15WAY7ECQcuK9RtWg4xeEVL7")
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(block))
}

func TestGetBlockOperationHashesHead(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all operation hashes at head block")
	opHashes, err := gt.GetBlockOperationHashesHead()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(opHashes))
}

func TestGetBlockOperationHashesAtLevel(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all operation hashes at level 100000")
	opHashes, err := gt.GetBlockOperationHashesAtLevel(100000)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(opHashes))
}

func TestGetBlockOperationHashes(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all operation hashes at hash BLz6yCE4BUL4ppo1zsEWdK9FRCt15WAY7ECQcuK9RtWg4xeEVL7")
	opHashes, err := gt.GetBlockOperationHashes("BLz6yCE4BUL4ppo1zsEWdK9FRCt15WAY7ECQcuK9RtWg4xeEVL7")
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(opHashes))
}

func TestGetAccountBalanceAtSnapshot(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting account balance for tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc, at cycle 15")
	balance, err := gt.GetAccountBalanceAtSnapshot("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 15)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(balance))
}

func TestGetAccountBalance(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting account balance for tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	balance, err := gt.GetAccountBalance("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(balance))
}

func TestGetDelegateStakingBalance(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting staking balance for tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc at cycle 15")
	balance, err := gt.GetDelegateStakingBalance("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 15)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(balance))
}

func TestGetCurrentCycle(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting the current cycle")
	cycle, err := gt.GetCurrentCycle()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(cycle))
}

func TestGetAccountBalanceAtBlock(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting staking balance for tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc at cycle 15")
	balance, err := gt.GetDelegateStakingBalance("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 15)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(balance))
}

func TestGetChainId(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting the chain ID for the network")
	chainId, err := gt.GetChainId()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(chainId))
}

func TestGetDelegationsForDelegate(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting delegations for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	delegations, err := gt.GetDelegationsForDelegate("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(delegations))
}

func TestGetDelegationsForDelegateByCycle(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting delegations for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycle 60")
	delegations, err := gt.GetDelegationsForDelegateByCycle("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(delegations))
}

func TestGetRewardsForDelegateForCycles(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting rewards for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycles 60-64")
	rewards, err := gt.GetRewardsForDelegateForCycles("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 60, 64)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(rewards))
}

func TestGetRewardsForDelegateCycle(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting rewards for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycle 60")
	rewards, err := gt.GetRewardsForDelegateCycle("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(rewards))
}

func TestGetCycleRewards(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting rewards for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycles 60")
	rewards, err := gt.getCycleRewards("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(rewards))
}

func TestGetDelegateRewardsForCycle(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting the rewards earned by delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for a cycle 60.")
	rewards, err := gt.GetDelegateRewardsForCycle("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(rewards))
}

func TestGetContractRewardsForDelegate(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	rewards, err := gt.GetDelegateRewardsForCycle("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log("Getting gross rewards and share for all delegations for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	contractRewards, err := gt.getContractRewardsForDelegate("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", rewards, 60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(contractRewards))
}

func TestGetShareOfContract(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting the the share for delegation KT1EidADxWfYeBgK8L1ZTbf7a9zyjKwCFjfH on tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycle 60")
	share, _, err := gt.GetShareOfContract("KT1EidADxWfYeBgK8L1ZTbf7a9zyjKwCFjfH", "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(share))
}

func TestGetDelegate(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting info for delegate on tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc head block")
	delegate, err := gt.GetDelegate("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(delegate))
}

func TestGetStakingBalanceAtCycle(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting the staking balance for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc at cycle 60")
	stakingBalance, err := gt.GetStakingBalanceAtCycle("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(stakingBalance))
}

func TestGetBakingRights(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all baking rights for cycle 60")
	bakingRights, err := gt.GetBakingRights(60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(bakingRights))
}

func TestGetBakingRightsForDelegate(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting baking rights with priotrity 2 for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycle 60")
	bakingRights, err := gt.GetBakingRightsForDelegate(60, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 2)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(bakingRights))
}

func TestGetBakingRightsForDelegateForCycles(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting baking rights with priotrity 2 for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycles 60-64")
	bakingRights, err := gt.GetBakingRightsForDelegateForCycles(60, 64, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc", 2)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(bakingRights))
}

func TestGetEndorsingRightsForDelegate(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all endorsing rights for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycles 60")
	endorsingRights, err := gt.GetEndorsingRightsForDelegate(60, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(endorsingRights))
}

func TestGetEndorsingRightsForDelegateForCycles(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all endorsing rights for delegate tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc for cycles 60-64")
	endorsingRights, err := gt.GetEndorsingRightsForDelegateForCycles(60, 64, "tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(endorsingRights))
}

func TestGetEndorsingRights(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all endorsing rights for cycles 60")
	endorsingRights, err := gt.GetEndorsingRights(60)
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(endorsingRights))
}

func TestGetAllDelegatesByHash(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all delegates at block hash BLz6yCE4BUL4ppo1zsEWdK9FRCt15WAY7ECQcuK9RtWg4xeEVL7")
	delegates, err := gt.GetAllDelegatesByHash("BLz6yCE4BUL4ppo1zsEWdK9FRCt15WAY7ECQcuK9RtWg4xeEVL7")
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(delegates))
}

func TestGetAllDelegates(t *testing.T) {
	gt := NewGoTezos()
	client := NewTezosRPCClient("rpc.tzbeta.net", "443")
	gt.AddNewClient(client)

	t.Log("Getting all delegates at head")
	delegates, err := gt.GetAllDelegates()
	if err != nil {
		t.Errorf("%s", err)
	}

	t.Log(PrettyReport(delegates))
}

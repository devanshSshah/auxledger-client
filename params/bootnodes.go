// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://22b4633caa46175da609c646403ee80b7a8f07656d65cf9457d17b05e416d7ee310dd94031431350cbd4f7fb53fbe3f95985ba9e57b7c03823e96cdd9fd9ba0f@172.16.10.98:30303",

	"enode://efaf812328b98f88ade38a2ebaaf4cf5e1ebace4fe5679333a16bf185b488831b4406a9455593e8a56515624211ecd879f18a20b8d8c9c80c1ee659f7cc6d10b@172.16.10.98:30403",
   
    "enode://071a4883b9c8c4beca13a1de6e5b9c96ee46ddd7c4fc13353e33b707bf1523261a5b1f21cde44627ecf5e08f00343a922ef3446290dc9ba87c99b75275262db2@172.16.10.98:30503",

	"enode://0705d61137c63dafd5e2c36799f336c7224273fa21dcfee3b84b5f70b89643c67fd40d6231ed991aef8ab22e658c0d1e39f870b8cad3340032f42549568be19a@172.16.10.98:30603",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}

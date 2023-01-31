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

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// TSF/DEV Go Bootnodes
	"enode://810c9a99dd53f0e8136432f31cfcf1f5925fe8dff4adfa53c462591516447ffc4d0b16e7b954e0b9d1a312c67876d8a16e563b29afe587596d28bd6ff0225e26@167.99.192.153:59997",
	"enode://10a4991c6b5b5eb831a07e7c35724d8266893c7986b526729348d1bf37b37a61c19c895013065e348b124b6431d7c9a9e8e970b3786eb95af2f921e7f93514f0@46.101.187.225:59997",
	"enode://f70c7ea4bd48dc7116a95b87b6e42fe61f0dc5f8a70e7157f68c99d13f393ec1bb3a0089701f6657a861e1e174e1aab650bc585246ac429b9cbe84b8077930db@165.232.99.209:59997",
	"enode://a126d9ce098720d00435745cfec981c6da1b2d04a65401639a36873ee1262ce9525ddc76b5668ec5aa8369e98f54d7814ad4b49b73a26fe516eb8aa1cf7a6b1e@167.71.54.120:59997",
	"enode://891aa70a41e8410aa9c6a10a3e18f53bb8d31283d3e40cde0f66dc4dab506e7e6f415e76651ce66fdb98af0b126b560f7e44d5e0a6e3a6cc920b3192bba231f4@68.183.214.221:59997",
	"enode://ee4a8ceeb6c68c4a62b4f97f2ff6b2f7d80a89d6527310d74d811f79299cafea8a71cc8370aa481de8d5f6ed2b0c01a897f07e42ab5e72b6d327b42927004c91@167.172.97.130:59997",
	"enode://666d2daffc08ddae814a799946637a521a830ff2e1b327119e315507f546c174d27013d94a690b1ffff846ea437d667c7e92dd85440e0b7719952056109f7c3b@167.172.174.6:59997",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://9fb686d265a7d6e3c5cd4de44e8b50eee40a618fed41ac6640174225a8bdaa1b7248f92369b9299e4576be8a532eae601353114b8443f274f7263a3e3d0b4c50@104.248.32.50:59997",
	"enode://d43e62049e8ca3028c6479ae408cc614f07f84d813947988c9a71b463a99e891746433deb9b03e1950b1bd091612d0cab624a45dcca07f4ecd7f05cfeb04faee@134.122.87.200:59997",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}

// These DNS names provide bootstrap connectivity for public testnets and the mainnet.
// See https://github.com/ethereum/discv4-dns-lists for more information.
var KnownDNSNetworks = map[common.Hash]string{}

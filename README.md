# Cosmos Hub
![banner](./docs/images/cosmos-hub-image.jpg)

[![CircleCI](https://circleci.com/gh/cosmos/gaia/tree/master.svg?style=shield)](https://circleci.com/gh/cosmos/gaia/tree/master)
[![codecov](https://codecov.io/gh/cosmos/gaia/branch/master/graph/badge.svg)](https://codecov.io/gh/cosmos/gaia)
[![Go Report Card](https://goreportcard.com/badge/github.com/cosmos/gaia)](https://goreportcard.com/report/github.com/cosmos/gaia)
[![license](https://img.shields.io/github/license/cosmos/gaia.svg)](https://github.com/cosmos/gaia/blob/master/LICENSE)
[![LoC](https://tokei.rs/b1/github/cosmos/gaia)](https://github.com/cosmos/gaia)
[![GolangCI](https://golangci.com/badges/github.com/cosmos/gaia.svg)](https://golangci.com/r/github.com/cosmos/gaia)
[![riot.im](https://img.shields.io/badge/riot.im-JOIN%20CHAT-green.svg)](https://riot.im/app/#/room/#cosmos-sdk:matrix.org)

This repository hosts `Gaia`, the first implementation of the Cosmos Hub based on the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk).

**Note**: Requires [Go 1.13+](https://golang.org/dl/)

## Cosmos Hub Mainnet

To run a full-node for the mainnet of the Cosmos Hub, first [install `gaiad`](./docs/gaia-tutorials/installation.md), then follow [the guide](./docs/gaia-tutorials/join-mainnet.md).

For status updates and genesis file, see the [launch repo](https://github.com/cosmos/launch).

## Quick Start

```
make install
```

## Cosmos Hub Archives

Archives of the Cosmos Hub can be found on this page [here](./docs/resources/archives.md).

## Disambiguation

This Cosmos-SDK project is not related to the [React-Cosmos](https://github.com/react-cosmos/react-cosmos) project (yet). Many thanks to Evan Coury and Ovidiu (@skidding) for this Github organization name. As per our agreement, this disambiguation notice will stay here.

# MEI Technical Specification

## Key
* CDPStoreKeyPrefix
* CDPStoreKey function (CDPStoreKeyPrefix + address)

## Keeper
* Bank keeper
* Mei keeper

## Mei Keeper
SetCDPbyAddress
```
Input: sdk.AccAddress, cdp
Output: nil, error
```

GetCDPbyAddress
```
Input: Address
Output: CDP, nil
```

SetMessage
```
Input: msg
Output: nil, error
```

GetMessageByID
```
Input: ID
Output: msg, nil
```

## Struct
CDP Structure
```
Properties:
Owner: sdk.AccAddress
CollateralAmount: sdk.Coins (Atom)
DebtAmount: sdk.Coins (Mei)
```

## Messages
* LockAtomMsg
* UnlockAtomMsg
* ReturnMeiMsg
* BorrowMeiMsg
* LiquidatedMsg

## Message Structure
LockAtomMsg
```
LockAmount: sdk.Coins
Owner: sdk.AccAddress
```

UnlockAtomMsg (Required: ATOM price)
```
UnlockAmount: sdk.Coins
Owner: sdk.AccAddress
```

ReturnMeiMsg
```
ReturnAmount: sdk.Coins
Owner: sdk.AccAddress
```

BorrowMeiMsg (Required: ATOM price)
```
BorrowAmount: sdk.Coins
Owner: sdk.AccAddress
```

LiquidatedMsg
```
CdpOwner: sdk.AccAddress
Liquidator: sdk.AccAddress
```

## Message Handler
LockAtomMsgHandler
1. Get CDP by address from keeper 
2. Add Atom amount with collateral amount on CDP
3. Update CDP into keeper
4. Move Atom from sender account to CDP module account

UnlockAtomMsgHandler
1. Get CDP by address from keeper
2. Calculate new collateral ratio
(If collateral ratio >= 150%: continue;
Else: return error;)
5. Subtract Atom Amount from collateral amount on CDP
6. Update CDP into keeper
7. Move Atom from CDP module account to sender account

ReturnMeiMsgHandler
1. Get CDP by address from keeper
2. Subtract MEI Amount from debt amount on CDP
3. Update CDP into keeper
4. Move Mei from CDP module account to sender account

BorrowMeiMsgHandler
1. Get CDP by address from keeper
2. Calculate new collateral ratio
(If collateral ratio >= 150%: continue;
Else: return error;)
5. Add MEI Amount to dept amount on CDP
6. Update CDP into keeper
7. Move Mei from CDP address to sender account

LiquidatedMsgHandler
1. Get CDP by address from keeper
2. Calculate new collateral ratio
(If collateral ratio >= 150%: return error;
Else: continue;)
5. Get debt amount from CDP
6. Burn MEI from liquidator equal to dept amount of this CDP
7. Set CDP collateral amount and debt amount to 0
8. Send atom collateral to liquidator account


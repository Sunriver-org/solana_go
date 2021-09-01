# Solana SDK library for Go

[![GoDoc](https://pkg.go.dev/badge/github.com/gagliardetto/solana-go?status.svg)](https://pkg.go.dev/github.com/gagliardetto/solana-go@v0.4.1?tab=doc)
[![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gagliardetto/solana-go?include_prereleases&label=release-tag)](https://github.com/gagliardetto/solana-go/releases)
[![Build Status](https://github.com/gagliardetto/solana-go/workflows/tests/badge.svg?branch=main)](https://github.com/gagliardetto/solana-go/actions?query=branch%3Amain)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gagliardetto/solana-go/main)](https://www.tickgit.com/browse?repo=github.com/gagliardetto/solana-go&branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/gagliardetto/solana-go)](https://goreportcard.com/report/github.com/gagliardetto/solana-go)

Go library to interface with Solana JSON RPC and WebSocket interfaces.

Clients for Solana native programs, Solana Program Library (SPL), and [Serum DEX](https://dex.projectserum.com) are in development.

More contracts to come.

<div align="center">
    <img src="https://user-images.githubusercontent.com/15271561/128235229-1d2d9116-23bb-464e-b2cc-8fb6355e3b55.png" margin="auto" height="175"/>
</div>

## Contents

- [Solana SDK library for Go](#solana-sdk-library-for-go)
  - [Contents](#contents)
  - [Features](#features)
  - [Current development status](#current-development-status)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [RPC usage examples](#rpc-usage-examples)
  - [Contributing](#contributing)
  - [License](#license)
  - [Credits](#credits)
  - [Installation](#installation)
  - [Examples](#examples)
    - [RPC Methods](#rpc-methods)
      - [GetAccountInfo](#index--rpc--getaccountinfo)
      - [GetBalance](#index--rpc--getbalance)
      - [GetBlock](#index--rpc--getblock)
      - [GetBlockCommitment](#index--rpc--getblockcommitment)
      - [GetBlockHeight](#index--rpc--getblockheight)
      - [GetBlockProduction](#index--rpc--getblockproduction)
      - [GetBlockTime](#index--rpc--getblocktime)
      - [GetBlocks](#index--rpc--getblocks)
      - [GetBlocksWithLimit](#index--rpc--getblockswithlimit)
      - [GetClusterNodes](#index--rpc--getclusternodes)
      - [GetConfirmedBlock](#index--rpc--getconfirmedblock)
      - [GetConfirmedBlocks](#index--rpc--getconfirmedblocks)
      - [GetConfirmedBlocksWithLimit](#index--rpc--getconfirmedblockswithlimit)
      - [GetConfirmedSignaturesForAddress2](#index--rpc--getconfirmedsignaturesforaddress2)
      - [GetConfirmedTransaction](#index--rpc--getconfirmedtransaction)
      - [GetEpochInfo](#index--rpc--getepochinfo)
      - [GetEpochSchedule](#index--rpc--getepochschedule)
      - [GetFeeCalculatorForBlockhash](#index--rpc--getfeecalculatorforblockhash)
      - [GetFeeRateGovernor](#index--rpc--getfeerategovernor)
      - [GetFees](#index--rpc--getfees)
      - [GetFirstAvailableBlock](#index--rpc--getfirstavailableblock)
      - [GetGenesisHash](#index--rpc--getgenesishash)
      - [GetHealth](#index--rpc--gethealth)
      - [GetIdentity](#index--rpc--getidentity)
      - [GetInflationGovernor](#index--rpc--getinflationgovernor)
      - [GetInflationRate](#index--rpc--getinflationrate)
      - [GetInflationReward](#index--rpc--getinflationreward)
      - [GetLargestAccounts](#index--rpc--getlargestaccounts)
      - [GetLeaderSchedule](#index--rpc--getleaderschedule)
      - [GetMaxRetransmitSlot](#index--rpc--getmaxretransmitslot)
      - [GetMaxShredInsertSlot](#index--rpc--getmaxshredinsertslot)
      - [GetMinimumBalanceForRentExemption](#index--rpc--getminimumbalanceforrentexemption)
      - [GetMultipleAccounts](#index--rpc--getmultipleaccounts)
      - [GetProgramAccounts](#index--rpc--getprogramaccounts)
      - [GetRecentBlockhash](#index--rpc--getrecentblockhash)
      - [GetRecentPerformanceSamples](#index--rpc--getrecentperformancesamples)
      - [GetSignatureStatuses](#index--rpc--getsignaturestatuses)
      - [GetSignaturesForAddress](#index--rpc--getsignaturesforaddress)
      - [GetSlot](#index--rpc--getslot)
      - [GetSlotLeader](#index--rpc--getslotleader)
      - [GetSlotLeaders](#index--rpc--getslotleaders)
      - [GetSnapshotSlot](#index--rpc--getsnapshotslot)
      - [GetStakeActivation](#index--rpc--getstakeactivation)
      - [GetSupply](#index--rpc--getsupply)
      - [GetTokenAccountBalance](#index--rpc--gettokenaccountbalance)
      - [GetTokenAccountsByDelegate](#index--rpc--gettokenaccountsbydelegate)
      - [GetTokenAccountsByOwner](#index--rpc--gettokenaccountsbyowner)
      - [GetTokenLargestAccounts](#index--rpc--gettokenlargestaccounts)
      - [GetTokenSupply](#index--rpc--gettokensupply)
      - [GetTransaction](#index--rpc--gettransaction)
      - [GetTransactionCount](#index--rpc--gettransactioncount)
      - [GetVersion](#index--rpc--getversion)
      - [GetVoteAccounts](#index--rpc--getvoteaccounts)
      - [MinimumLedgerSlot](#index--rpc--minimumledgerslot)
      - [RequestAirdrop](#index--rpc--requestairdrop)
      - [SendTransaction](#index--rpc--sendtransaction)
      - [SimulateTransaction](#index--rpc--simulatetransaction)
    - [Websocket Subscriptions](#websocket-subscriptions)
      - [AccountSubscribe](#index--ws-subscriptions--accountsubscribe)
      - [LogsSubscribe](#index--ws-subscriptions--logssubscribe)
      - [ProgramSubscribe](#index--ws-subscriptions--programsubscribe)
      - [RootSubscribe](#index--ws-subscriptions--rootsubscribe)
      - [SignatureSubscribe](#index--ws-subscriptions--signaturesubscribe)
      - [SlotSubscribe](#index--ws-subscriptions--slotsubscribe)
      - [VoteSubscribe](#index--ws-subscriptions--votesubscribe)


## Features

- [x] Full JSON RPC API
- [x] Full WebSocket JSON streaming API
- [ ] Wallet, account, and keys management
- [ ] Clients for native programs
- [ ] Clients for Solana Program Library
- [ ] Client for Serum
- [ ] More programs

## Current development status

There is currently **no stable release**. The SDK is actively developed and latest is `v0.4.1` which is an `alpha` release.

The RPC and WS client implementation is based on [this RPC spec](https://github.com/solana-labs/solana/blob/dff9c88193da142693cabebfcd3bf68fa8e8b873/docs/src/developing/clients/jsonrpc-api.md).

## Requirements

- Go 1.16 or later

## Installation

> :warning: `solana-go` works using SemVer but in 0 version, which means that the 'minor' will be changed when some broken changes are introduced into the application, and the 'patch' will be changed when a new feature with new changes is added or for bug fixing. As soon as v1.0.0 be released, `solana-go` will start to use SemVer as usual.

```bash
$ cd my-project
$ go get github.com/gagliardetto/solana-go@v0.4.1
```

## Examples

### Create account

```go
package main

import (
  "context"
  "fmt"

  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  // Create a new account:
  account := solana.NewAccount()
  fmt.Println("account private key:", account.PrivateKey)
  fmt.Println("account public key:", account.PublicKey())

  // Create a new RPC client:
  client := rpc.New(rpc.TestNet_RPC)

  // Airdrop 1 sol to the new account:
  out, err := client.RequestAirdrop(
    context.TODO(),
    account.PublicKey(),
    solana.LAMPORTS_PER_SOL,
    "",
  )
  if err != nil {
    panic(err)
  }
  fmt.Println("airdrop transaction signature:", out)
}
```

### Transfer Sol from one wallet to another wallet

```go
package main

import (
  "context"
  "fmt"
  "time"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/programs/system"
  "github.com/gagliardetto/solana-go/rpc"
  "github.com/gagliardetto/solana-go/rpc/jsonrpc"
)

func main() {
  // Create a new RPC client:
  client := rpc.New(rpc.DevNet_RPC)

  // Load the account that you will send funds FROM:
  accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile("/path/to/.config/solana/id.json")
  if err != nil {
    panic(err)
  }
  fmt.Println("accountFrom private key:", accountFrom)
  fmt.Println("accountFrom public key:", accountFrom.PublicKey())

  if true {
    // Airdrop 10 sol to the account so it will have something to transfer:
    out, err := client.RequestAirdrop(
      context.TODO(),
      accountFrom.PublicKey(),
      solana.LAMPORTS_PER_SOL*10,
      rpc.CommitmentFinalized,
    )
    if err != nil {
      panic(err)
    }
    fmt.Println("airdrop transaction signature:", out)
    time.Sleep(time.Second * 5)
  }
  //---------------

  // The public key of the account that you will send sol TO:
  accountTo := solana.MustPublicKeyFromBase58("TODO")

  recent, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
  if err != nil {
    panic(err)
  }

  // The amount to send (in lamports);
  // 1 sol = 1000000000 lamports
  amount := uint64(3333)
  tx, err := solana.NewTransaction(
    []solana.Instruction{
      system.NewTransferInstruction(
        amount,
        accountFrom.PublicKey(),
        accountTo,
      ).Build(),
    },
    recent.Value.Blockhash,
    solana.TransactionPayer(accountFrom.PublicKey()),
  )
  if err != nil {
    panic(err)
  }

  _, err = tx.Sign(
    func(key solana.PublicKey) *solana.PrivateKey {
      if accountFrom.PublicKey().Equals(key) {
        return &accountFrom
      }
      return nil
    },
  )
  if err != nil {
    panic(fmt.Errorf("unable to sign transaction: %w", err))
  }
  spew.Dump(tx)

  sig, err := client.SendTransactionWithOpts(
    context.TODO(),
    tx,
    false,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    spew.Config.DisablePointerMethods = true
    spew.Dump(err)
    spew.Dump(err.(*jsonrpc.RPCError).Code)
    spew.Dump(err.(*jsonrpc.RPCError).Data)
    spew.Dump(err.(*jsonrpc.RPCError).Message)
    panic(err)
  }
  spew.Dump(sig)
}

```

## RPC usage examples

### RPC Methods

#### [index](#contents) > [RPC](#rpc-methods) > GetAccountInfo

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  bin "github.com/dfuse-io/binary"
  solana "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/programs/token"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.MainNetBeta_RPC
  client := rpc.New(endpoint)

  {
    pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
    // basic usage
    resp, err := client.GetAccountInfo(
      context.TODO(),
      pubKey,
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(resp)

    var mint token.Mint
    // Account{}.Data.GetBinary() returns the *decoded* binary data
    // regardless the original encoding (it can handle them all). 
    err = bin.NewDecoder(resp.Value.Data.GetBinary()).Decode(&mint)
    if err != nil {
      panic(err)
    }
    spew.Dump(mint)
    // NOTE: The supply is mint.Supply, with the mint.Decimals:
    // mint.Supply = 9998022451607088
    // mint.Decimals = 6
    // ... which means that the supply is 9998022451.607088
  }
  {
    // Or you can use `GetAccountDataIn` which does all of the above in one call:
    pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
    var mint token.Mint
    // Get the account, and decode its data into the provided mint object:
    err := client.GetAccountDataIn(
      context.TODO(),
      pubKey,
      &mint,
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(mint)
  }
  {
    pubKey := solana.MustPublicKeyFromBase58("4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R") // raydium token
    // advanced usage
    resp, err := client.GetAccountInfoWithOpts(
      context.TODO(),
      pubKey,
      // You can specify more options here:
      &rpc.GetAccountInfoOpts{
        Encoding:   solana.EncodingBase64Zstd,
        Commitment: rpc.CommitmentFinalized,
        // You can get just a part of the account data by specify a DataSlice:
        // DataSlice: &rpc.DataSlice{
        //  Offset: pointer.ToUint64(0),
        //  Length: pointer.ToUint64(1024),
        // },
      },
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(resp)

    var mint token.Mint
    err = bin.NewDecoder(resp.Value.Data.GetBinary()).Decode(&mint)
    if err != nil {
      panic(err)
    }
    spew.Dump(mint)
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetBalance

```go
package main

import (
  "context"
  "fmt"
  "math/big"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.MainNetBeta_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932")
  out, err := client.GetBalance(
    context.TODO(),
    pubKey,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
  spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports

  var lamportsOnAccount = new(big.Float).SetUint64(uint64(out.Value))
  // Convert lamports to sol:
  var solBalance = new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))

  // WARNING: this is not a precise conversion.
  fmt.Println("◎", solBalance.Text('f', 10))
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetBlock

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
  if err != nil {
    panic(err)
  }

  {
    out, err := client.GetBlock(context.TODO(), uint64(example.Context.Slot))
    if err != nil {
      panic(err)
    }
    // spew.Dump(out) // NOTE: This generates a lot of output.
    spew.Dump(len(out.Transactions))
  }

  {
    includeRewards := false
    out, err := client.GetBlockWithOpts(
      context.TODO(),
      uint64(example.Context.Slot),
      // You can specify more options here:
      &rpc.GetBlockOpts{
        Encoding:   solana.EncodingBase64,
        Commitment: rpc.CommitmentFinalized,
        // Get only signatures:
        TransactionDetails: rpc.TransactionDetailsSignatures,
        // Exclude rewards:
        Rewards: &includeRewards,
      },
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetBlockCommitment

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
  if err != nil {
    panic(err)
  }

  out, err := client.GetBlockCommitment(
    context.TODO(),
    uint64(example.Context.Slot),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetBlockHeight

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetBlockHeight(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetBlockProduction

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  {
    out, err := client.GetBlockProduction(context.TODO())
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
  {
    out, err := client.GetBlockProductionWithOpts(
      context.TODO(),
      &rpc.GetBlockProductionOpts{
        Commitment: rpc.CommitmentFinalized,
        // Range: &rpc.SlotRangeRequest{
        //  FirstSlot: XXXXXX,
        //  Identity:  solana.MustPublicKeyFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
        // },
      },
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetBlockTime

```go
package main

import (
  "context"
  "time"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }

  out, err := client.GetBlockTime(
    context.TODO(),
    uint64(example.Context.Slot),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
  spew.Dump(out.Time().Format(time.RFC1123))
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetBlocks

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }

  endSlot := uint64(example.Context.Slot)
  out, err := client.GetBlocks(
    context.TODO(),
    uint64(example.Context.Slot-3),
    &endSlot,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetBlocksWithLimit

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }

  limit := uint64(4)
  out, err := client.GetBlocksWithLimit(
    context.TODO(),
    uint64(example.Context.Slot-10),
    limit,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetClusterNodes

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetClusterNodes(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetConfirmedBlock

```go
package main

import (
  "context"

  "github.com/AlekSi/pointer"
  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }

  { // deprecated and is going to be removed in solana-core v1.8
    out, err := client.GetConfirmedBlock(
      context.TODO(),
      uint64(example.Context.Slot),
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
  {
    slot := uint64(example.Context.Slot)
    out, err := client.GetConfirmedBlockWithOpts(
      context.TODO(),
      slot,
      // You can specify more options here:
      &rpc.GetConfirmedBlockOpts{
        Encoding:   solana.EncodingBase64,
        Commitment: rpc.CommitmentFinalized,
        // Get only signatures:
        TransactionDetails: rpc.TransactionDetailsSignatures,
        // Exclude rewards:
        Rewards: pointer.ToBool(false),
      },
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetConfirmedBlocks

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }

  {
    endSlot := uint64(example.Context.Slot)
    // deprecated and is going to be removed in solana-core v1.8
    out, err := client.GetConfirmedBlocks(
      context.TODO(),
      uint64(example.Context.Slot-3),
      &endSlot,
      rpc.CommitmentFinalized,
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetConfirmedBlocksWithLimit

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }

  limit := uint64(3)
  { // deprecated and is going to be removed in solana-core v1.8
    out, err := client.GetConfirmedBlocksWithLimit(
      context.TODO(),
      uint64(example.Context.Slot-10),
      limit,
      rpc.CommitmentFinalized,
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetConfirmedSignaturesForAddress2

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
  {
    // deprecated and is going to be removed in solana-core v1.8
    out, err := client.GetConfirmedSignaturesForAddress2(
      context.TODO(),
      pubKey,
      // TODO:
      nil,
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetConfirmedTransaction

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
  // Let's get a valid transaction to use in the example:
  example, err := client.GetConfirmedSignaturesForAddress2(
    context.TODO(),
    pubKey,
    nil,
  )
  if err != nil {
    panic(err)
  }

  out, err := client.GetConfirmedTransaction(
    context.TODO(),
    example[0].Signature,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetEpochInfo

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetEpochInfo(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetEpochSchedule

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetEpochSchedule(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetFeeCalculatorForBlockhash

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  example, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }

  out, err := client.GetFeeCalculatorForBlockhash(
    context.TODO(),
    example.Value.Blockhash,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetFeeRateGovernor

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetFeeRateGovernor(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetFees

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetFees(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetFirstAvailableBlock

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetFirstAvailableBlock(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetGenesisHash

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetGenesisHash(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetHealth

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetHealth(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
  spew.Dump(out == rpc.HealthOk)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetIdentity

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetIdentity(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetInflationGovernor

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetInflationGovernor(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetInflationRate

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetInflationRate(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetInflationReward

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("6dmNQ5jwLeLk5REvio1JcMshcbvkYMwy26sJ8pbkvStu")

  out, err := client.GetInflationReward(
    context.TODO(),
    []solana.PublicKey{
      pubKey,
    },
    &rpc.GetInflationRewardOpts{
      Commitment: rpc.CommitmentFinalized,
    },
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetLargestAccounts

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetLargestAccounts(
    context.TODO(),
    rpc.CommitmentFinalized,
    rpc.LargestAccountsFilterCirculating,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetLeaderSchedule

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetLeaderSchedule(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out) // NOTE: this creates a lot of output
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetMaxRetransmitSlot

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetMaxRetransmitSlot(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetMaxShredInsertSlot

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetMaxShredInsertSlot(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetMinimumBalanceForRentExemption

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  dataSize := uint64(1024 * 9)
  out, err := client.GetMinimumBalanceForRentExemption(
    context.TODO(),
    dataSize,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetMultipleAccounts

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.MainNetBeta_RPC
  client := rpc.New(endpoint)

  {
    out, err := client.GetMultipleAccounts(
      context.TODO(),
      solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"),  // serum token
      solana.MustPublicKeyFromBase58("4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R"), // raydium token
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
  {
    out, err := client.GetMultipleAccountsWithOpts(
      context.TODO(),
      []solana.PublicKey{solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"), // serum token
        solana.MustPublicKeyFromBase58("4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R"), // raydium token
      },
      &rpc.GetMultipleAccountsOpts{
        Encoding:   solana.EncodingBase64Zstd,
        Commitment: rpc.CommitmentFinalized,
        // You can get just a part of the account data by specify a DataSlice:
        // DataSlice: &rpc.DataSlice{
        //  Offset: pointer.ToUint64(0),
        //  Length: pointer.ToUint64(1024),
        // },
      },
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetProgramAccounts

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetProgramAccounts(
    context.TODO(),
    solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(len(out))
  spew.Dump(out) // NOTE: this can generate a lot of output
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetRecentBlockhash

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  recent, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(recent)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetRecentPerformanceSamples

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  limit := uint(3)
  out, err := client.GetRecentPerformanceSamples(
    context.TODO(),
    &limit,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetSignatureStatuses

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetSignatureStatuses(
    context.TODO(),
    true,
    // All the transactions you want the get the status for:
    solana.MustSignatureFromBase58("2CwH8SqVZWFa1EvsH7vJXGFors1NdCuWJ7Z85F8YqjCLQ2RuSHQyeGKkfo1Tj9HitSTeLoMWnxpjxF2WsCH8nGWh"),
    solana.MustSignatureFromBase58("5YJHZPeHZuZjhunBc1CCB1NDRNf2tTJNpdb3azGsR7PfyEncCDhr95wG8EWrvjNXBc4wCKixkheSbCxoC2NCG3X7"),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetSignaturesForAddress

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetSignaturesForAddress(
    context.TODO(),
    solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetSlot

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetSlot(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetSlotLeader

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetSlotLeader(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetSlotLeaders

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  recent, err := client.GetRecentBlockhash(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }

  out, err := client.GetSlotLeaders(
    context.TODO(),
    uint64(recent.Context.Slot),
    10,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetSnapshotSlot

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetSnapshotSlot(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetStakeActivation

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("EW2p7QCJNHMVj5nQCcW7Q2BDETtNBXn68FyucU4RCjvb")
  out, err := client.GetStakeActivation(
    context.TODO(),
    pubKey,
    rpc.CommitmentFinalized,
    nil,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetSupply

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetSupply(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetTokenAccountBalance

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("EzK5qLWhftu8Z2znVa5fozVtobbjhd8Gdu9hQHpC8bec")
  out, err := client.GetTokenAccountBalance(
    context.TODO(),
    pubKey,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetTokenAccountsByDelegate

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("AfkALUPjQp8R1rUwE6KhT38NuTYWCncwwHwcJu7UtAfV")
  out, err := client.GetTokenAccountsByDelegate(
    context.TODO(),
    pubKey,
    &rpc.GetTokenAccountsConfig{
      Mint: solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112"),
    },
    nil,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetTokenAccountsByOwner

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("7HZaCWazgTuuFuajxaaxGYbGnyVKwxvsJKue1W4Nvyro")
  out, err := client.GetTokenAccountsByOwner(
    context.TODO(),
    pubKey,
    &rpc.GetTokenAccountsConfig{
      Mint: solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112"),
    },
    nil,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetTokenLargestAccounts

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.MainNetBeta_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
  out, err := client.GetTokenLargestAccounts(
    context.TODO(),
    pubKey,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetTokenSupply

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.MainNetBeta_RPC
  client := rpc.New(endpoint)

  pubKey := solana.MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt") // serum token
  out, err := client.GetTokenSupply(
    context.TODO(),
    pubKey,
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetTransaction

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  txSig := solana.MustSignatureFromBase58("4bjVLV1g9SAfv7BSAdNnuSPRbSscADHFe4HegL6YVcuEBMY83edLEvtfjE4jfr6rwdLwKBQbaFiGgoLGtVicDzHq")
  {
    out, err := client.GetTransaction(
      context.TODO(),
      txSig,
      nil,
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
    spew.Dump(out.Transaction.GetParsedTransaction())
  }
  {
    out, err := client.GetTransaction(
      context.TODO(),
      txSig,
      &rpc.GetTransactionOpts{
        Encoding: solana.EncodingJSON,
      },
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
    spew.Dump(out.Transaction.GetParsedTransaction())
  }
  {
    out, err := client.GetTransaction(
      context.TODO(),
      txSig,
      &rpc.GetTransactionOpts{
        Encoding: solana.EncodingBase58,
      },
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
    spew.Dump(out.Transaction.GetBinary())
  }
  {
    out, err := client.GetTransaction(
      context.TODO(),
      txSig,
      &rpc.GetTransactionOpts{
        Encoding: solana.EncodingBase64,
      },
    )
    if err != nil {
      panic(err)
    }
    spew.Dump(out)
    spew.Dump(out.Transaction.GetBinary())
  }
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetTransactionCount

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetTransactionCount(
    context.TODO(),
    rpc.CommitmentFinalized,
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetVersion

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetVersion(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > GetVoteAccounts

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.GetVoteAccounts(
    context.TODO(),
    &rpc.GetVoteAccountsOpts{
      VotePubkey: solana.MustPublicKeyFromBase58("vot33MHDqT6nSwubGzqtc6m16ChcUywxV7tNULF19Vu"),
    },
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > MinimumLedgerSlot

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  out, err := client.MinimumLedgerSlot(
    context.TODO(),
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > RequestAirdrop

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
)

func main() {
  endpoint := rpc.TestNet_RPC
  client := rpc.New(endpoint)

  amount := solana.LAMPORTS_PER_SOL // 1 sol
  pubKey := solana.MustPublicKeyFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
  out, err := client.RequestAirdrop(
    context.TODO(),
    pubKey,
    amount,
    "",
  )
  if err != nil {
    panic(err)
  }
  spew.Dump(out)
}
```

#### [index](#contents) > [RPC](#rpc-methods) > SendTransaction

```go
package main

func main() {

}
```

#### [index](#contents) > [RPC](#rpc-methods) > SimulateTransaction

```go
package main

func main() {

}
```


### Websocket Subscriptions

#### [index](#contents) > [WS Subscriptions](#websocket-subscriptions) > AccountSubscribe

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
  "github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
  client, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
  if err != nil {
    panic(err)
  }
  program := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin") // serum

  {
    sub, err := client.AccountSubscribe(
      program,
      "",
    )
    if err != nil {
      panic(err)
    }
    defer sub.Unsubscribe()

    for {
      got, err := sub.Recv()
      if err != nil {
        panic(err)
      }
      spew.Dump(got)
    }
  }
  if false {
    sub, err := client.AccountSubscribeWithOpts(
      program,
      "",
      // You can specify the data encoding of the returned accounts:
      solana.EncodingBase64,
    )
    if err != nil {
      panic(err)
    }
    defer sub.Unsubscribe()

    for {
      got, err := sub.Recv()
      if err != nil {
        panic(err)
      }
      spew.Dump(got)
    }
  }
}
```

#### [index](#contents) > [WS Subscriptions](#websocket-subscriptions) > LogsSubscribe

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
  "github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
  client, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
  if err != nil {
    panic(err)
  }
  program := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin") // serum

  {
    // Subscribe to log events that mention the provided pubkey:
    sub, err := client.LogsSubscribeMentions(
      program,
      rpc.CommitmentRecent,
    )
    if err != nil {
      panic(err)
    }
    defer sub.Unsubscribe()

    for {
      got, err := sub.Recv()
      if err != nil {
        panic(err)
      }
      spew.Dump(got)
    }
  }
  if false {
    // Subscribe to all log events:
    sub, err := client.LogsSubscribe(
      ws.LogsSubscribeFilterAll,
      rpc.CommitmentRecent,
    )
    if err != nil {
      panic(err)
    }
    defer sub.Unsubscribe()

    for {
      got, err := sub.Recv()
      if err != nil {
        panic(err)
      }
      spew.Dump(got)
    }
  }
}
```

#### [index](#contents) > [WS Subscriptions](#websocket-subscriptions) > ProgramSubscribe

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
  "github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
  client, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
  if err != nil {
    panic(err)
  }
  program := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA") // token

  sub, err := client.ProgramSubscribeWithOpts(
    program,
    rpc.CommitmentRecent,
    solana.EncodingBase64Zstd,
    nil,
  )
  if err != nil {
    panic(err)
  }
  defer sub.Unsubscribe()

  for {
    got, err := sub.Recv()
    if err != nil {
      panic(err)
    }
    spew.Dump(got)

    decodedBinary := got.Value.Account.Data.GetBinary()
    if decodedBinary != nil {
      // spew.Dump(decodedBinary)
    }

    // or if you requested solana.EncodingJSONParsed and it is supported:
    rawJSON := got.Value.Account.Data.GetRawJSON()
    if rawJSON != nil {
      // spew.Dump(rawJSON)
    }
  }
}
```

#### [index](#contents) > [WS Subscriptions](#websocket-subscriptions) > RootSubscribe

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
  "github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
  client, err := ws.Connect(context.Background(), rpc.TestNet_WS)
  if err != nil {
    panic(err)
  }

  sub, err := client.RootSubscribe()
  if err != nil {
    panic(err)
  }

  for {
    got, err := sub.Recv()
    if err != nil {
      panic(err)
    }
    spew.Dump(got)
  }
}
```

#### [index](#contents) > [WS Subscriptions](#websocket-subscriptions) > SignatureSubscribe

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go"
  "github.com/gagliardetto/solana-go/rpc"
  "github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
  client, err := ws.Connect(context.Background(), rpc.TestNet_WS)
  if err != nil {
    panic(err)
  }

  txSig := solana.MustSignatureFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

  sub, err := client.SignatureSubscribe(
    txSig,
    "",
  )
  if err != nil {
    panic(err)
  }
  defer sub.Unsubscribe()

  for {
    got, err := sub.Recv()
    if err != nil {
      panic(err)
    }
    spew.Dump(got)
  }
}
```

#### [index](#contents) > [WS Subscriptions](#websocket-subscriptions) > SlotSubscribe

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
  "github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
  client, err := ws.Connect(context.Background(), rpc.TestNet_WS)
  if err != nil {
    panic(err)
  }

  sub, err := client.SlotSubscribe()
  if err != nil {
    panic(err)
  }
  defer sub.Unsubscribe()

  for {
    got, err := sub.Recv()
    if err != nil {
      panic(err)
    }
    spew.Dump(got)
  }
}
```

#### [index](#contents) > [WS Subscriptions](#websocket-subscriptions) > VoteSubscribe

```go
package main

import (
  "context"

  "github.com/davecgh/go-spew/spew"
  "github.com/gagliardetto/solana-go/rpc"
  "github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
  client, err := ws.Connect(context.Background(), rpc.MainNetBeta_WS)
  if err != nil {
    panic(err)
  }

  // NOTE: this subscription must be enabled by the node you're connecting to.
  // This subscription is disabled by default.
  sub, err := client.VoteSubscribe()
  if err != nil {
    panic(err)
  }
  defer sub.Unsubscribe()

  for {
    got, err := sub.Recv()
    if err != nil {
      panic(err)
    }
    spew.Dump(got)
  }
}
```

## Contributing

We encourage everyone to contribute, submit issues, PRs, discuss. Every kind of help is welcome.

## License

[Apache 2.0](LICENSE)

## Credits

- Gopher logo was originally created by Takuya Ueda (https://twitter.com/tenntenn). Licensed under the Creative Commons 3.0 Attributions license.

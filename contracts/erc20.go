// Copyright 2022 Atrix Foundation
// This file is part of the Atrix Network packages.
//
// Atrix is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Atrix packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Atrix packages. If not, see https://github.com/Atrix/Atrix/blob/main/LICENSE

package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/Atrix/ethermint/x/evm/types"

	"github.com/Atrix/Atrix/v11/x/erc20/types"
)

var (
	//go:embed compiled_contracts/ERC20MinterBurnerDecimals.json
	ERC20MinterBurnerDecimalsJSON []byte // nolint: golint

	// ERC20MinterBurnerDecimalsContract is the compiled erc20 contract
	ERC20MinterBurnerDecimalsContract evmtypes.CompiledContract

	// ERC20MinterBurnerDecimalsAddress is the erc20 module address
	ERC20MinterBurnerDecimalsAddress common.Address
)

func init() {
	ERC20MinterBurnerDecimalsAddress = types.ModuleAddress

	err := json.Unmarshal(ERC20MinterBurnerDecimalsJSON, &ERC20MinterBurnerDecimalsContract)
	if err != nil {
		panic(err)
	}

	if len(ERC20MinterBurnerDecimalsContract.Bin) == 0 {
		panic("load contract failed")
	}
}

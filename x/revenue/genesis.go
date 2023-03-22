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

package revenue

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Atrix/Atrix/v11/x/revenue/keeper"
	"github.com/Atrix/Atrix/v11/x/revenue/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) {
	err := k.SetParams(ctx, data.Params)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed setting params"))
	}

	for _, revenue := range data.Revenues {
		contract := revenue.GetContractAddr()
		deployer := revenue.GetDeployerAddr()
		withdrawer := revenue.GetWithdrawerAddr()

		// Set initial contracts receiving transaction fees
		k.SetRevenue(ctx, revenue)
		k.SetDeployerMap(ctx, deployer, contract)

		if len(withdrawer) != 0 {
			k.SetWithdrawerMap(ctx, withdrawer, contract)
		}
	}
}

// ExportGenesis export module state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:   k.GetParams(ctx),
		Revenues: k.GetRevenues(ctx),
	}
}

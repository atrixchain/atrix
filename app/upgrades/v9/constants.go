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

package v9

const (
	// UpgradeName is the shared upgrade plan name for mainnet
	UpgradeName = "v9.0.0"
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/Atrix/Atrix/releases/download/v9.0.0/Atrix_9.0.0_Darwin_arm64.tar.gz","darwin/amd64":"https://github.com/Atrix/Atrix/releases/download/v9.0.0/Atrix_9.0.0_Darwin_amd64.tar.gz","linux/arm64":"https://github.com/Atrix/Atrix/releases/download/v9.0.0/Atrix_9.0.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/Atrix/Atrix/releases/download/v9.0.0/Atrix_9.0.0_Linux_amd64.tar.gz","windows/x86_64":"https://github.com/Atrix/Atrix/releases/download/v9.0.0/Atrix_9.0.0_Windows_x86_64.zip"}}'`
	// MaxRecover is the maximum amount of coins to be redistributed in the upgrade
	MaxRecover = "93590289356801768542679"
)

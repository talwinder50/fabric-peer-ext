/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetRoles(t *testing.T) {
	oldVal := viper.Get(confRoles)
	defer viper.Set(confRoles, oldVal)

	roles := "endorser,committer"
	viper.Set(confRoles, roles)
	assert.Equal(t, roles, GetRoles())
}

func TestGetPvtDataCacheSize(t *testing.T) {
	oldVal := viper.Get(confPvtDataCacheSize)
	defer viper.Set(confPvtDataCacheSize, oldVal)

	val := GetPvtDataCacheSize()
	assert.Equal(t, val, 10)

	viper.Set(confPvtDataCacheSize, 99)
	val = GetPvtDataCacheSize()
	assert.Equal(t, val, 99)

}

func TestGetTransientDataLevelDBPath(t *testing.T) {
	oldVal := viper.Get("peer.fileSystemPath")
	defer viper.Set("peer.fileSystemPath", oldVal)

	viper.Set("peer.fileSystemPath", "/tmp123")

	assert.Equal(t, "/tmp123/ledgersData/transientDataLeveldb", GetTransientDataLevelDBPath())
}

func TestGetTransientDataExpiredIntervalTime(t *testing.T) {
	oldVal := viper.Get(confTransientDataCleanupIntervalTime)
	defer viper.Set(confTransientDataCleanupIntervalTime, oldVal)

	viper.Set(confTransientDataCleanupIntervalTime, "")
	assert.Equal(t, defaultTransientDataCleanupIntervalTime, GetTransientDataExpiredIntervalTime())

	viper.Set(confTransientDataCleanupIntervalTime, 111*time.Second)
	assert.Equal(t, 111*time.Second, GetTransientDataExpiredIntervalTime())
}

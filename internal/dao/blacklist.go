// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-starter/internal/dao/internal"
)

// internalBlacklistDao is internal type for wrapping internal DAO implements.
type internalBlacklistDao = *internal.BlacklistDao

// blacklistDao is the data access object for table blacklist.
// You can define custom methods on it to extend its functionality as you wish.
type blacklistDao struct {
	internalBlacklistDao
}

var (
	// Blacklist is globally public accessible object for table blacklist operations.
	Blacklist = blacklistDao{
		internal.NewBlacklistDao(),
	}
)

// Fill with you ideas below.
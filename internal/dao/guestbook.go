// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"coderblog-interface/internal/dao/internal"
)

// internalGuestbookDao is internal type for wrapping internal DAO implements.
type internalGuestbookDao = *internal.GuestbookDao

// guestbookDao is the data access object for table guestbook.
// You can define custom methods on it to extend its functionality as you wish.
type guestbookDao struct {
	internalGuestbookDao
}

var (
	// Guestbook is globally public accessible object for table guestbook operations.
	Guestbook = guestbookDao{
		internal.NewGuestbookDao(),
	}
)

// Fill with you ideas below.
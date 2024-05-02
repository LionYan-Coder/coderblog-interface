// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MailboxDao is the data access object for table mailbox.
type MailboxDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns MailboxColumns // columns contains all the column names of Table for convenient usage.
}

// MailboxColumns defines and stores column names for table mailbox.
type MailboxColumns struct {
	Id       string //
	UserId   string // 用户id
	Content  string // 留言内容
	CreateAt string // 创建时间
	UpdateAt string // 更新时间
}

// mailboxColumns holds the columns for table mailbox.
var mailboxColumns = MailboxColumns{
	Id:       "id",
	UserId:   "user_id",
	Content:  "content",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewMailboxDao creates and returns a new DAO object for table data access.
func NewMailboxDao() *MailboxDao {
	return &MailboxDao{
		group:   "default",
		table:   "mailbox",
		columns: mailboxColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MailboxDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MailboxDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MailboxDao) Columns() MailboxColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MailboxDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MailboxDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MailboxDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

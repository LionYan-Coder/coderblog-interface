// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotebookCommentDao is the data access object for table notebook_comment.
type NotebookCommentDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns NotebookCommentColumns // columns contains all the column names of Table for convenient usage.
}

// NotebookCommentColumns defines and stores column names for table notebook_comment.
type NotebookCommentColumns struct {
	Id         string // 笔记id
	UserId     string // 用户id
	NotebookId string // 笔记id
	ReplyId    string // 回复id
	Content    string // 内容
	CreateAt   string // 创建时间
	UpdateAt   string // 更新时间
}

// notebookCommentColumns holds the columns for table notebook_comment.
var notebookCommentColumns = NotebookCommentColumns{
	Id:         "id",
	UserId:     "user_id",
	NotebookId: "notebook_id",
	ReplyId:    "reply_id",
	Content:    "content",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
}

// NewNotebookCommentDao creates and returns a new DAO object for table data access.
func NewNotebookCommentDao() *NotebookCommentDao {
	return &NotebookCommentDao{
		group:   "default",
		table:   "notebook_comment",
		columns: notebookCommentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NotebookCommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NotebookCommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NotebookCommentDao) Columns() NotebookCommentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NotebookCommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NotebookCommentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NotebookCommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

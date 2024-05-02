// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotebookDao is the data access object for table notebook.
type NotebookDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns NotebookColumns // columns contains all the column names of Table for convenient usage.
}

// NotebookColumns defines and stores column names for table notebook.
type NotebookColumns struct {
	Id        string //
	UserId    string // 用户id
	Title     string // 笔记标题
	Summary   string // 笔记概要
	Tags      string // 笔记标签
	Content   string // 笔记内容
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	Published string // 发布状态（1: 已发布，0: 未发布）
}

// notebookColumns holds the columns for table notebook.
var notebookColumns = NotebookColumns{
	Id:        "id",
	UserId:    "user_id",
	Title:     "title",
	Summary:   "summary",
	Tags:      "tags",
	Content:   "content",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	Published: "published",
}

// NewNotebookDao creates and returns a new DAO object for table data access.
func NewNotebookDao() *NotebookDao {
	return &NotebookDao{
		group:   "default",
		table:   "notebook",
		columns: notebookColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NotebookDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NotebookDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NotebookDao) Columns() NotebookColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NotebookDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NotebookDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NotebookDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoteDao is the data access object for table note.
type NoteDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns NoteColumns // columns contains all the column names of Table for convenient usage.
}

// NoteColumns defines and stores column names for table note.
type NoteColumns struct {
	Id       string //
	Author   string // 笔记作者
	Title    string // 笔记标题
	Summary  string // 笔记概要
	Content  string // 笔记内容
	CreateAt string // 创建时间
	UpdateAt string // 更新时间
}

// noteColumns holds the columns for table note.
var noteColumns = NoteColumns{
	Id:       "id",
	Author:   "author",
	Title:    "title",
	Summary:  "summary",
	Content:  "content",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewNoteDao creates and returns a new DAO object for table data access.
func NewNoteDao() *NoteDao {
	return &NoteDao{
		group:   "default",
		table:   "note",
		columns: noteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoteDao) Columns() NoteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

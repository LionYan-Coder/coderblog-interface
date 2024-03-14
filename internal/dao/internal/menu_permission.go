// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuPermissionDao is the data access object for table menu_permission.
type MenuPermissionDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns MenuPermissionColumns // columns contains all the column names of Table for convenient usage.
}

// MenuPermissionColumns defines and stores column names for table menu_permission.
type MenuPermissionColumns struct {
	Id           string // 菜单-权限 关联id
	MenuId       string // 菜单id
	PermissionId string // 权限id
	CreateAt     string // 创建时间
	UpdateAt     string // 更新时间
}

// menuPermissionColumns holds the columns for table menu_permission.
var menuPermissionColumns = MenuPermissionColumns{
	Id:           "id",
	MenuId:       "menu_id",
	PermissionId: "permission_id",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
}

// NewMenuPermissionDao creates and returns a new DAO object for table data access.
func NewMenuPermissionDao() *MenuPermissionDao {
	return &MenuPermissionDao{
		group:   "default",
		table:   "menu_permission",
		columns: menuPermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MenuPermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MenuPermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MenuPermissionDao) Columns() MenuPermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MenuPermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MenuPermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MenuPermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

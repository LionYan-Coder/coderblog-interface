// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BannerDao is the data access object for table banner.
type BannerDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns BannerColumns // columns contains all the column names of Table for convenient usage.
}

// BannerColumns defines and stores column names for table banner.
type BannerColumns struct {
	Id          string // 横幅ID
	UserId      string // 用户id
	Title       string // 横幅标题
	Info        string // 横幅信息
	CoverUrl    string // 横幅图片地址
	DisplayDate string // 展示日期
	CreateAt    string // 创建日期
	UpdateAt    string // 更新日期
}

// bannerColumns holds the columns for table banner.
var bannerColumns = BannerColumns{
	Id:          "id",
	UserId:      "user_id",
	Title:       "title",
	Info:        "info",
	CoverUrl:    "cover_url",
	DisplayDate: "display_date",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
}

// NewBannerDao creates and returns a new DAO object for table data access.
func NewBannerDao() *BannerDao {
	return &BannerDao{
		group:   "default",
		table:   "banner",
		columns: bannerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BannerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BannerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BannerDao) Columns() BannerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BannerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BannerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BannerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

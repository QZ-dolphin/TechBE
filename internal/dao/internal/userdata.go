// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserdataDao is the data access object for table userdata.
type UserdataDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserdataColumns // columns contains all the column names of Table for convenient usage.
}

// UserdataColumns defines and stores column names for table userdata.
type UserdataColumns struct {
	Id        string //
	Username  string //
	Email     string //
	Password  string //
	ClientIp  string //
	Stackdata string //
	CreatedAt string //
	UpdatedAt string //
}

// userdataColumns holds the columns for table userdata.
var userdataColumns = UserdataColumns{
	Id:        "id",
	Username:  "username",
	Email:     "email",
	Password:  "password",
	ClientIp:  "client_ip",
	Stackdata: "stackdata",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserdataDao creates and returns a new DAO object for table data access.
func NewUserdataDao() *UserdataDao {
	return &UserdataDao{
		group:   "default",
		table:   "userdata",
		columns: userdataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserdataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserdataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserdataDao) Columns() UserdataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserdataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserdataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserdataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

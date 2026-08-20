package mysql

import (
	"context"
	"database/sql/driver"
)

func (stmt *mysqlStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	// ... existing logic ...
	rows, err := stmt.query(ctx, args)
	if err != nil {
		if ctx.Err() != nil {
			return nil, driver.ErrBadConn
		}
		return nil, err
	}
	return rows, nil
}

func (stmt *mysqlStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	// ... existing logic ...
	res, err := stmt.exec(ctx, args)
	if err != nil {
		if ctx.Err() != nil {
			return nil, driver.ErrBadConn
		}
		return nil, err
	}
	return res, nil
}
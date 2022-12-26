package postgres

import (
	"bytes"
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}
func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	qq, err := q.FormattedQuery()
	if err != nil {
		fmt.Println(err)

	}
	str1 := bytes.NewBuffer(qq).String()
	fmt.Println(str1)

	return nil
}
func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}

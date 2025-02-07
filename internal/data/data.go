package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stonelgh/kratos-layout/internal/conf"
)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

package filter

import "rec-sys/common"

type Filter interface {
	Filter([]*common.Product) []*common.Product
	Name() string
}

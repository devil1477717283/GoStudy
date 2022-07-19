package sort

import "rec-sys/common"

type Sorter interface {
	Sort([]*common.Product) []*common.Product
	Name() string
}

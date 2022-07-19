package recall

import "rec-sys/common"

type Recaller interface {
	Recall(n int) []*common.Product
	Name() string
}

package recall

import (
	"rec-sys/common"
	"sort"
)

var allProducts = []*common.Product{
	{Id: 1, Sale: 38, Name: "p1"},
	{Id: 2, Sale: 50, Name: "p2"},
	{Id: 3, Sale: 51, Name: "p3"},
	{Id: 4, Sale: 40, Name: "p4"},
	{Id: 5, Sale: 25, Name: "p5"},
	{Id: 6, Sale: 28, Name: "p6"},
	{Id: 7, Sale: 60, Name: "p7"},
}

type HotRecall struct {
	Tag string
}

func (h *HotRecall) Name() string {
	return h.Tag
}
func (h *HotRecall) Recall(n int) []*common.Product {
	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].Sale > allProducts[j].Sale
	})
	rect := make([]*common.Product, 0, n)
	for _, product := range allProducts {
		rect = append(rect, product)
		if len(rect) >= n {
			break
		}
	}
	return rect
}

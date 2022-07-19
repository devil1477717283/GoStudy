package main

import (
	"log"
	"rec-sys/common"
	"rec-sys/filter"
	"rec-sys/recall"
	"rec-sys/sort"
	"time"
)

type Recommender struct {
	Recallers []recall.Recaller
	Sorter    sort.Sorter
	Filters   []filter.Filter
}

func (rec *Recommender) Rec() []*common.Product {
	RecallMap := make(map[int]*common.Product, 100)
	for _, recaller := range rec.Recallers {
		begin := time.Now()
		products := recaller.Recall(10)
		log.Printf("召回%s耗时%dns,召回了%d个商品\n", recaller.Name(), time.Since(begin).Nanoseconds(), len(products))
		for _, product := range products {
			RecallMap[product.Id] = product
		}
	}
	log.Printf("一共召回了%d个商品\n", len(RecallMap))
	RecallSlice := make([]*common.Product, 0, len(RecallMap))
	for _, product := range RecallMap {
		RecallSlice = append(RecallSlice, product)
	}
	begin := time.Now()
	SortedResult := rec.Sorter.Sort(RecallSlice)
	log.Printf("排序耗时%dns", time.Since(begin).Nanoseconds())
	FilterResult := SortedResult
	for _, filter := range rec.Filters {
		begin := time.Now()
		FilterResult = filter.Filter(FilterResult)
		log.Printf("过滤规则%s耗时%dns\n", filter.Name(), time.Since(begin).Nanoseconds())
	}
	return FilterResult
}

func main() {

}

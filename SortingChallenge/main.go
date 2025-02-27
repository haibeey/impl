package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

type Product struct {
	Id         uint64     `json:"id"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Created    *time.Time `json:"created"`
	SalesCount uint64     `json:"sales_count"`
	ViewsCount uint64     `json:"views_count"`
}

// For printing to stdout. It prints only the ID of the product 
func (p Product) String() string {
	return fmt.Sprintf("%d", p.Id)
}

func (p1 Product) Equals(p2 Product) bool {
	if p1.Id != p2.Id || p1.Name != p2.Name || p1.Price != p2.Price ||
		p1.SalesCount != p2.SalesCount || p1.ViewsCount != p2.ViewsCount {
		return false
	}

	if (p1.Created == nil) != (p2.Created == nil) {
		return false
	}
	if p1.Created != nil && p2.Created != nil && !p1.Created.Equal(*p2.Created) {
		return false
	}

	return true
}

type SortBy func(p1, p2 *Product) bool

type Products struct {
	products []Product
	sortBy   SortBy
}

func (sb SortBy) Sort(products []Product) {
	ps := &Products{
		products: products,
		sortBy:   sb,
	}
	sort.Sort(ps)
}

func (ps *Products) Len() int {
	return len(ps.products)
}

func (ps *Products) Swap(i, j int) {
	ps.products[i], ps.products[j] = ps.products[j], ps.products[i]
}

func (ps *Products) Less(i, j int) bool {
	return ps.sortBy(&ps.products[i], &ps.products[j])
}

func parseDate(dateStr string) *time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil
	}
	return &t
}

func main() {
	products := []Product{
		{
			Id:         1,
			Name:       "Alabaster Table",
			Price:      12.99,
			Created:    parseDate("2019-01-04"),
			SalesCount: 32,
			ViewsCount: 730,
		},
		{
			Id:         2,
			Name:       "Zebra Table",
			Price:      44.49,
			Created:    parseDate("2012-01-04"),
			SalesCount: 301,
			ViewsCount: 3279,
		},
		{
			Id:         3,
			Name:       "Coffee Table",
			Price:      10.00,
			Created:    parseDate("2014-05-28"),
			SalesCount: 1048,
			ViewsCount: 20123,
		},
	}

	sortBySalesPerView := func(p1, p2 *Product) bool {
		ratio1 := float64(p1.SalesCount) / float64(p1.ViewsCount)
		ratio2 := float64(p2.SalesCount) / float64(p2.ViewsCount)
		return ratio1 < ratio2
	}

	SortBy(sortBySalesPerView).Sort(products)

	log.Println(products) 
}

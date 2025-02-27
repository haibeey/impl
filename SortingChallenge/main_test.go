package main

import (
	"testing"
)

func productEquals(ps1, ps2 []Product) bool {
	if len(ps1) != len(ps2) {
		return false
	}

	for i := range ps1 {
		if !ps1[i].Equals(ps2[i]) {
			return false
		}
	}

	return true
}

func TestSorter(t *testing.T) {

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

	sortByID := func(p1, p2 *Product) bool {
		return p1.Id < p2.Id
	}

	sortByName := func(p1, p2 *Product) bool {
		return p1.Name < p2.Name
	}

	sortByCreated := func(p1, p2 *Product) bool {
		return p1.Created.Before(*p2.Created)
	}
	sortBySalesCount := func(p1, p2 *Product) bool {
		return p1.SalesCount < p2.SalesCount
	}
	sortByViewsCount := func(p1, p2 *Product) bool {
		return p1.ViewsCount < p2.ViewsCount
	}

	productsUnsorted := make([]Product, len(products))
	copy(productsUnsorted, products) // copy data
	SortBy(sortByID).Sort(products)

	// should be the same as original
	if !productEquals(productsUnsorted, products) {
		t.Fatal("Sorted product has wrong ordering")
	}

	SortBy(sortByName).Sort(products)

	if products[1].Name != productsUnsorted[2].Name {
		t.Fatal("Sorted product has wrong ordering")
	}

	SortBy(sortByCreated).Sort(products)
	if products[0].Name != productsUnsorted[1].Name {
		t.Fatal("Product created 2012 should be first in the list")
	}

	SortBy(sortBySalesCount).Sort(products)
	// should be the same as original
	if !productEquals(productsUnsorted, products) {
		t.Fatal("Sorted product has wrong ordering")
	}

	SortBy(sortByViewsCount).Sort(products)
	// should be the same as original
	if !productEquals(productsUnsorted, products) {
		t.Fatal("Sorted product has wrong ordering")
	}
}

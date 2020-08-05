package product

import (
	"testing"
)

func TestNewProduct(t *testing.T) {
	const name = "Alcohol"
	const price = 120.20
	p, err := NewProduct(name, price)
	if err != nil {
		t.Fatalf("NewProduct error: %v", err)
	}
	if p.ID == "" {
		t.Error("NewProduct retorno ID vacio")
	}
	if p == nil {
		t.Error("NewProduct retorno product nil")
	}
	if p.Name != name {
		t.Errorf("NewProduct retorno producto con name %s pero debia ser %s", p.Name, name)
	}
	if p.Price != price {
		t.Errorf("NewProduct retorno producto con price %f pero debia ser %f", p.Price, price)
	}
}

func BenchmarkNewProduct(b *testing.B) {
	const name = "Alcohol"
	const price = 120.20
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := NewProduct(name, price)
		if err != nil {
			b.Fatalf("NewProduct error: %v", err)
		}
	}
}

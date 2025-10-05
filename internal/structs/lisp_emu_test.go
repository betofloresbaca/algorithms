package structs

import (
	"testing"
)

func TestConsCarCdr(t *testing.T) {
	// Construir la lista (1 2 3)
	l := Cons(1, Cons(2, Cons(3, nil)))

	if Car(l) != 1 {
		t.Errorf("Car(l) = %v; want 1", Car(l))
	}
	if Car(Cdr(l)) != 2 {
		t.Errorf("Car(Cdr(l)) = %v; want 2", Car(Cdr(l)))
	}
	if Car(Cdr(Cdr(l))) != 3 {
		t.Errorf("Car(Cdr(Cdr(l))) = %v; want 3", Car(Cdr(Cdr(l))))
	}
	if Cdr(Cdr(Cdr(l))) != nil {
		t.Error("Cdr(Cdr(Cdr(l))) should be nil")
	}
}

func TestEmptyList(t *testing.T) {
	var l *cell[int] = nil
	if Car(l) != 0 {
		t.Errorf("Car(nil) = %v; want 0 (zero value)", Car(l))
	}
	if Cdr(l) != nil {
		t.Error("Cdr(nil) should be nil")
	}
}

func TestConsLargeList(t *testing.T) {
	var l *cell[int] = nil
	// Construir la lista [0, 1, 2, ..., 99] (al revés por Cons)
	for i := 99; i >= 0; i-- {
		l = Cons(i, l)
	}
	// Recorrer y verificar
	n := l
	for i := 0; i < 100; i++ {
		if n == nil {
			t.Fatalf("List ended early at index %d", i)
		}
		if Car(n) != i {
			t.Errorf("Car(n) at index %d = %d; want %d", i, Car(n), i)
		}
		n = Cdr(n)
	}
	if n != nil {
		t.Error("List has more than 100 elements")
	}
}

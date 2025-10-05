package structs

import (
	"testing"
)

func TestPushAndString(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	expected := "[1 2 3]"
	if ll.String() != expected {
		t.Errorf("Push/String failed: got %s, want %s", ll.String(), expected)
	}
}

func TestUnshift(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Unshift(2)
	ll.Unshift(1)
	expected := "[1 2]"
	if ll.String() != expected {
		t.Errorf("Unshift failed: got %s, want %s", ll.String(), expected)
	}
}

func TestPop(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	v, err := ll.Pop()
	if err != nil || v != 3 {
		t.Errorf("Pop failed: got %v, err %v, want 3, nil", v, err)
	}
	expected := "[1 2]"
	if ll.String() != expected {
		t.Errorf("Pop failed: got %s, want %s", ll.String(), expected)
	}
}

func TestShift(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	v, err := ll.Shift()
	if err != nil || v != 1 {
		t.Errorf("Shift failed: got %v, err %v, want 1, nil", v, err)
	}
	expected := "[2 3]"
	if ll.String() != expected {
		t.Errorf("Shift failed: got %s, want %s", ll.String(), expected)
	}
}

func TestPopEmpty(t *testing.T) {
	ll := &LinkedList[int]{}
	_, err := ll.Pop()
	if err == nil {
		t.Error("Pop on empty list should return error")
	}
}

func TestShiftEmpty(t *testing.T) {
	ll := &LinkedList[int]{}
	_, err := ll.Shift()
	if err == nil {
		t.Error("Shift on empty list should return error")
	}
}

func TestCount(t *testing.T) {
	ll := &LinkedList[int]{}
	ll.Push(1)
	ll.Push(2)
	ll.Unshift(0)
	if ll.Count() != 3 {
		t.Errorf("Count failed: got %d, want 3", ll.Count())
	}
}

func TestStringEmpty(t *testing.T) {
	ll := &LinkedList[int]{}
	expected := "[]"
	if ll.String() != expected {
		t.Errorf("String empty failed: got %s, want %s", ll.String(), expected)
	}
}

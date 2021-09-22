package netflix

import "testing"

func TestLRUCache(t *testing.T) {
	cache := &LRUCache{
		capacity: 3,
		cache:    make(map[int]*DoubleLinkedListNode),
	}

	cache.Set(10, 20)
	want := "(10, 20)"
	if want != cache.Print() {
		t.Errorf("want %s, but get %s", want, cache.Print())
	}

	cache.Set(15, 25)
	want = "(10, 20)(15, 25)"
	if want != cache.Print() {
		t.Errorf("want %s, but get %s", want, cache.Print())
	}
	cache.Set(20, 35)
	want = "(10, 20)(15, 25)(20, 35)"
	if want != cache.Print() {
		t.Errorf("want %s, but get %s", want, cache.Print())
	}
	cache.Set(5, 40)
	want = "(15, 25)(20, 35)(5, 40)"
	if want != cache.Print() {
		t.Errorf("want %s, but get %s", want, cache.Print())
	}
	node := cache.Get(15)
	t.Log(node)
	// .
	want = "(20, 35)(5, 40)(15, 25)"
	if want != cache.Print() {
		t.Errorf("want %s, but get %s", want, cache.Print())
	}
}

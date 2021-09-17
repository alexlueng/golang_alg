package netflix

import (
	"fmt"
	"testing"
)

func TestMedianAge(t *testing.T) {
	mda := newMedianAge()
	mda.InsertNum(22)
	mda.InsertNum(35)
	want := float64((22 + 35) / 2.0)
	if mda.FindMedian() != want {
		t.Errorf("the media method doesn't work, got %f, want %f", mda.FindMedian(), want)
	}
	mda.InsertNum(30)
	want = float64(30)
	if mda.FindMedian() != want {
		t.Errorf("the media method doesn't work, got %f, want %f", mda.FindMedian(), want)
	}
	mda.InsertNum(25)

	fmt.Println(mda.maxHeap)
	fmt.Println(mda.minHeap)
	want = float64((25 + 30) / 2.0)
	if mda.FindMedian() != want {
		fmt.Println(mda.maxHeap)
		fmt.Println(mda.minHeap)
		t.Errorf("the media method doesn't work, got %f, want %f", mda.FindMedian(), want)
	}

}

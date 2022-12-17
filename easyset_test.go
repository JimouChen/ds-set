package ds_set

import (
	"fmt"
	"github.com/JimouChen/ds-set/easyset"
	"testing"
)

func TestClear(t *testing.T) {
	s := easyset.Set{}
	s.InitSet()
	s.Add("234")
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Show()
	s.Clear()
	s.Show()
	fmt.Println(s.Length())
}

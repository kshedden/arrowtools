package arrowtools

import (
	"testing"

	"github.com/apache/arrow/go/arrow/array"
)

func TestDropNA(t *testing.T) {

	tbl1 := table1()
	tbl2 := DropNA(tbl1)
	tbl6 := table6()

	b, msg := TablesEqual(tbl2, tbl6)
	if !b {
		println(msg)
		t.Fail()
	}
}

func TestFilterRows(t *testing.T) {

	sf := func(x []array.Interface, msk []bool) {
		c1 := array.NewInt64Data(x[0].Data())
		c2 := array.NewStringData(x[2].Data())

		for i := 0; i < c1.Len(); i++ {
			msk[i] = false
			if c1.Value(i) == 3 || c2.Value(i) == "str-7" {
				msk[i] = true
			}
		}
	}

	tbl1 := table1()
	tbl2 := FilterRows(tbl1, sf)
	tbl7 := table7()

	b, msg := TablesEqual(tbl2, tbl7)
	if !b {
		println(msg)
		t.Fail()
	}
}

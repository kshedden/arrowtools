package arrowtools

import (
	"testing"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

func TestAppendColumns1(t *testing.T) {

	x := [][]float64{
		[]float64{1, 2, 3},
		[]float64{4, 5},
		[]float64{6, 7, 8},
	}

	y := [][]float32{
		[]float32{1, 2, 3},
		[]float32{4, 5},
		[]float32{6, 7, 8},
	}

	col1 := ColumnFromFloat64Slices(x, "x")
	col2 := ColumnFromFloat32Slices(y, "y")

	tbl1 := TableFromColumns([]array.Column{*col1})
	tbl2 := AppendColumns(tbl1, []array.Column{*col2})
	tbl3 := TableFromColumns([]array.Column{*col1, *col2})

	b, msg := TablesEqual(tbl2, tbl3)
	if !b {
		println(msg)
		t.Fail()
	}
}

func TestDropColumns1(t *testing.T) {

	x := [][]float64{
		[]float64{1, 2, 3},
		[]float64{4, 5},
		[]float64{6, 7, 8},
	}

	y := [][]float32{
		[]float32{1, 2, 3},
		[]float32{4, 5},
		[]float32{6, 7, 8},
	}

	col1 := ColumnFromFloat64Slices(x, "x")
	col2 := ColumnFromFloat32Slices(y, "y")

	tbl1 := TableFromColumns([]array.Column{*col1, *col2})
	tbl2 := DropColumns(tbl1, "x")
	tbl3 := TableFromColumns([]array.Column{*col2})

	b, msg := TablesEqual(tbl2, tbl3)
	if !b {
		println(msg)
		t.Fail()
	}
}

func TestSelectColumns1(t *testing.T) {

	tbl1 := table1()
	tbl2 := table5()

	tbl3 := SelectColumns(tbl1, "str", "f64")

	b, msg := TablesEqual(tbl2, tbl3)
	if !b {
		println(msg)
		t.Fail()
	}
}

func TestReplaceColumns1(t *testing.T) {

	x := [][]float64{
		[]float64{1, 2, 3},
		[]float64{4, 5},
		[]float64{6, 7, 8},
	}

	y := [][]float32{
		[]float32{1, 2, 3},
		[]float32{4, 5},
		[]float32{6, 7, 8},
	}

	z := [][]float64{
		[]float64{1, 3, 2},
		[]float64{4, 5},
		[]float64{6, 7, 8},
	}

	col1 := ColumnFromFloat64Slices(x, "x")
	col2 := ColumnFromFloat32Slices(y, "y")
	col3 := ColumnFromFloat64Slices(z, "x")

	tbl1 := TableFromColumns([]array.Column{*col1, *col2})
	tbl2 := TableFromColumns([]array.Column{*col3, *col2})

	tbl3 := ReplaceColumn(tbl1, *col3)

	b, msg := TablesEqual(tbl2, tbl3)
	if !b {
		println(msg)
		t.Fail()
	}
}

func TestTableFromColumns1(t *testing.T) {

	x := [][]float64{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8},
	}

	y := [][]float32{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8},
	}

	z := [][]uint64{
		{1, 3, 2},
		{4, 5},
		{6, 7, 8},
	}

	col1 := ColumnFromFloat64Slices(x, "x")
	col2 := ColumnFromFloat32Slices(y, "y")
	col3 := ColumnFromUint64Slices(z, "z")

	tbl1 := TableFromColumns([]array.Column{*col1, *col2, *col3})

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "x", Type: arrow.PrimitiveTypes.Float64},
		{Name: "y", Type: arrow.PrimitiveTypes.Float32},
		{Name: "z", Type: arrow.PrimitiveTypes.Uint64},
	}, nil)

	tbl2 := array.NewTable(schema, []array.Column{*col1, *col2, *col3}, -1)

	b, msg := TablesEqual(tbl1, tbl2)
	if !b {
		println(msg)
		t.Fail()
	}
}

package arrowtools

import (
	"fmt"
	"strings"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

func ExampleColumnFromUint8Slices() {

	var sh SliceHelper

	col1 := sh.Uint8Column([][]uint8{{0, 1, 2}, {3, 4}}, nil, "col1")
	col2 := sh.Uint8Column([][]uint8{{0, 1, 2}, {3, 4}}, nil, "col2")
	col3 := sh.Uint8Column([][]uint8{{5, 2, 3}, {4, 1}}, nil, "col2")

	b1, _ := ColumnsEqual(col1, col2)
	b2, _ := ColumnsEqual(col1, col3)
	fmt.Printf("%t %t", b1, b2)
	// Output:
	// true false
}

func ExampleTableFromColumns() {

	var sh SliceHelper

	col1 := sh.Float32Column([][]float32{{4, 1, 3}, {2, 5}}, nil, "col1")
	col2 := sh.Float32Column([][]float32{{0, 1, 2}, {3, 4}}, nil, "col2")
	col3 := sh.Float32Column([][]float32{{5, 2, 3}, {4, 1}}, nil, "col3")

	tbl := TableFromColumns([]array.Column{*col1, *col2, *col3})
	fmt.Printf("%d %d", tbl.NumRows(), tbl.NumCols())
	// Output:
	// 5 3
}

func ExampleCSVWithHeader() {

	raw := `## a simple set of data: int64;float64;string
i64;f64;str
0;0;str-0
1;;str-1
2;2;str-2
3;3;str-3
;4;str-4
`
	buf := strings.NewReader(raw)

	fields := []arrow.Field{
		{Name: "str", Type: arrow.BinaryTypes.String},
		{Name: "i64", Type: arrow.PrimitiveTypes.Int64},
	}

	tbl := ReadCSV(buf, fields, WithHeader(), WithComma(';'), WithComment('#'))

	tr := array.NewTableReader(tbl, 1000)
	for tr.Next() {
		rec := tr.Record()
		rh := NewRecordHelper(rec)
		rh.SetPos(0)
		col1 := rh.StringSlice()
		fmt.Printf("%v\n", col1)
		rh.SetName("i64")
		col2 := rh.Int64Slice()
		fmt.Printf("%v\n", col2)
	}

	// Output:
	// [str-0 str-1 str-2 str-3 str-4]
	// [0 1 2 3 0]
}

func ExampleCSVWithoutHeader() {

	raw := `## a simple set of data: int64,float32,string
0,0,str-0
1,,str-1
2,2,str-2
3,3,str-3
,4,str-4
`
	buf := strings.NewReader(raw)

	fields := []arrow.Field{
		{Name: "i64", Type: arrow.PrimitiveTypes.Int64},
		{Name: "f32", Type: arrow.PrimitiveTypes.Float32},
		{Name: "str", Type: arrow.BinaryTypes.String},
	}

	tbl := ReadCSV(buf, fields, WithComment('#'))

	tr := array.NewTableReader(tbl, 1000)
	for tr.Next() {
		rec := tr.Record()
		rh := NewRecordHelper(rec)
		rh.SetPos(0)
		col1 := rh.Int64Slice()
		fmt.Printf("%v\n", col1)
		rh.SetPos(1)
		col2 := rh.Float32Slice()
		fmt.Printf("%v\n", col2)
		rh.SetName("str")
		col3 := rh.StringSlice()
		fmt.Printf("%v\n", col3)
	}

	// Output:
	// [0 1 2 3 0]
	// [0 0 2 3 4]
	// [str-0 str-1 str-2 str-3 str-4]
}

func ExampleDropNA() {

	var sh SliceHelper

	col1 := sh.Float64Column([][]float64{{0, 1, 2, 3}},
		[][]bool{{true, true, false, true}}, "x")
	col2 := sh.Int16Column([][]int16{{0, 1, 2, 3}},
		[][]bool{{false, true, false, true}}, "y")

	tbl1 := TableFromColumns([]array.Column{*col1, *col2})
	tbl2 := DropNA(tbl1)

	fmt.Printf("%d %d", tbl2.NumRows(), tbl2.NumCols())

	// Output:
	// 2 2
}

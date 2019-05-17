package arrowtools

import (
	"fmt"
	"strings"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

func ExampleColumnFromUint8Slices() {

	col1 := ColumnFromUint8Slices([][]uint8{{0, 1, 2}, {3, 4}}, "col1")
	col2 := ColumnFromUint8Slices([][]uint8{{0, 1, 2}, {3, 4}}, "col2")
	col3 := ColumnFromUint8Slices([][]uint8{{5, 2, 3}, {4, 1}}, "col2")

	b1, _ := ColumnsEqual(col1, col2)
	b2, _ := ColumnsEqual(col1, col3)
	fmt.Printf("%t %t", b1, b2)
	// Output:
	// true false
}

func ExampleTableFromColumns() {

	col1 := ColumnFromFloat32Slices([][]float32{{4, 1, 3}, {2, 5}}, "col1")
	col2 := ColumnFromFloat32Slices([][]float32{{0, 1, 2}, {3, 4}}, "col2")
	col3 := ColumnFromFloat32Slices([][]float32{{5, 2, 3}, {4, 1}}, "col3")

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
		col1 := GetStringSliceFromRecord(rec, 0)
		fmt.Printf("%v\n", col1)
		col2 := GetInt64SliceFromRecord(rec, 1)
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
		col1 := GetInt64SliceFromRecord(rec, 0)
		fmt.Printf("%v\n", col1)
		col2 := GetFloat32SliceFromRecord(rec, 1)
		fmt.Printf("%v\n", col2)
		col3 := GetStringSliceFromRecord(rec, 2)
		fmt.Printf("%v\n", col3)
	}

	// Output:
	// [0 1 2 3 0]
	// [0 0 2 3 4]
	// [str-0 str-1 str-2 str-3 str-4]
}

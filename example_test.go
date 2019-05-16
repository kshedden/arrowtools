package arrowtools

import (
	"fmt"

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

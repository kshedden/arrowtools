// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"testing"
)

func TestEqualityUint8(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]uint8{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]uint8{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]uint8{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]uint8{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromUint8Slices(w, nil, "name")
	col3 := ColumnFromUint8Slices(x, nil, "name")
	col4 := ColumnFromUint8Slices(y, nil, "name")
	col5 := ColumnFromUint8Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityUint16(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]uint16{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]uint16{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]uint16{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]uint16{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromUint16Slices(w, nil, "name")
	col3 := ColumnFromUint16Slices(x, nil, "name")
	col4 := ColumnFromUint16Slices(y, nil, "name")
	col5 := ColumnFromUint16Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityUint32(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]uint32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]uint32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]uint32{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]uint32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromUint32Slices(w, nil, "name")
	col3 := ColumnFromUint32Slices(x, nil, "name")
	col4 := ColumnFromUint32Slices(y, nil, "name")
	col5 := ColumnFromUint32Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityUint64(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]uint64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]uint64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]uint64{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]uint64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromUint64Slices(w, nil, "name")
	col3 := ColumnFromUint64Slices(x, nil, "name")
	col4 := ColumnFromUint64Slices(y, nil, "name")
	col5 := ColumnFromUint64Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityInt8(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]int8{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]int8{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]int8{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]int8{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromInt8Slices(w, nil, "name")
	col3 := ColumnFromInt8Slices(x, nil, "name")
	col4 := ColumnFromInt8Slices(y, nil, "name")
	col5 := ColumnFromInt8Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityInt16(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]int16{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]int16{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]int16{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]int16{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromInt16Slices(w, nil, "name")
	col3 := ColumnFromInt16Slices(x, nil, "name")
	col4 := ColumnFromInt16Slices(y, nil, "name")
	col5 := ColumnFromInt16Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityInt32(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]int32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]int32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]int32{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]int32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromInt32Slices(w, nil, "name")
	col3 := ColumnFromInt32Slices(x, nil, "name")
	col4 := ColumnFromInt32Slices(y, nil, "name")
	col5 := ColumnFromInt32Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityInt64(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]int64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]int64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]int64{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]int64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromInt64Slices(w, nil, "name")
	col3 := ColumnFromInt64Slices(x, nil, "name")
	col4 := ColumnFromInt64Slices(y, nil, "name")
	col5 := ColumnFromInt64Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityFloat32(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]float32{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromFloat32Slices(w, nil, "name")
	col3 := ColumnFromFloat32Slices(x, nil, "name")
	col4 := ColumnFromFloat32Slices(y, nil, "name")
	col5 := ColumnFromFloat32Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

func TestEqualityFloat64(t *testing.T) {

	v := [][]float32{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	w := [][]float64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	x := [][]float64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}

	y := [][]float64{
		{0, 1, 2, 4},
		{4, 5, 6},
		{7, 8},
	}

	z := [][]float64{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7, 8},
	}

	col1 := ColumnFromFloat32Slices(v, nil, "name")
	col2 := ColumnFromFloat64Slices(w, nil, "name")
	col3 := ColumnFromFloat64Slices(x, nil, "name")
	col4 := ColumnFromFloat64Slices(y, nil, "name")
	col5 := ColumnFromFloat64Slices(z, nil, "name")

	if b, _ := ColumnsEqual(col2, col5); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col5, col2); !b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col2, col3); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col3, col2); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col1, col4); b {
		t.Fail()
	}

	if b, _ := ColumnsEqual(col4, col1); b {
		t.Fail()
	}
}

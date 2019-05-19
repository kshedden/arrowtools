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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Uint8Column(w, nil, "name")
	col3 := sh.Uint8Column(x, nil, "name")
	col4 := sh.Uint8Column(y, nil, "name")
	col5 := sh.Uint8Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Uint16Column(w, nil, "name")
	col3 := sh.Uint16Column(x, nil, "name")
	col4 := sh.Uint16Column(y, nil, "name")
	col5 := sh.Uint16Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Uint32Column(w, nil, "name")
	col3 := sh.Uint32Column(x, nil, "name")
	col4 := sh.Uint32Column(y, nil, "name")
	col5 := sh.Uint32Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Uint64Column(w, nil, "name")
	col3 := sh.Uint64Column(x, nil, "name")
	col4 := sh.Uint64Column(y, nil, "name")
	col5 := sh.Uint64Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Int8Column(w, nil, "name")
	col3 := sh.Int8Column(x, nil, "name")
	col4 := sh.Int8Column(y, nil, "name")
	col5 := sh.Int8Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Int16Column(w, nil, "name")
	col3 := sh.Int16Column(x, nil, "name")
	col4 := sh.Int16Column(y, nil, "name")
	col5 := sh.Int16Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Int32Column(w, nil, "name")
	col3 := sh.Int32Column(x, nil, "name")
	col4 := sh.Int32Column(y, nil, "name")
	col5 := sh.Int32Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Int64Column(w, nil, "name")
	col3 := sh.Int64Column(x, nil, "name")
	col4 := sh.Int64Column(y, nil, "name")
	col5 := sh.Int64Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Float32Column(w, nil, "name")
	col3 := sh.Float32Column(x, nil, "name")
	col4 := sh.Float32Column(y, nil, "name")
	col5 := sh.Float32Column(z, nil, "name")

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

	var sh SliceHelper

	col1 := sh.Float32Column(v, nil, "name")
	col2 := sh.Float64Column(w, nil, "name")
	col3 := sh.Float64Column(x, nil, "name")
	col4 := sh.Float64Column(y, nil, "name")
	col5 := sh.Float64Column(z, nil, "name")

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

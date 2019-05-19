// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"testing"
)

func TestColumnSliceUint8(t *testing.T) {

	x := [][]uint8{
		[]uint8{0, 1, 2, 3, 4},
		[]uint8{5, 6, 7},
		[]uint8{8, 9},
	}

	var sh SliceHelper
	col := sh.Uint8Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Uint8Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceUint16(t *testing.T) {

	x := [][]uint16{
		[]uint16{0, 1, 2, 3, 4},
		[]uint16{5, 6, 7},
		[]uint16{8, 9},
	}

	var sh SliceHelper
	col := sh.Uint16Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Uint16Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceUint32(t *testing.T) {

	x := [][]uint32{
		[]uint32{0, 1, 2, 3, 4},
		[]uint32{5, 6, 7},
		[]uint32{8, 9},
	}

	var sh SliceHelper
	col := sh.Uint32Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Uint32Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceUint64(t *testing.T) {

	x := [][]uint64{
		[]uint64{0, 1, 2, 3, 4},
		[]uint64{5, 6, 7},
		[]uint64{8, 9},
	}

	var sh SliceHelper
	col := sh.Uint64Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Uint64Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceInt8(t *testing.T) {

	x := [][]int8{
		[]int8{0, 1, 2, 3, 4},
		[]int8{5, 6, 7},
		[]int8{8, 9},
	}

	var sh SliceHelper
	col := sh.Int8Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Int8Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceInt16(t *testing.T) {

	x := [][]int16{
		[]int16{0, 1, 2, 3, 4},
		[]int16{5, 6, 7},
		[]int16{8, 9},
	}

	var sh SliceHelper
	col := sh.Int16Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Int16Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceInt32(t *testing.T) {

	x := [][]int32{
		[]int32{0, 1, 2, 3, 4},
		[]int32{5, 6, 7},
		[]int32{8, 9},
	}

	var sh SliceHelper
	col := sh.Int32Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Int32Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceInt64(t *testing.T) {

	x := [][]int64{
		[]int64{0, 1, 2, 3, 4},
		[]int64{5, 6, 7},
		[]int64{8, 9},
	}

	var sh SliceHelper
	col := sh.Int64Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Int64Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceFloat32(t *testing.T) {

	x := [][]float32{
		[]float32{0, 1, 2, 3, 4},
		[]float32{5, 6, 7},
		[]float32{8, 9},
	}

	var sh SliceHelper
	col := sh.Float32Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Float32Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

func TestColumnSliceFloat64(t *testing.T) {

	x := [][]float64{
		[]float64{0, 1, 2, 3, 4},
		[]float64{5, 6, 7},
		[]float64{8, 9},
	}

	var sh SliceHelper
	col := sh.Float64Column(x, nil, "name")

	ch := NewColumnHelper(col)
	y := ch.Float64Slices()

	if len(x) != len(y) {
		t.Fail()
	}

	for i := range x {
		for j := range x[i] {
			if x[i][j] != y[i][j] {
				t.Fail()
			}
		}
	}
}

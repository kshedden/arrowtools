// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"github.com/apache/arrow/go/arrow/array"
	"testing"
)

func TestGetSliceFromRecord(t *testing.T) {

	var sh SliceHelper

	var cols []array.Column
	{
		x := [][]uint8{{3, 4, 5}, {6, 7}}
		col := sh.Uint8Column(x, nil, "Uint8")
		cols = append(cols, *col)
	}
	{
		x := [][]uint16{{3, 4, 5}, {6, 7}}
		col := sh.Uint16Column(x, nil, "Uint16")
		cols = append(cols, *col)
	}
	{
		x := [][]uint32{{3, 4, 5}, {6, 7}}
		col := sh.Uint32Column(x, nil, "Uint32")
		cols = append(cols, *col)
	}
	{
		x := [][]uint64{{3, 4, 5}, {6, 7}}
		col := sh.Uint64Column(x, nil, "Uint64")
		cols = append(cols, *col)
	}
	{
		x := [][]int8{{3, 4, 5}, {6, 7}}
		col := sh.Int8Column(x, nil, "Int8")
		cols = append(cols, *col)
	}
	{
		x := [][]int16{{3, 4, 5}, {6, 7}}
		col := sh.Int16Column(x, nil, "Int16")
		cols = append(cols, *col)
	}
	{
		x := [][]int32{{3, 4, 5}, {6, 7}}
		col := sh.Int32Column(x, nil, "Int32")
		cols = append(cols, *col)
	}
	{
		x := [][]int64{{3, 4, 5}, {6, 7}}
		col := sh.Int64Column(x, nil, "Int64")
		cols = append(cols, *col)
	}
	{
		x := [][]float32{{3, 4, 5}, {6, 7}}
		col := sh.Float32Column(x, nil, "Float32")
		cols = append(cols, *col)
	}
	{
		x := [][]float64{{3, 4, 5}, {6, 7}}
		col := sh.Float64Column(x, nil, "Float64")
		cols = append(cols, *col)
	}

	tbl := TableFromColumns(cols)
	tr := array.NewTableReader(tbl, 1000)
	v := 3

	for tr.Next() {

		rec := tr.Record()
		rh := NewRecordHelper(rec)
		j := 0
		{
			rh.SetPos(j)
			x := rh.Uint8Slice()
			for i := range x {
				if x[i] != uint8(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Uint16Slice()
			for i := range x {
				if x[i] != uint16(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Uint32Slice()
			for i := range x {
				if x[i] != uint32(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Uint64Slice()
			for i := range x {
				if x[i] != uint64(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Int8Slice()
			for i := range x {
				if x[i] != int8(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Int16Slice()
			for i := range x {
				if x[i] != int16(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Int32Slice()
			for i := range x {
				if x[i] != int32(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Int64Slice()
			for i := range x {
				if x[i] != int64(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Float32Slice()
			for i := range x {
				if x[i] != float32(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			rh.SetPos(j)
			x := rh.Float64Slice()
			for i := range x {
				if x[i] != float64(v+i) {
					t.Fail()
				}
			}
			j++
		}
		v += 3
	}
}

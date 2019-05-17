// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"github.com/apache/arrow/go/arrow/array"
	"testing"
)

func TestGetSliceFromRecord(t *testing.T) {

	var cols []array.Column
	{
		x := [][]uint8{{3, 4, 5}, {6, 7}}
		col := ColumnFromUint8Slices(x, nil, "Uint8")
		cols = append(cols, *col)
	}
	{
		x := [][]uint16{{3, 4, 5}, {6, 7}}
		col := ColumnFromUint16Slices(x, nil, "Uint16")
		cols = append(cols, *col)
	}
	{
		x := [][]uint32{{3, 4, 5}, {6, 7}}
		col := ColumnFromUint32Slices(x, nil, "Uint32")
		cols = append(cols, *col)
	}
	{
		x := [][]uint64{{3, 4, 5}, {6, 7}}
		col := ColumnFromUint64Slices(x, nil, "Uint64")
		cols = append(cols, *col)
	}
	{
		x := [][]int8{{3, 4, 5}, {6, 7}}
		col := ColumnFromInt8Slices(x, nil, "Int8")
		cols = append(cols, *col)
	}
	{
		x := [][]int16{{3, 4, 5}, {6, 7}}
		col := ColumnFromInt16Slices(x, nil, "Int16")
		cols = append(cols, *col)
	}
	{
		x := [][]int32{{3, 4, 5}, {6, 7}}
		col := ColumnFromInt32Slices(x, nil, "Int32")
		cols = append(cols, *col)
	}
	{
		x := [][]int64{{3, 4, 5}, {6, 7}}
		col := ColumnFromInt64Slices(x, nil, "Int64")
		cols = append(cols, *col)
	}
	{
		x := [][]float32{{3, 4, 5}, {6, 7}}
		col := ColumnFromFloat32Slices(x, nil, "Float32")
		cols = append(cols, *col)
	}
	{
		x := [][]float64{{3, 4, 5}, {6, 7}}
		col := ColumnFromFloat64Slices(x, nil, "Float64")
		cols = append(cols, *col)
	}

	tbl := TableFromColumns(cols)
	tr := array.NewTableReader(tbl, 1000)
	v := 3

	for tr.Next() {

		rec := tr.Record()
		j := 0
		{
			x := GetUint8SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != uint8(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetUint16SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != uint16(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetUint32SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != uint32(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetUint64SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != uint64(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetInt8SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != int8(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetInt16SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != int16(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetInt32SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != int32(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetInt64SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != int64(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetFloat32SliceFromRecord(rec, j)
			for i := range x {
				if x[i] != float32(v+i) {
					t.Fail()
				}
			}
			j++
		}
		{
			x := GetFloat64SliceFromRecord(rec, j)
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

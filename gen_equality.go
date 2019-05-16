// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/stretchr/testify/assert"
)

func ColumnsEqual(col1, col2 *array.Column) (bool, string) {

	if col1.DataType().ID() != col2.DataType().ID() {
		return false, "Inconsistent types"
	}

	chunks1 := col1.Data().Chunks()
	chunks2 := col2.Data().Chunks()

	if len(chunks1) != len(chunks2) {
		return false, fmt.Sprintf("Unequal chunk counts, %d != %d", len(chunks1), len(chunks2))
	}

	for k := range chunks1 {

		chunk1 := chunks1[k]
		chunk2 := chunks2[k]

		switch col1.DataType() {
		case arrow.PrimitiveTypes.Uint8:
			y1 := array.NewUint8Data(chunk1.Data()).Uint8Values()
			y2 := array.NewUint8Data(chunk2.Data()).Uint8Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal uint8 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Uint16:
			y1 := array.NewUint16Data(chunk1.Data()).Uint16Values()
			y2 := array.NewUint16Data(chunk2.Data()).Uint16Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal uint16 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Uint32:
			y1 := array.NewUint32Data(chunk1.Data()).Uint32Values()
			y2 := array.NewUint32Data(chunk2.Data()).Uint32Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal uint32 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Uint64:
			y1 := array.NewUint64Data(chunk1.Data()).Uint64Values()
			y2 := array.NewUint64Data(chunk2.Data()).Uint64Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal uint64 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Int8:
			y1 := array.NewInt8Data(chunk1.Data()).Int8Values()
			y2 := array.NewInt8Data(chunk2.Data()).Int8Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal int8 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Int16:
			y1 := array.NewInt16Data(chunk1.Data()).Int16Values()
			y2 := array.NewInt16Data(chunk2.Data()).Int16Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal int16 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Int32:
			y1 := array.NewInt32Data(chunk1.Data()).Int32Values()
			y2 := array.NewInt32Data(chunk2.Data()).Int32Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal int32 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Int64:
			y1 := array.NewInt64Data(chunk1.Data()).Int64Values()
			y2 := array.NewInt64Data(chunk2.Data()).Int64Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal int64 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Float32:
			y1 := array.NewFloat32Data(chunk1.Data()).Float32Values()
			y2 := array.NewFloat32Data(chunk2.Data()).Float32Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal float32 values in chunk %d\n", k)
			}
		case arrow.PrimitiveTypes.Float64:
			y1 := array.NewFloat64Data(chunk1.Data()).Float64Values()
			y2 := array.NewFloat64Data(chunk2.Data()).Float64Values()
			if !assert.ObjectsAreEqualValues(y1, y2) {
				return false, fmt.Sprintf("Unequal float64 values in chunk %d\n", k)
			}
		case arrow.BinaryTypes.String:
			y1 := array.NewStringData(chunk1.Data())
			y2 := array.NewStringData(chunk2.Data())
			if y1.Len() != y2.Len() {
				return false, fmt.Sprintf("Unequal lengths of string values in chunk %d\n", k)
			}
			for i := 0; i < y1.Len(); i++ {
				if y1.Value(i) != y2.Value(i) {
					return false, fmt.Sprintf("Unequal string values in chunk %d\n", k)
				}
			}
		default:
			panic("unknown type")
		}
	}

	return true, ""
}

func TablesEqual(tbl1, tbl2 array.Table) (bool, string) {

	m1 := tbl1.NumCols()
	m2 := tbl2.NumCols()
	if m1 != m2 {
		return false, fmt.Sprintf("Inconsistent number of columns, %d != %d", m1, m2)
	}

	for i := 0; i < int(m1); i++ {

		col1 := tbl1.Column(i)
		col2 := tbl2.Column(i)

		b, msg := ColumnsEqual(col1, col2)
		if !b {
			return false, msg
		}
	}

	return true, ""
}

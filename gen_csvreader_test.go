// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"bytes"
	"io"
	"testing"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

func TestCSV1(t *testing.T) {

	var buf bytes.Buffer
	io.WriteString(&buf, "xfloat64,xfloat32,xint64,xint32,xint16,xint8,xuint64,xuint32,xuint16,xuint8\n")
	io.WriteString(&buf, "0,0,0,0,0,0,0,0,0,0\n")
	io.WriteString(&buf, "1,1,1,1,1,1,1,1,1,1\n")
	io.WriteString(&buf, "2,2,2,2,2,2,2,2,2,2\n")
	io.WriteString(&buf, "3,3,3,3,3,3,3,3,3,3\n")

	xuint8 := ColumnFromUint8Slices([][]uint8{{0, 1, 2, 3}}, nil, "xuint8")
	xuint16 := ColumnFromUint16Slices([][]uint16{{0, 1, 2, 3}}, nil, "xuint16")
	xuint32 := ColumnFromUint32Slices([][]uint32{{0, 1, 2, 3}}, nil, "xuint32")
	xuint64 := ColumnFromUint64Slices([][]uint64{{0, 1, 2, 3}}, nil, "xuint64")
	xint8 := ColumnFromInt8Slices([][]int8{{0, 1, 2, 3}}, nil, "xint8")
	xint16 := ColumnFromInt16Slices([][]int16{{0, 1, 2, 3}}, nil, "xint16")
	xint32 := ColumnFromInt32Slices([][]int32{{0, 1, 2, 3}}, nil, "xint32")
	xint64 := ColumnFromInt64Slices([][]int64{{0, 1, 2, 3}}, nil, "xint64")
	xfloat32 := ColumnFromFloat32Slices([][]float32{{0, 1, 2, 3}}, nil, "xfloat32")
	xfloat64 := ColumnFromFloat64Slices([][]float64{{0, 1, 2, 3}}, nil, "xfloat64")

	cols := []array.Column{
		*xuint8,
		*xuint16,
		*xuint32,
		*xuint64,
		*xint8,
		*xint16,
		*xint32,
		*xint64,
		*xfloat32,
		*xfloat64,
	}

	fields := []arrow.Field{
		{Name: "xuint8", Type: arrow.PrimitiveTypes.Uint8},
		{Name: "xuint16", Type: arrow.PrimitiveTypes.Uint16},
		{Name: "xuint32", Type: arrow.PrimitiveTypes.Uint32},
		{Name: "xuint64", Type: arrow.PrimitiveTypes.Uint64},
		{Name: "xint8", Type: arrow.PrimitiveTypes.Int8},
		{Name: "xint16", Type: arrow.PrimitiveTypes.Int16},
		{Name: "xint32", Type: arrow.PrimitiveTypes.Int32},
		{Name: "xint64", Type: arrow.PrimitiveTypes.Int64},
		{Name: "xfloat32", Type: arrow.PrimitiveTypes.Float32},
		{Name: "xfloat64", Type: arrow.PrimitiveTypes.Float64},
	}

	tbl1 := ReadCSV(&buf, fields, WithHeader())
	tbl2 := array.NewTable(arrow.NewSchema(fields, nil), cols, -1)

	if b, msg := TablesEqual(tbl1, tbl2); !b {
		println(msg)
		t.Fail()
	}
}

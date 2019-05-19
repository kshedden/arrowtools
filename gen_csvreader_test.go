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

	var sh SliceHelper

	xuint8 := sh.Uint8Column([][]uint8{{0, 1, 2, 3}}, nil, "xuint8")
	xuint16 := sh.Uint16Column([][]uint16{{0, 1, 2, 3}}, nil, "xuint16")
	xuint32 := sh.Uint32Column([][]uint32{{0, 1, 2, 3}}, nil, "xuint32")
	xuint64 := sh.Uint64Column([][]uint64{{0, 1, 2, 3}}, nil, "xuint64")
	xint8 := sh.Int8Column([][]int8{{0, 1, 2, 3}}, nil, "xint8")
	xint16 := sh.Int16Column([][]int16{{0, 1, 2, 3}}, nil, "xint16")
	xint32 := sh.Int32Column([][]int32{{0, 1, 2, 3}}, nil, "xint32")
	xint64 := sh.Int64Column([][]int64{{0, 1, 2, 3}}, nil, "xint64")
	xfloat32 := sh.Float32Column([][]float32{{0, 1, 2, 3}}, nil, "xfloat32")
	xfloat64 := sh.Float64Column([][]float64{{0, 1, 2, 3}}, nil, "xfloat64")

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

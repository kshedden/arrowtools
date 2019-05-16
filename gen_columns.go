// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

// ColumnFromUint8Slices returns a pointer to an array.Column value that
// holds the given uint8 data.
func ColumnFromUint8Slices(x [][]uint8, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Uint8}
	bld := array.NewUint8Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromUint16Slices returns a pointer to an array.Column value that
// holds the given uint16 data.
func ColumnFromUint16Slices(x [][]uint16, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Uint16}
	bld := array.NewUint16Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromUint32Slices returns a pointer to an array.Column value that
// holds the given uint32 data.
func ColumnFromUint32Slices(x [][]uint32, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Uint32}
	bld := array.NewUint32Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromUint64Slices returns a pointer to an array.Column value that
// holds the given uint64 data.
func ColumnFromUint64Slices(x [][]uint64, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Uint64}
	bld := array.NewUint64Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromInt8Slices returns a pointer to an array.Column value that
// holds the given int8 data.
func ColumnFromInt8Slices(x [][]int8, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int8}
	bld := array.NewInt8Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromInt16Slices returns a pointer to an array.Column value that
// holds the given int16 data.
func ColumnFromInt16Slices(x [][]int16, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int16}
	bld := array.NewInt16Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromInt32Slices returns a pointer to an array.Column value that
// holds the given int32 data.
func ColumnFromInt32Slices(x [][]int32, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int32}
	bld := array.NewInt32Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromInt64Slices returns a pointer to an array.Column value that
// holds the given int64 data.
func ColumnFromInt64Slices(x [][]int64, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int64}
	bld := array.NewInt64Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromFloat32Slices returns a pointer to an array.Column value that
// holds the given float32 data.
func ColumnFromFloat32Slices(x [][]float32, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Float32}
	bld := array.NewFloat32Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromFloat64Slices returns a pointer to an array.Column value that
// holds the given float64 data.
func ColumnFromFloat64Slices(x [][]float64, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Float64}
	bld := array.NewFloat64Builder(mem)
	for _, z := range x {
		bld.AppendValues(z, nil)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// SlicesFromUint8Column returns a slice of slices holding the
// data from the given column.
func SlicesFromUint8Column(col *array.Column) [][]uint8 {

	var x [][]uint8

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewUint8Data(c.Data()).Uint8Values())
	}

	return x
}

// SlicesFromUint16Column returns a slice of slices holding the
// data from the given column.
func SlicesFromUint16Column(col *array.Column) [][]uint16 {

	var x [][]uint16

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewUint16Data(c.Data()).Uint16Values())
	}

	return x
}

// SlicesFromUint32Column returns a slice of slices holding the
// data from the given column.
func SlicesFromUint32Column(col *array.Column) [][]uint32 {

	var x [][]uint32

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewUint32Data(c.Data()).Uint32Values())
	}

	return x
}

// SlicesFromUint64Column returns a slice of slices holding the
// data from the given column.
func SlicesFromUint64Column(col *array.Column) [][]uint64 {

	var x [][]uint64

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewUint64Data(c.Data()).Uint64Values())
	}

	return x
}

// SlicesFromInt8Column returns a slice of slices holding the
// data from the given column.
func SlicesFromInt8Column(col *array.Column) [][]int8 {

	var x [][]int8

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewInt8Data(c.Data()).Int8Values())
	}

	return x
}

// SlicesFromInt16Column returns a slice of slices holding the
// data from the given column.
func SlicesFromInt16Column(col *array.Column) [][]int16 {

	var x [][]int16

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewInt16Data(c.Data()).Int16Values())
	}

	return x
}

// SlicesFromInt32Column returns a slice of slices holding the
// data from the given column.
func SlicesFromInt32Column(col *array.Column) [][]int32 {

	var x [][]int32

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewInt32Data(c.Data()).Int32Values())
	}

	return x
}

// SlicesFromInt64Column returns a slice of slices holding the
// data from the given column.
func SlicesFromInt64Column(col *array.Column) [][]int64 {

	var x [][]int64

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewInt64Data(c.Data()).Int64Values())
	}

	return x
}

// SlicesFromFloat32Column returns a slice of slices holding the
// data from the given column.
func SlicesFromFloat32Column(col *array.Column) [][]float32 {

	var x [][]float32

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewFloat32Data(c.Data()).Float32Values())
	}

	return x
}

// SlicesFromFloat64Column returns a slice of slices holding the
// data from the given column.
func SlicesFromFloat64Column(col *array.Column) [][]float64 {

	var x [][]float64

	for _, c := range col.Data().Chunks() {
		x = append(x, array.NewFloat64Data(c.Data()).Float64Values())
	}

	return x
}

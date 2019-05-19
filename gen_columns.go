// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

// ColumnFromUint8Slices returns a pointer to an array.Column value that
// holds the given uint8 data.
func (sh *SliceHelper) Uint8Column(x [][]uint8, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Uint8}
	bld := array.NewUint8Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromUint16Slices returns a pointer to an array.Column value that
// holds the given uint16 data.
func (sh *SliceHelper) Uint16Column(x [][]uint16, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Uint16}
	bld := array.NewUint16Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromUint32Slices returns a pointer to an array.Column value that
// holds the given uint32 data.
func (sh *SliceHelper) Uint32Column(x [][]uint32, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Uint32}
	bld := array.NewUint32Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromUint64Slices returns a pointer to an array.Column value that
// holds the given uint64 data.
func (sh *SliceHelper) Uint64Column(x [][]uint64, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Uint64}
	bld := array.NewUint64Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromInt8Slices returns a pointer to an array.Column value that
// holds the given int8 data.
func (sh *SliceHelper) Int8Column(x [][]int8, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int8}
	bld := array.NewInt8Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromInt16Slices returns a pointer to an array.Column value that
// holds the given int16 data.
func (sh *SliceHelper) Int16Column(x [][]int16, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int16}
	bld := array.NewInt16Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromInt32Slices returns a pointer to an array.Column value that
// holds the given int32 data.
func (sh *SliceHelper) Int32Column(x [][]int32, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int32}
	bld := array.NewInt32Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromInt64Slices returns a pointer to an array.Column value that
// holds the given int64 data.
func (sh *SliceHelper) Int64Column(x [][]int64, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Int64}
	bld := array.NewInt64Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromFloat32Slices returns a pointer to an array.Column value that
// holds the given float32 data.
func (sh *SliceHelper) Float32Column(x [][]float32, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Float32}
	bld := array.NewFloat32Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// ColumnFromFloat64Slices returns a pointer to an array.Column value that
// holds the given float64 data.
func (sh *SliceHelper) Float64Column(x [][]float64, valid [][]bool, name string) *array.Column {

	mem := memory.DefaultAllocator
	var y []array.Interface

	fld := arrow.Field{Name: name, Type: arrow.PrimitiveTypes.Float64}
	bld := array.NewFloat64Builder(mem)
	for i, z := range x {
		var v []bool
		if valid != nil {
			v = valid[i]
		}
		bld.AppendValues(z, v)
		y = append(y, bld.NewArray())
	}

	chunks := array.NewChunked(fld.Type, y)
	return array.NewColumn(fld, chunks)
}

// Uint8Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Uint8Slices() [][]uint8 {
	var x [][]uint8
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewUint8Data(c.Data()).Uint8Values())
	}
	return x
}

// Uint16Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Uint16Slices() [][]uint16 {
	var x [][]uint16
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewUint16Data(c.Data()).Uint16Values())
	}
	return x
}

// Uint32Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Uint32Slices() [][]uint32 {
	var x [][]uint32
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewUint32Data(c.Data()).Uint32Values())
	}
	return x
}

// Uint64Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Uint64Slices() [][]uint64 {
	var x [][]uint64
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewUint64Data(c.Data()).Uint64Values())
	}
	return x
}

// Int8Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Int8Slices() [][]int8 {
	var x [][]int8
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewInt8Data(c.Data()).Int8Values())
	}
	return x
}

// Int16Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Int16Slices() [][]int16 {
	var x [][]int16
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewInt16Data(c.Data()).Int16Values())
	}
	return x
}

// Int32Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Int32Slices() [][]int32 {
	var x [][]int32
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewInt32Data(c.Data()).Int32Values())
	}
	return x
}

// Int64Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Int64Slices() [][]int64 {
	var x [][]int64
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewInt64Data(c.Data()).Int64Values())
	}
	return x
}

// Float32Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Float32Slices() [][]float32 {
	var x [][]float32
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewFloat32Data(c.Data()).Float32Values())
	}
	return x
}

// Float64Column returns a slice of slices holding the
// data from the given column.
func (ch *ColumnHelper) Float64Slices() [][]float64 {
	var x [][]float64
	for _, c := range ch.col.Data().Chunks() {
		x = append(x, array.NewFloat64Data(c.Data()).Float64Values())
	}
	return x
}

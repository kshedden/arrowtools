// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"github.com/apache/arrow/go/arrow/array"
)

// GetUint8SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Uint8Slice() []uint8 {
	return array.NewUint8Data(rh.rec.Column(rh.curpos).Data()).Uint8Values()
}

// GetUint16SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Uint16Slice() []uint16 {
	return array.NewUint16Data(rh.rec.Column(rh.curpos).Data()).Uint16Values()
}

// GetUint32SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Uint32Slice() []uint32 {
	return array.NewUint32Data(rh.rec.Column(rh.curpos).Data()).Uint32Values()
}

// GetUint64SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Uint64Slice() []uint64 {
	return array.NewUint64Data(rh.rec.Column(rh.curpos).Data()).Uint64Values()
}

// GetInt8SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Int8Slice() []int8 {
	return array.NewInt8Data(rh.rec.Column(rh.curpos).Data()).Int8Values()
}

// GetInt16SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Int16Slice() []int16 {
	return array.NewInt16Data(rh.rec.Column(rh.curpos).Data()).Int16Values()
}

// GetInt32SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Int32Slice() []int32 {
	return array.NewInt32Data(rh.rec.Column(rh.curpos).Data()).Int32Values()
}

// GetInt64SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Int64Slice() []int64 {
	return array.NewInt64Data(rh.rec.Column(rh.curpos).Data()).Int64Values()
}

// GetFloat32SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Float32Slice() []float32 {
	return array.NewFloat32Data(rh.rec.Column(rh.curpos).Data()).Float32Values()
}

// GetFloat64SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func (rh *RecordHelper) Float64Slice() []float64 {
	return array.NewFloat64Data(rh.rec.Column(rh.curpos).Data()).Float64Values()
}

func (rh *RecordHelper) StringSlice() []string {
	// TODO probably not the best way to do this.
	da := array.NewStringData(rh.rec.Column(rh.curpos).Data())
	vals := make([]string, da.Len())
	for i := 0; i < da.Len(); i++ {
		vals[i] = da.Value(i)
	}
	return vals
}

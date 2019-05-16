// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"github.com/apache/arrow/go/arrow/array"
)

// GetUint8SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetUint8SliceFromRecord(rec array.Record, pos int) []uint8 {
	return array.NewUint8Data(rec.Column(pos).Data()).Uint8Values()
}

// GetUint16SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetUint16SliceFromRecord(rec array.Record, pos int) []uint16 {
	return array.NewUint16Data(rec.Column(pos).Data()).Uint16Values()
}

// GetUint32SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetUint32SliceFromRecord(rec array.Record, pos int) []uint32 {
	return array.NewUint32Data(rec.Column(pos).Data()).Uint32Values()
}

// GetUint64SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetUint64SliceFromRecord(rec array.Record, pos int) []uint64 {
	return array.NewUint64Data(rec.Column(pos).Data()).Uint64Values()
}

// GetInt8SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetInt8SliceFromRecord(rec array.Record, pos int) []int8 {
	return array.NewInt8Data(rec.Column(pos).Data()).Int8Values()
}

// GetInt16SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetInt16SliceFromRecord(rec array.Record, pos int) []int16 {
	return array.NewInt16Data(rec.Column(pos).Data()).Int16Values()
}

// GetInt32SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetInt32SliceFromRecord(rec array.Record, pos int) []int32 {
	return array.NewInt32Data(rec.Column(pos).Data()).Int32Values()
}

// GetInt64SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetInt64SliceFromRecord(rec array.Record, pos int) []int64 {
	return array.NewInt64Data(rec.Column(pos).Data()).Int64Values()
}

// GetFloat32SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetFloat32SliceFromRecord(rec array.Record, pos int) []float32 {
	return array.NewFloat32Data(rec.Column(pos).Data()).Float32Values()
}

// GetFloat64SliceFromRecord returns a slice corresponding to the given
// column position in the given record.
func GetFloat64SliceFromRecord(rec array.Record, pos int) []float64 {
	return array.NewFloat64Data(rec.Column(pos).Data()).Float64Values()
}

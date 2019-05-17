// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

func getnull(col array.Interface, msk []bool, da *array.Data) {

	switch col.DataType() {
	case arrow.PrimitiveTypes.Uint8:
		x := array.NewUint8Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Uint16:
		x := array.NewUint16Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Uint32:
		x := array.NewUint32Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Uint64:
		x := array.NewUint64Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Int8:
		x := array.NewInt8Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Int16:
		x := array.NewInt16Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Int32:
		x := array.NewInt32Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Int64:
		x := array.NewInt64Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Float32:
		x := array.NewFloat32Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.PrimitiveTypes.Float64:
		x := array.NewFloat64Data(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	case arrow.BinaryTypes.String:
		x := array.NewStringData(da)
		for i := range msk {
			msk[i] = msk[i] || x.IsNull(i)
		}
	default:
		panic("unknown type")
	}
}

func reduce(rec array.Record, bld *array.RecordBuilder, inds []int) {

	for j := range rec.Columns() {

		col := rec.Column(j)
		da := col.Data()

		switch col.DataType() {
		case arrow.PrimitiveTypes.Uint8:
			x := array.NewUint8Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Uint8Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Uint8Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Uint16:
			x := array.NewUint16Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Uint16Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Uint16Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Uint32:
			x := array.NewUint32Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Uint32Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Uint32Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Uint64:
			x := array.NewUint64Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Uint64Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Uint64Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Int8:
			x := array.NewInt8Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Int8Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Int8Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Int16:
			x := array.NewInt16Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Int16Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Int16Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Int32:
			x := array.NewInt32Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Int32Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Int32Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Int64:
			x := array.NewInt64Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Int64Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Int64Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Float32:
			x := array.NewFloat32Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Float32Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Float32Builder).Append(x.Value(i))
				}
			}
		case arrow.PrimitiveTypes.Float64:
			x := array.NewFloat64Data(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.Float64Builder).AppendNull()
				} else {
					bld.Field(j).(*array.Float64Builder).Append(x.Value(i))
				}
			}
		case arrow.BinaryTypes.String:
			x := array.NewStringData(da)
			for _, i := range inds {
				if x.IsNull(i) {
					bld.Field(j).(*array.StringBuilder).AppendNull()
				} else {
					bld.Field(j).(*array.StringBuilder).Append(x.Value(i))
				}
			}
		default:
			panic("unknown type")
		}
	}
}

// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"strconv"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

func (rdr *reader) readRow(recs []string) {

	for i, str := range recs {

		switch rdr.schema.Field(i).Type.(type) {
		case *arrow.Uint8Type:
			v, b := rdr.readUint8(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Uint8Builder).Append(v)
			}
		case *arrow.Uint16Type:
			v, b := rdr.readUint16(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Uint16Builder).Append(v)
			}
		case *arrow.Uint32Type:
			v, b := rdr.readUint32(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Uint32Builder).Append(v)
			}
		case *arrow.Uint64Type:
			v, b := rdr.readUint64(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Uint64Builder).Append(v)
			}
		case *arrow.Int8Type:
			v, b := rdr.readInt8(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Int8Builder).Append(v)
			}
		case *arrow.Int16Type:
			v, b := rdr.readInt16(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Int16Builder).Append(v)
			}
		case *arrow.Int32Type:
			v, b := rdr.readInt32(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Int32Builder).Append(v)
			}
		case *arrow.Int64Type:
			v, b := rdr.readInt64(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Int64Builder).Append(v)
			}
		case *arrow.Float32Type:
			v, b := rdr.readFloat32(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Float32Builder).Append(v)
			}
		case *arrow.Float64Type:
			v, b := rdr.readFloat64(str)
			if b {
				rdr.bld.Field(i).AppendNull()
			} else {
				rdr.bld.Field(i).(*array.Float64Builder).Append(v)
			}
		case *arrow.StringType:
			// No way to define nulls now on strings
			// r.bld.Field(i).AppendNull()
			rdr.bld.Field(i).(*array.StringBuilder).Append(str)
		case *arrow.BooleanType:
			var v bool
			switch str {
			case "false", "False", "0":
				v = false
			case "true", "True", "1":
				v = true
			}
			rdr.bld.Field(i).(*array.BooleanBuilder).Append(v)
		default:
			panic("unknown type")
		}
	}
}

func (r *reader) readUint8(str string) (uint8, bool) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, true
	}
	return uint8(v), false
}

func (r *reader) readUint16(str string) (uint16, bool) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, true
	}
	return uint16(v), false
}

func (r *reader) readUint32(str string) (uint32, bool) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, true
	}
	return uint32(v), false
}

func (r *reader) readUint64(str string) (uint64, bool) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, true
	}
	return uint64(v), false
}

func (r *reader) readInt8(str string) (int8, bool) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, true
	}
	return int8(v), false
}

func (r *reader) readInt16(str string) (int16, bool) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, true
	}
	return int16(v), false
}

func (r *reader) readInt32(str string) (int32, bool) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, true
	}
	return int32(v), false
}

func (r *reader) readInt64(str string) (int64, bool) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, true
	}
	return int64(v), false
}

func (r *reader) readFloat32(str string) (float32, bool) {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, true
	}
	return float32(v), false
}

func (r *reader) readFloat64(str string) (float64, bool) {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, true
	}
	return float64(v), false
}

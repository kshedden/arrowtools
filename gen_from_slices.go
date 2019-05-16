// GENERATED CODE, DO NOT EDIT

package arrowtools

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

func getSchema(slices []interface{}, names []string) *arrow.Schema {

	var fields []arrow.Field

	for i, x := range slices {
		switch x.(type) {
		case []uint8:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Uint8})
		case []uint16:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Uint16})
		case []uint32:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Uint32})
		case []uint64:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Uint64})
		case []int8:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Int8})
		case []int16:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Int16})
		case []int32:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Int32})
		case []int64:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Int64})
		case []float32:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Float32})
		case []float64:
			fields = append(fields, arrow.Field{Name: names[i], Type: arrow.PrimitiveTypes.Float64})
		default:
			panic("unkown type")
		}
	}

	return arrow.NewSchema(fields, nil)
}

func TableFromSlices(slices []interface{}, names []string) array.Table {

	schema := getSchema(slices, names)
	mem := memory.DefaultAllocator
	bld := array.NewRecordBuilder(mem, schema)

	for i, x := range slices {
		switch x := x.(type) {
		case []uint8:
			bld.Field(i).(*array.Uint8Builder).AppendValues(x, nil)
		case []uint16:
			bld.Field(i).(*array.Uint16Builder).AppendValues(x, nil)
		case []uint32:
			bld.Field(i).(*array.Uint32Builder).AppendValues(x, nil)
		case []uint64:
			bld.Field(i).(*array.Uint64Builder).AppendValues(x, nil)
		case []int8:
			bld.Field(i).(*array.Int8Builder).AppendValues(x, nil)
		case []int16:
			bld.Field(i).(*array.Int16Builder).AppendValues(x, nil)
		case []int32:
			bld.Field(i).(*array.Int32Builder).AppendValues(x, nil)
		case []int64:
			bld.Field(i).(*array.Int64Builder).AppendValues(x, nil)
		case []float32:
			bld.Field(i).(*array.Float32Builder).AppendValues(x, nil)
		case []float64:
			bld.Field(i).(*array.Float64Builder).AppendValues(x, nil)
		default:
			panic("unkown type\n")
		}
	}

	rec1 := bld.NewRecord()
	tbl := array.NewTableFromRecords(schema, []array.Record{rec1})

	return tbl
}

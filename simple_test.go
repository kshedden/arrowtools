package arrowtools

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

const (
	raw1 = `## a simple set of data: int64;float64;string
i64;f64;str
0;0;str-0
1;;str-1
2;2;str-2
3;3;str-3
;4;str-4
5;5;str-5
6;6;str-6
7;7;str-7
8;8;str-8
9;9;str-9
`

	// Differs from raw1 only in row 4
	raw2 = `## a simple set of data: int64;float64;string
i64;f64;str
0;0;str-0
1;;str-1
2;2;str-2
3;2;str-3
;4;str-4
5;5;str-5
6;6;str-6
7;7;str-7
8;8;str-8
9;9;str-9
`

	// Data are the same as raw1, columns are selected/reordered
	raw3 = `## a simple set of data: string;float64
str;f64
str-0;0
str-1;
str-2;2
str-3;3
str-4;4
str-5;5
str-6;6
str-7;7
str-8;8
str-9;9
`

	// raw1, with all rows with missing data dropped
	raw4 = `## a simple set of data: int64;float64;string
i64;f64;str
0;0;str-0
2;2;str-2
3;3;str-3
5;5;str-5
6;6;str-6
7;7;str-7
8;8;str-8
9;9;str-9
`

	raw5 = `## a simple set of data: int64;float64;string
i64;f64;str
0;0;str-0
1;;str-1
2;2;str-2
;4;str-4
5;5;str-5
6;6;str-6
8;8;str-8
9;9;str-9
`
)

func table1() array.Table {
	f := bytes.NewBufferString(raw1)
	fields := []arrow.Field{
		{Name: "i64", Type: arrow.PrimitiveTypes.Int64},
		{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
		{Name: "str", Type: arrow.BinaryTypes.String},
	}
	return ReadCSV(f, fields, WithComma(';'), WithComment('#'), WithHeader())
}

func table2() array.Table {
	f := bytes.NewBufferString(raw2)
	fields := []arrow.Field{
		{Name: "i64", Type: arrow.PrimitiveTypes.Int64},
		{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
		{Name: "str", Type: arrow.BinaryTypes.String},
	}
	return ReadCSV(f, fields, WithComma(';'), WithComment('#'), WithHeader())
}

func table3() array.Table {
	f := bytes.NewBufferString(raw3)
	fields := []arrow.Field{
		{Name: "str", Type: arrow.BinaryTypes.String},
		{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
	}
	return ReadCSV(f, fields, WithComma(';'), WithComment('#'), WithHeader())
}

func table4() array.Table {
	f := bytes.NewBufferString(raw1)
	fields := []arrow.Field{
		{Name: "str", Type: arrow.BinaryTypes.String},
		{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
	}
	return ReadCSV(f, fields, WithComma(';'), WithComment('#'), WithHeader())
}

func table5() array.Table {
	f := bytes.NewBufferString(raw1)
	fields := []arrow.Field{
		{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
		{Name: "str", Type: arrow.BinaryTypes.String},
	}
	return ReadCSV(f, fields, WithComma(';'), WithComment('#'), WithHeader())
}

func table6() array.Table {
	f := bytes.NewBufferString(raw4)
	fields := []arrow.Field{
		{Name: "i64", Type: arrow.PrimitiveTypes.Int64},
		{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
		{Name: "str", Type: arrow.BinaryTypes.String},
	}
	return ReadCSV(f, fields, WithComma(';'), WithComment('#'), WithHeader())
}

func table7() array.Table {
	f := bytes.NewBufferString(raw5)
	fields := []arrow.Field{
		{Name: "i64", Type: arrow.PrimitiveTypes.Int64},
		{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
		{Name: "str", Type: arrow.BinaryTypes.String},
	}
	return ReadCSV(f, fields, WithComma(';'), WithComment('#'), WithHeader())
}

func TestReader1(t *testing.T) {

	tbl3 := table3()
	tbl4 := table4()

	b, msg := TablesEqual(tbl3, tbl4)
	if !b {
		fmt.Printf("TestReader1: msg=%v\n", msg)
		t.Fail()
	}
}

package arrowtools

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

type reader struct {
	r io.Reader

	cr *csv.Reader

	schema *arrow.Schema

	mem memory.Allocator

	bld *array.RecordBuilder

	chunk int

	useCols []int

	// Character used as field delimiter in the CSV file
	comma rune

	// Lines beginning with this characeter in the CSV file are skipped
	comment rune

	hasHeader bool

	header []string
}

// Option provides a configuration option for the CSV reader.
type Option func(*reader)

// WithHeader indicates that the CSV file has a header row.
func WithHeader() Option {
	return func(rdr *reader) {
		rdr.hasHeader = true
	}
}

// WithComma defines a delimiter to the underlying csv reader.  If not set, the delimiter
// is a comma.
func WithComma(c rune) Option {
	return func(rdr *reader) {
		rdr.comma = c
	}
}

// WithComment provides a comment symbol to the underlying csv reader.
func WithComment(c rune) Option {
	return func(rdr *reader) {
		rdr.comment = c
	}
}

// ReadCSV is a CSV reader that produces a arrow Table from the data in a CSV source.
// If a header row is present, the provided Schema can contain only a subset of the
// columns in the file, and only those columns are retained.  If there is no
// header, then the length of the schema should match the number of columns
// in the source file.
func ReadCSV(r io.Reader, fields []arrow.Field, opts ...Option) array.Table {

	rdr := &reader{
		r:      r,
		cr:     csv.NewReader(r),
		schema: arrow.NewSchema(fields, nil),
	}

	for _, opt := range opts {
		opt(rdr)
	}

	if rdr.comma != 0 {
		rdr.cr.Comma = rdr.comma
	}
	rdr.cr.Comment = rdr.comment

	if rdr.mem == nil {
		rdr.mem = memory.DefaultAllocator
	}

	rdr.bld = array.NewRecordBuilder(rdr.mem, rdr.schema)

	if rdr.hasHeader {
		rdr.readHeader()
	}

	return rdr.Read()
}

func (rdr *reader) readHeader() {
	var err error
	rdr.header, err = rdr.cr.Read()
	if err != nil {
		panic(err)
	}

	vn := make(map[string]int)
	for k, n := range rdr.header {
		vn[n] = k
	}

	for _, f := range rdr.schema.Fields() {
		q, ok := vn[f.Name]
		if !ok {
			msg := fmt.Sprintf("Variable '%s' not found\n", f.Name)
			panic(msg)
		}
		rdr.useCols = append(rdr.useCols, q)
	}
}

func (rdr *reader) Read() array.Table {

	var recx []string
	if rdr.useCols != nil {
		recx = make([]string, len(rdr.useCols))
	}

	for {
		recs, err := rdr.cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if rdr.useCols != nil {
			for i, j := range rdr.useCols {
				recx[i] = recs[j]
			}
		} else {
			recx = recs
		}

		if len(recx) != len(rdr.schema.Fields()) {
			panic("Inconsistent number of columns\n")
		}

		rdr.readRow(recx)
	}

	nr := rdr.bld.NewRecord()
	tbl := array.NewTableFromRecords(rdr.schema, []array.Record{nr})

	return tbl
}

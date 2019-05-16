package arrowtools

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

// ReplaceColumn returns a new table in which one column is replaced with a
// given column.  The name of the given column must match the name of one
// column in the provided table.
func ReplaceColumn(tbl array.Table, col array.Column) array.Table {

	schema := tbl.Schema()
	pos := -1
	for i, fld := range schema.Fields() {
		if fld.Name == col.Name() {
			pos = i
		}
	}

	if pos == -1 {
		msg := fmt.Sprintf("ReplaceColumn: Column '%s' not found.", col.Name())
		panic(msg)
	}

	var newcols []array.Column
	for j := 0; j < int(tbl.NumCols()); j++ {
		newcols = append(newcols, *tbl.Column(j))
	}
	newcols[pos] = col

	return array.NewTable(arrow.NewSchema(schema.Fields(), nil), newcols, -1)
}

// AppendColumns returns a table that contains the columns of the provided table, and appends
// the given additional columns.  These new columns must not have the same name as any
// existing columns.
func AppendColumns(tbl array.Table, cols []array.Column) array.Table {

	schema := tbl.Schema()
	flds := schema.Fields()
	for i := range cols {
		flds = append(flds, arrow.Field{Name: cols[i].Name(), Type: cols[i].DataType()})
	}
	newschema := arrow.NewSchema(flds, nil)

	var newcols []array.Column
	for j := 0; j < int(tbl.NumCols()); j++ {
		newcols = append(newcols, *tbl.Column(j))
	}
	newcols = append(newcols, cols...)

	return array.NewTable(newschema, newcols, -1)
}

// DropColumns returns a table that contains the columns of the provided table, except for the columns
// with the given names.
func DropColumns(tbl array.Table, names ...string) array.Table {

	nm := make(map[string]bool)
	for _, na := range names {
		nm[na] = true
	}

	schema := tbl.Schema()
	var newflds []arrow.Field
	var newcols []array.Column
	for i, fld := range schema.Fields() {
		if !nm[fld.Name] {
			newflds = append(newflds, fld)
			newcols = append(newcols, *tbl.Column(i))
		}
	}
	newschema := arrow.NewSchema(newflds, nil)

	return array.NewTable(newschema, newcols, -1)
}

// SelectColumns returns a Table whose columns are a subset of the columns
// of the provided table.  The retained columns are shallow copies of
// the source columns.
func SelectColumns(tbl array.Table, keep ...string) array.Table {

	vn := make(map[string]bool)
	for _, na := range keep {
		vn[na] = true
	}

	schema := tbl.Schema()
	var newcols []array.Column
	var newfields []arrow.Field

	f := schema.Fields()
	for i, v := range f {
		if vn[v.Name] {
			newfields = append(newfields, v)
			newcols = append(newcols, *tbl.Column(i))
		}
	}

	newschema := arrow.NewSchema(newfields, nil)
	newtbl := array.NewTable(newschema, newcols, -1)

	return newtbl
}

// TableFromColumns creates an array.Table from the given columns.
func TableFromColumns(cols []array.Column) array.Table {

	var fields []arrow.Field
	for _, col := range cols {
		fields = append(fields, arrow.Field{Name: col.Name(), Type: col.DataType()})
	}

	schema := arrow.NewSchema(fields, nil)

	return array.NewTable(schema, cols, -1)
}

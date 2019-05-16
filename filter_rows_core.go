package arrowtools

import (
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

type SelectFunc func([]array.Interface, []bool)

// FilterRows selects a subset of the rows of a table.  Rows where the
// selector function is false are dropped.
func FilterRows(tbl array.Table, selector SelectFunc) array.Table {

	schema := tbl.Schema()
	mem := memory.DefaultAllocator
	bld := array.NewRecordBuilder(mem, schema)

	r := array.NewTableReader(tbl, 1000)
	var msk []bool
	var inds []int
	var recs []array.Record

	for r.Next() {

		rec := r.Record()
		col := rec.Column(0)
		da := col.Data()

		if cap(msk) < da.Len() {
			msk = make([]bool, da.Len())
		} else {
			msk = msk[0:da.Len()]
			for i := range msk {
				msk[i] = false
			}
		}

		selector(rec.Columns(), msk)

		inds = inds[0:0]
		for j, b := range msk {
			if !b {
				inds = append(inds, j)
			}
		}

		reduce(rec, bld, inds)

		recx := bld.NewRecord()
		recs = append(recs, recx)
	}

	return array.NewTableFromRecords(schema, recs)
}

// DropNA drops all rows from the table that contain at least one null value.
func DropNA(tbl array.Table) array.Table {

	schema := tbl.Schema()
	mem := memory.DefaultAllocator
	bld := array.NewRecordBuilder(mem, schema)

	r := array.NewTableReader(tbl, 1000)
	var msk []bool
	var inds []int
	var recs []array.Record

	for r.Next() {

		rec := r.Record()
		for j := range rec.Columns() {

			col := rec.Column(j)
			da := col.Data()

			// Set the size of the null mask
			if j == 0 {
				if cap(msk) < da.Len() {
					msk = make([]bool, da.Len())
				} else {
					msk = msk[0:da.Len()]
					for i := range msk {
						msk[i] = false
					}
				}
			}

			getnull(col, msk, da)
		}

		inds = inds[0:0]
		for j, b := range msk {
			if !b {
				inds = append(inds, j)
			}
		}

		reduce(rec, bld, inds)

		recx := bld.NewRecord()
		recs = append(recs, recx)
	}

	tblx := array.NewTableFromRecords(schema, recs)

	return tblx
}

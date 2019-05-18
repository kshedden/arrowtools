package arrowtools

import (
	"github.com/apache/arrow/go/arrow/array"
)

type RecordHelper struct {
	rec array.Record

	curpos int

	vnames map[string]int
}

// NewRecordHelper returns a RecordHelper that can obtain data slices
// from the given Record.
func NewRecordHelper(rec array.Record) *RecordHelper {

	rh := &RecordHelper{
		rec: rec,
	}

	vnames := make(map[string]int)
	for k := range rec.Columns() {
		vnames[rec.ColumnName(k)] = k
	}
	rh.vnames = vnames

	return rh
}

// SetPos sets the column position from which the next SliceFromRecord
// call will return data.
func (rh *RecordHelper) SetPos(pos int) {
	rh.curpos = pos
}

// SetName sets the column name from which the next SliceFromRecord
// call will return data.
func (rh *RecordHelper) SetName(name string) {
	rh.curpos = rh.vnames[name]
}

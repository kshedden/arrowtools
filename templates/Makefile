all: columns equality filter_rows csvreader table

.PHONY: columns equality from_slices filter_rows csvreader

columns:
	go run main.go -template=columns.tmpl
	go run main.go -template=columns_test.tmpl

equality:
	go run main.go -template=equality.tmpl
	go run main.go -template=equality_test.tmpl

filter_rows:
	go run main.go -template=filter_rows.tmpl

table:
	go run main.go -template=table.tmpl
	go run main.go -template=table_test.tmpl

csvreader:
	go run main.go -template=csvreader.tmpl
	go run main.go -template=csvreader_test.tmpl

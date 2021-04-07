package server

var PgTypeMap = map[string]int{
	"bool":  16,
	"bytea": 17,
	"char":  18,
	"int8":  20,
	"int2":  21,
	"int4":  23,
	"text":  25,
	"oid":   26,
}

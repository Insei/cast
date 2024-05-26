[![codecov](https://codecov.io/github/Insei/cast/graph/badge.svg?token=KPPYANT85P)](https://codecov.io/github/Insei/cast)
[![build](https://github.com/insei/cast/actions/workflows/go.yml/badge.svg)](https://github.com/Insei/cast/actions/workflows/go.yml)
[![Goreport](https://goreportcard.com/badge/github.com/insei/cast)](https://goreportcard.com/report/github.com/insei/cast)
[![GoDoc](https://godoc.org/github.com/insei/cast?status.svg)](https://godoc.org/github.com/insei/cast)
# Cast
Cast is string casting library, support cast string to generic type, reflect type and value type.

## Supported types
```
(*)int
(*)int8
(*)int16
(*)int32
(*)int64
(*)uint
(*)uint8
(*)uint16
(*)uint32
(*)uint64
(*)float32
(*)float64
(*)bool
(*)string
(*)time.Time
(*)uuid.UUID (Google)
```
## Examples
`To[int]` example.
```go
valInt, err := cast.To[int]("56")
// or valPtrInt, err := To[*int]("56")
if err != nil {
	panic(err)
}
```
`To[time.Time]` example (supports strings only in time.RFC3339).
```go
valTime, err := cast.To[time.Time]("2024-05-22T11:36:57+03:00")
// or valPtrTime, err := To[*time.Time]("2024-05-22T11:36:57+03:00")
if err != nil {
	panic(err)
}
```
`ToReflect(string, reflect.Type)` example.
```go
timeType := reflect.TypeOf(time.Time{})
valTime, err := cast.ToReflect("2024-05-22T11:36:57+03:00", timeType)
// or valPtrTime, err := ToReflect("2024-05-22T11:36:57+03:00", reflect.PointerTo(timeType))
if err != nil {
	panic(err)
}
```
`ToFrom(string, any)` example.
```go
date := time.Time{}
err := cast.ToFrom("2024-05-22T11:36:57+03:00", &date)
if err != nil {
	panic(err)
}
```

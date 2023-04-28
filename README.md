# flagx

## Usage

Assume that you have the following `Date` type.

```go
type Date struct {
  Year, Month, Day int
}
```

Configure `Parse` and `String` function to use flags of type `Date`.

```go
var dateFlag = flagx.Config[Date]{
  Parse: func(s string) (Date, error) {
    tm, err := time.Parse("2006-01-02", s)
    if err != nil {
      return Date{}, err
    }
    year, month, day := tm.Date()
    return Date{Year: year, Month: int(month), Day: day}, nil
  },
  String: func(d Date) string {
    return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
  },
}
```

Now you can use `dateFlag.Value()` to create `flag.Value`, which can be used for `flag.Var()`.

```go
var date Date

func init() {
  flag.Var(dateFlag.Value(&date, Date{Year: 2006, Month: 1, Day: 2}), "date", "usage")
}
```

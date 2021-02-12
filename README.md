# reldatelib

A simple Go library to compute relative dates from string similar to the unix tool `date`

## Examples
`datelib.Parse("now")`

Fri 12 Feb 2021 11:00:42 AM EST

`datelib.Parse("yesterday")`

Thu 11 Feb 2021 11:01:23 AM EST

`datelib.Parse("2 weeks ago")`

Fri 29 Jan 2021 11:01:40 AM EST

`datelib.Parse("last wednesday")`

Wed 10 Feb 2021 12:00:00 AM EST

## Possible options
`"LAST", "THIS", "NEXT", "FIRST", "THIRD", "FOURTH", "FIFTH", "SIXTH", "SEVENTH", "EIGHTH", "NINTH", "TENTH", "ELEVENTH", "TWELFTH"`

`"SEC", "SECOND", "MIN", "MINUTE", "HOUR", "DAY", "WEEK", "FORTNIGHT", "MONTH", "YEAR"`

`"AGO", "HENCE"`

`"NOW", "TODAY", "YESTERDAY", "TOMORROW"`

`"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"`

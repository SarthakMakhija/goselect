# [![Actions Status](https://github.com/SarthakMakhija/goselect/workflows/GoSelectCI/badge.svg)](https://github.com/SarthakMakhija/goselect/actions) goselect
SQL like 'select' interface for files

# Some examples
- select fName from /home/apps
- select fSize, fName from /home/apps where fSize > 10MB

# Initial thoughts

The idea is to provide support for `selecting` files based on various conditions. At this stage some of the features that are planned include:
- Support for `where` clause
- Support for searching text within a file
- Support for some shortcuts for common files types like images, videos, text:
  - select * from /home/apps where fileType = 'image' and fileSize > 10MB
  - select * from /home/apps where fileType = 'text' and textContains = 'get('
- Support for projections
  - [X] projections with attribute name: name, size
  - [X] projections with scalar functions: contains, lower
  - [ ] projections with aggregate functions: min, max
  - [ ] projections with expression: 1 + 2
- Support for `order by` clause
  - [X] order by with positions: order by 1
  - [X] order by with descending order: order by 1 desc
  - [X] order by with optional ascending order: order by 1 asc
- Support for `limit` clause
  - [X] limit clause with a value: limit 10
- Support for `aggregation functions`
  - [ ] min
  - [ ] max
  - [ ] avg
  - [ ] sum
  - [ ] count
  - [ ] median
- Support for various `scalar functions`
  - [X] lower
  - [X] upper
  - [X] title
  - [X] base64
  - [X] length
  - [X] lTrim
  - [X] rTrim
  - [X] trim
  - [X] now
  - [X] date
  - [X] day
  - [X] month
  - [X] year
  - [X] dayOfWeek
  - [X] working directory (wd)
  - [X] concat
  - [X] concat with separator (concatWs)
  - [X] contains
  - [ ] substr
  - [ ] replace
  - [ ] replaceAll
  - [ ] formatSize
- Support for formatting the results
  - [X] Json formatter
  - [X] Html formatter
- Support for exporting the formatted result
  - [X] Console
  - [X] File
- Design consideration for searching files in a directory with a huge number of files

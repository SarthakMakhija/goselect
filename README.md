# goselect
SQL like interface for files

# Some examples
- goselect fName from /home/apps
- goselect fSize, fName from /home/apps where fSize > 10MB

# Initial thoughts

The idea is to provide support for `selecting` files based on various conditions. At this stage some of the features that are planned include:
- Support for `where` clause
- Support for projections
- Support for `order by`
- Support for `aggregations`
- Support for custom filter
- Support for custom map-reduce function
- Support for searching text within a file
- Support for some shortcuts for common files types like images, videos, text:
  - goselect * from /home/apps where fileType = 'image' and fileSize > 10MB
  - goselect * from /home/apps where fileType = 'text' and textContains = 'get('
- Design consideration for searching files in a directory with a huge number of files
- [Not sure] Support for caching the query results
- [Not sure] Support for union and intersection


# sqlcgetters

[![godoc](https://pkg.go.dev/badge/github.com/willabides/sqlcgetters.svg)](https://pkg.go.dev/github.com/willabides/sqlcgetters)

`sqlcgetters` command line tool and go module that generates getters for stucts that were generated by `sqlc`.

<!--- everything between the next line and the "end usage output" comment is generated by script/generate-readme --->
<!--- start usage output --->
```
Usage: sqlcgetters <path>

sqlcgetters generates getters for structs generated by sqlc.

Arguments:
  <path>    Path to sqlc output directory.

Flags:
  -h, --help       Show context-sensitive help.
      --version    Output the version and exit.
```
<!--- end usage output --->

## Why do I need this?

`sqlc` generates structs with exported fields, so why would you want getters? It's all about interfaces. Take these 
queries for example: 

```postgresql
-- name: BooksByTitleYear :many
SELECT book_id, title, authors.name AS author_name, isbn
FROM books
  LEFT JOIN authors ON books.author_id = authors.author_id
WHERE title = $1 AND year = $2;

-- name: BooksByTags :many
SELECT book_id, title, authors.name AS author_name, isbn, tags
FROM books
  LEFT JOIN authors ON books.author_id = authors.author_id
WHERE tags && $1::varchar[];
```

`sqlc` will generate two queries with these structs as their results:

```go
type BooksByTitleYearRow struct {
    BookID       int32
    Title        string
    AuthorName   sql.NullString
    Isbn         string
}

type BooksByTagsRow struct {
    BookID       int32
    Title        string
    AuthorName   sql.NullString
    Isbn         string
    Tags         []string
}
```

You have two different structs with the same fields. Writing a function that handles either of them is difficult. 
You need to accept `interface{}` and the type switch to handle the different structs. But if you had a Get*() 
function for each field, your function can accept an interface with the fields you need.

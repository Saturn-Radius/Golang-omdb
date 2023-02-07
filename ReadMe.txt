### Assessment:

Please write a small Golang server to search movies from http://www.omdbapi.com/ and
get a movie detail.


### Running

>go mod tidy
-Run
>go run .

-- in browser:
http://localhost:8080/search?s=Batman&page2

http://localhost:8080/detail/tt4853102

-Test

>cd handler
>go test
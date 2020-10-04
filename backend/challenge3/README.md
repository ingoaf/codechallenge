# Simple and extended search

## Description
Execute a simple search. You may recognize that the results can contain many companies, that are similar to the searched name.
To narrow the search result to the company you are looking for, you can use an extended search.
These can contain more informated queries.

## Steps
Use atreugo to implement an API with two routes.
The first should execute the simple search, using the name from the path as query.
The second should execute the extended search. The query comes from the body of the POST request.

1. Get familiar with atreugo. You can find the documentation [here](https://pkg.go.dev/github.com/savsgio/atreugo/v11)
1. gocloaksession can be used as a middleware for atreugo. Get familiar with it [here]{https://github.com/Clarilab/gocloaksession}
1. The functions, that are called by atreugo should include these steps:
    1. Read the parameter or body from the request.
    1. Use resty to call the search-API. The documentation is [here](https://godoc.org/github.com/go-resty/resty)
        1. The simple search is a GET request, and no body is needed
        1. The extended search uses the `SearchRequest`as body for the POST request
    1. Check the response for failures
    1. Return the result from the request as `SearchResponse`

## For the supervisor
1. Provide values for keycloak authentication (clientID, clientSecret, realm, url)
1. Provide url to API
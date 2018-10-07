# Medicinal products registry API

This is a webserver which exposes a convenient GraphQL endpoint for fetching data from [polish register of medicinal products][registry] which is a public data from [dane.gov.pl].

New data is fetched every 20 minutes (and this can be changed in code) from the server.

## Compilation & usage

To compile, you need to install dependencies by `dep ensure` (you need [Dep]) in project's directory. Then, all you have to do is to run `go build`.

To run the server, you need to provide the `MPR_ADDR` environment variable to tell the server on what to listen (this can be `:8080`, for example).

```bash
MPR_ADDR=localhost:80 ./app
```

## API

Below screenshot shows all endpoints and fields exposed in the API.

![Insomnia screenshot][api-screenshot]

## Contributing

...is welcome! :)

[registry]: https://dane.gov.pl/dataset/397/resource/1851
[dane.gov.pl]: https://dane.gov.pl
[Dep]: https://golang.github.io/dep/
[api-screenshot]: https://i.imgur.com/xI8BmCO.png
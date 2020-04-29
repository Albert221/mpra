# Medicinal Products Registry API

This is a webserver which exposes a convenient GraphQL endpoint for fetching data from [polish register of medicinal products][registry] which is a public data from [dane.gov.pl].

![GraphQL Playground screenshot][api-screenshot]

## Compilation & usage

You can run this project using Docker or Go (1.14 or newer).

##### Docker

```bash
git clone https://github.com/Albert221/medicinal-products-registry-api.git mpra && cd mpra
docker build --tag mpra:latest .
docker run --publish 8080:8080 mpra:latest
```


##### Go

```bash
git clone https://github.com/Albert221/medicinal-products-registry-api.git mpra && cd mpra
go get github.com/markbates/pkger/cmd/pkger
pkger
go build -o mpra .
```

### Usage

```
$ ./mpra -help
Usage of ./mpra:
  -dataset string
        path to the file which the dataset will be cached to (default "dataset.xml")
  -host string
        hostname to listen on
  -port string
        port to listen on (default "8080")
  -refresh duration
        period every which the dataset will be refreshed (default 1h0m0s)
```

## Contributing

...is welcome! :)

[registry]: https://dane.gov.pl/dataset/397/resource/1851
[dane.gov.pl]: https://dane.gov.pl
[api-screenshot]: https://i.imgur.com/gfoPieX.png

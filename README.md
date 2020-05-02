# Medicinal Products Registry API

[![FOSSA Status][fossa-badge]][fossa-link]

This is a webserver that exposes a convenient GraphQL endpoint for fetching data from [polish register of medicinal products][registry] which is a public data from [dane.gov.pl].

![GraphQL Playground screenshot][api-screenshot]
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FAlbert221%2Fmpra.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FAlbert221%2Fmpra?ref=badge_shield)

## Compilation & usage

You can run this project using Docker or Go (1.14 or newer).

##### Docker

```bash
git clone https://github.com/Albert221/mpra.git mpra && cd mpra
docker build --tag mpra:latest .
docker run --publish 8080:8080 mpra:latest
```


##### Go

```bash
git clone https://github.com/Albert221/mpra.git mpra && cd mpra
go get github.com/markbates/pkger/cmd/pkger
pkger
go build -o mpra .
./mpra
```

You can set few arguments using the command line, use `./mpra -help` for help:

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
[fossa-badge]: https://app.fossa.io/api/projects/git%2Bgithub.com%2FAlbert221%2Fmpra.svg?type=shield
[fossa-link]: https://app.fossa.io/projects/git%2Bgithub.com%2FAlbert221%2Fmpra?ref=badge_shield


## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FAlbert221%2Fmpra.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FAlbert221%2Fmpra?ref=badge_large)

# benchlinksdk

This package provides a client SDK binding to the benchling API.

```sh
go install -x github.com/cosnicolaou/oapi-tool@latest
oapi-tool download --output='benchling.yaml' 'https://benchling.com/api/v2/openapi.yaml'
oapi-tool format --output='formated.yaml' -validate=false benchling.yaml
oapi-tool transform --output='transformed.yaml' --config=transformations.yaml formated.yaml
```

Note that github.com/deepmap/oapi-codegen contains a bug that prevents it from
processing the benchling spec. Instead, use the forked version github.com/cosnicolaou/oapi-codegen
until the fix is merged upstream, eg:

```
(h=$(pwd) ; cd ../oapi-codegen/cmd/oapi-codegen; go run . --package=benchlingsdk $h/transformed.yaml ) > benchlingsdk.go
```

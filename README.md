# 12factors-workshop

Este repositorio é um código de exemplo criado com carinho e propositalmente tem muitas falhas de design pra que a gente possa refator-lo, usando as boas praticas citadas no [12factor.net](https://12factor.net/).

## Funcionalidades

WIP...

## TODO

- [x] usar um txt pra armazenar coisas
- [ ] deixar algumas features funcionando apenas em 1 branch (ou fazemos forks)
- [x] não usar go mod ou dep pra fazer gestão da dependencia - nem usar makefiles
- [x] subir algo que use uma porta fixa no codigo (dae não roda concorrente - muito menos suporta rodar em portas diferentes)
- [x] quando morre com erro, nem sempre salva o banco de dados
- [ ] sempre precisa de um config que não ta salvo no repo

## Usage

> **IMPORTANTE:** Lembre de fazer o clone desse repositorio no diretório `~/go/src/github.com/golang-poa/12factors-workshop`

### set up 

Go version: `go1.13.6 darwin/amd64`.

Instale:

```
go get -u github.com/gin-gonic/gin
go get -u github.com/cosmtrek/air

```

crie o arquivo de banco de dados:

```bash
echo '[]' > /tmp/db.json
```
> inicie a app apenas depois de criar esse arquivo! senão vai ganhar um erro tipo esse:
> 
> ```bash
> 12factors-workshop on  master [?] via 🐹 v1.13.6
> ➜ air
> 
>   __    _   ___
>  / /\  | | | |_)
> /_/--\ |_| |_| \_ // live reload for Go apps [v1.12.0]
> 
> watching .
> !exclude tmp
> building...
> running...
> [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
> 
> [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
>  - using env:	export GIN_MODE=release
>  - using code:	gin.SetMode(gin.ReleaseMode)
> 
> [GIN-debug] GET    /                         --> main.init.0.func1 (3 handlers)
> [GIN-debug] GET    /api/v1/products/         --> main.init.0.func2 (3 handlers)
> [GIN-debug] POST   /api/v1/products/:name    --> main.init.0.func3 (3 handlers)
> carregando db file
> panic: could not read the database file: open /tmp/db.json: no such file or directory
> 
> goroutine 1 [running]:
> main.init.1()
> 	/Users/seba/projetos/pessoal/12factors-workshop/db.go:26 +0x44
> ```

### run

```bash
<WORKSHOP_DB="path/to/file.json"> go run *.go
```
> or run with the code live reload:
> 
> ```bash
> air
> ```

### API Usage

* `GET /api/v1/products/`
* `POST /api/v1/products/:name`


### CONFIG

| CONFIG  | DESCRIPTION | DEFAULT |
| ------------- | ------------- | ------------- |
| `WORKSHOP_DB`  |  Arquivo do banco de dados  |  `/tmp/db.json` |
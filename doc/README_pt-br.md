# go_houdini
<a href="https://github.com/luizfiuzaa/go_houdini/blob/main/README.md">[english]</a>

Uma ferramenta de linha de comando para criar arquivos `.go` com estrutura definida — pacote, imports e funções — sem precisar escrever boilerplate manualmente.

![Gopher wizard](./assets/gopher.png)

## Funcionalidades

- Prompts interativos para scaffolding de arquivos Go
- Nome do pacote configurável (padrão `main`)
- Imports e funções gerados a partir de entrada separada por espaços
- Impede a sobrescrita de arquivos existentes
- Zero dependências externas

## Requisitos

- [Go 1.23+](https://golang.org/dl/)

## Instalação

Clone o repositório e compile o binário:

```sh
git clone https://github.com/luizfiuzaa/go_houdini.git
cd go_houdini
go build -o houdini .
```

Para usar o `houdini` de qualquer diretório, mova o binário para uma pasta no seu `PATH`:

```sh
# Linux / macOS
mv houdini /usr/local/bin/

# Windows — mova o houdini.exe para uma pasta já no PATH, ex: C:\tools\
```

## Como usar

### Criar um novo arquivo Go

```sh
houdini create
```

Você também pode passar o nome do arquivo diretamente para pular o prompt de nome:

```sh
houdini create hello_world
```

**Exemplo de sessão:**

```
📦 Package name (default 'main'): main
📚 Imports (space separated): fmt time
🔧 Functions (space separated): initApp startServer

✅ File created successfully! Enjoy!
```

Isso gera o arquivo `hello_world.go`:

```go
package main

import(
  "fmt"
  "time"
)

func initApp() {}

func startServer() {}
```

### Ajuda

```sh
houdini --help
houdini -h
```

## Contribuindo

Pull requests são bem-vindos. Para mudanças significativas, abra uma issue antes para discutir o que deseja alterar.

## Licença

[MIT](../LICENSE)

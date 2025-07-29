# go_houdini
Uma maneira rápida de criar um arquivo `.go` com estrutura básica, imports e funções.

<img src="https://github.com/luizfiuzaa/go_houdini/blob/main/doc/assets/gopher.png" width="150">

## Funcionalidades

- Crie rapidamente um novo arquivo Go com:
  - Pergunta pelo nome do arquivo
  - Imports (separados por espaço)
  - Funções (separadas por espaço)
- Evita sobrescrever arquivos existentes
- Adiciona um comentário com informações do projeto

## Como usar

1. **Compile o script:**
   ```sh
   go build -o go_houdini.exe houdini.go
   ```

2. **Execute o script:**
   ```sh
   ./go_houdini.exe
   ```
   Siga as instruções para informar:
   - O nome do arquivo (não precisa adicionar `.go`)
   - Pacotes para importar (separados por espaço, ou pressione Enter para pular)
   - Nomes das funções (separados por espaço, ou pressione Enter para pular)

3. O script irá gerar um arquivo Go com a estrutura especificada.

## Como adicionar à variável de ambiente (Windows)

Para usar o `go_houdini` em qualquer terminal:

1. Copie o `go_houdini.exe` para uma pasta, por exemplo: `C:\tools\go_houdini`.
2. Adicione essa pasta ao seu `PATH` do sistema:
   - Pressione `Win + R`, digite `sysdm.cpl` e pressione Enter.
   - Vá até a aba **Avançado** e clique em **Variáveis de Ambiente**.
   - Em **Variáveis do sistema**, selecione `Path` e clique em **Editar**.
   - Clique em **Novo** e adicione `C:\tools\go_houdini`.
   - Clique em **OK** para salvar.

3. Abra um novo terminal e execute:
   ```sh
   go_houdini
   ```

Agora você pode usar o `go_houdini` de qualquer diretório.

---
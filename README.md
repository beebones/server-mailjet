# Server MailJet

Serviço backend para enviar PDF de conclusão de curso.

## Configuração

É necessário definir as variáveis de ambiente com a sua chave pública e secreta do MailJet, utilizando os nomes abaixo:
```
MJ_APIKEY_PUBLIC
MJ_APIKEY_PRIVATE
```

Tutorial ensinando a adicionar variável de ambiente no [Windows 10](https://winaero.com/blog/create-environment-variable-windows-10/) e no [Linux](https://www.todoespacoonline.com/w/2015/07/variaveis-de-ambiente-no-linux/).


## Como rodar a aplicação usando Go?

```
go run .
```

## Como compilar a aplicação?

```
go build .
```
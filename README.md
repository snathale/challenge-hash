# Hash Teste Back-end

**Cadidata**: Nathale Silva

## Requisitos

* [Docker-compose](https://docs.docker.com/compose/install/)
* [GitHub Token](https://github.com/settings/tokens)

## Arquitetura do projeto

Na construção dos serviços abaixo descritos, utilizou-se uma arquitetura baseada em camadas, de modo que possibilite a
manutenção, bem como a inserção de novos recursos a aplicação.

Como requisitos foram desenvolvidos microserviços, utilizando gRPC como padrão de comunicação entre estes. Um microserviço foi desenvolvido em Golang, sendo este o server
gRPC, outro microserviço em NodeJS e outro em PHP, como tecnologia de banco de dados foi utilizado o [ArangoDB](https://www.arangodb.com).

**product-list**: Api RESTFUL escrita em PHP responsável por receber requisições para retornar uma lista de produtos com discontos.

**product-list-node**: Api RESTFUL escrita em NodeJS responsável por receber requisições para retornar uma lista de produtos com discontos.

**calculator**: Microsserviço responsável por servir o disconto calculado para um determinado usuário e produto através
de gRPC.

**arangoDB**: Serviço de banco de dados noSql responsável por armazenar informações dos produtos e usuários.

Observação:
* Neste projeto há duas interfaces RESTFUL, uma vez que o build do serviço **product-list** escrito em PHP, mostrou-se 
  bastante lento, logo uma segunda opção de interface RESTFUL foi implementada buscando otimizar este processo.

## Iniciando a aplicação

### Product-list

Após clonar o projeto em sua máquina, com os requisitos já atendidos, configure o token de
acesso do github no arquivo ```composer.json``` que está no dentro do diretório ```/product-list```, alterando o
seguinte valor:

```
{
  "github-oauth": {
    "github.com": "XXXXXX"
  }
}
```
Em seguida estando na raiz do projeto execute arquivo bash com o seguinte comando:

```shell script
    /bin/bash init.sh
```

Esse comando subirá os quatro serviços, o servidor gRPC será executado na porta ```3000```, a api de listagem de produtos (ngnix)
será executado na porta ```8080```, php-fm e o banco de dados estará disponível na porta ```8529``` e também será inicializado
com registros dummy.

Observações:

* Para o serviço de cálculo de desconto ```/calculator``` as configurações de host, porta e conexão com o banco de dados
  são informadas pelo arquivo de configuração **config.json** localizado dentro do projeto.
* Para a API RESTFUL ```/product-list``` as configurações de host, porta e conexão com o banco de dados são informadas
  pelo arquivo de configuração **.env.example** localizado dentro do projeto.
* Uma vez executado o comando que faz o build dos dois serviços, caso deseje parar os serviços e executá-los novamente
  não se faz mais necessário utilizar o mesmo comando, basta utilizar: ```docker-compose up ```

#### Rota

```GET /product```

**Requisição**

Parâmetros (*Via Header*)

* **X-USER-ID**: Id de usuário

Observação: Este Header não é obrigatório

Exemplo:

```
curl -H "X-USER-ID: 1" -X GET 'localhost:8080/product'
```

##### Resposta

+ **Success** 200

```json
{
  "status": "success",
  "data": [
    {
      "id": "1",
      "price_in_cents": 4750,
      "title": "Iphone 12 Pro",
      "description": "The newest model phone pro",
      "discount": {
        "percentage": 0.5,
        "value_in_cents": 4750
      }
    }
  ]
} 
```

### Product-list-node

Após clonar o projeto em sua máquina, com os requisitos já atendidos e estando na raiz do projeto execute arquivo bash com o seguinte comando:

```shell script
    /bin/bash init-node.sh
```

Esse comando subirá os três serviços, o servidor gRPC será executado na porta ```3000```, a api de listagem de produtos
será executado na porta ```3333```, o banco de dados estará disponível na porta ```8529``` e também será inicializado
com registros dummy.

Observações:

* Para o serviço de cálculo de desconto ```/calculator``` as configurações de host, porta e conexão com o banco de dados
  são informadas pelo arquivo de configuração **config.json** localizado dentro do projeto.
* Para a API RESTFUL ```/product-list-node``` as configurações de host, porta e conexão com o banco de dados e conexão 
  com o servico ```calculator```são informadas pelo arquivo de configuração **.env.example** localizado dentro do projeto.

#### Rota

```GET /product```

**Requisição**

Parâmetros (*Via Header*)

* **X-USER-ID**: Id de usuário

Observação: Este Header não é obrigatório

Exemplo:

```
curl -H "X-USER-ID: 1" -X GET 'localhost:3333/product'
```

##### Resposta

+ **Success** 200

```json
{
  "status": "success",
  "data": [
    {
      "id": "1",
      "price_in_cents": 4750,
      "title": "Iphone 12 Pro",
      "description": "The newest model phone pro",
      "discount": {
        "percentage": 0.5,
        "value_in_cents": 4750
      }
    }
  ]
} 
```

#### Executando testes

**Requisitos**

* [Golang](https://golang.org/doc/install) 1.15

##### Comando

Para o serviço de cálculo de desconto, é necessário entrar no respositório ```/calculator``` e executar o comando:

``` shell script
make run-tests
```

Observação:
* Estes comando também executará um teste de integração com o a api desenvolvida em NodeJS, logo a stack docker 
  correspondente deve estar em pé, para isto, basta seguir o processo de inicialização desta stack, descrita a cima na 
  sessão de inicialização do servico **product-list-node** 


Para o serviço de listagem de produtos com descontos desenvolvido em PHP, é necessário entrar no repositório ```/product-list``` e executar
o seguinte comando:

```shell script
make run-tests
```

Observação:
* Estes comando executará testes de integração com o a api desenvolvida em PHP, logo a stack docker
  correspondente deve estar em pé, para isto, basta seguir o processo de inicialização desta stack, descrita a cima na
  sessão de inicialização do servico **product-list** 
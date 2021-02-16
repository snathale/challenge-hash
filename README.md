# Challenge-Hash

**Recrutadora**: July Demenjon

**Cadidata**: Nathale Silva

**Contato**: silva.nathale@gmail.com

#### Requisitos

* [Docker-compose](https://docs.docker.com/compose/install/)

#### Arquitetura do projeto

Na construção dos serviços abaixo descritos, utilizou-se uma arquitetura baseada em camadas, de modo que possibilite a manutenção, bem como a inserção de novos recursos a aplicação. 
Como requisitos foram desenvolvidos 2 microserviços, utilizando gRPC como padrão de comunicação entre estes. 
Um microserviço foi desenvolvido utilizando Golang e o segundo utilizando PHP, como tecnologia de banco de dados foi utilizado o [ArangoDB](https://www.arangodb.com).

**product-list**: Api RESTFUL responsável por receber requisições para retornar uma lista de produtos com discontos.

**calculator**: Microsserviço responsável por servir o disconto calculado para um determinado usuário e produto através
de gRPC.

**arangoDB**: Serviço de banco de dados noSql responsável por armazenar informações dos produtos e usuários.


#### Iniciando a aplicação

Após clonar o projeto em sua máquina, com os requisitos já instalados e estando na raiz do projeto, executar um arquivo
bash com o seguinte comando:

```shell script
    /bin/bash init.sh
```

Esse comando subirá os três serviços, o servidor gRPC será executado na porta ```3000```, a api de listagem de produtos
será executado na porta ```8080```, o banco de dados estará disponível na porta ```8529``` e também será inicializado
com registros dummy.

Observações:

* Para o serviço de cálculo de desconto ```/calculator``` as configurações de host, porta e conexão com o banco de dados
  são informadas pelo arquivo de configuração **config.json** localizado dentro do projeto.
* Para a API RESTFUL ```/product-list``` as configurações de host, porta e conexão com o banco de dados são informadas pelo arquivo de
  configuração **.env.example** localizado dentro do projeto.
* Uma vez executado o comando que faz o build dos dois serviços, caso deseje parar os serviços e 
  executá-los novamente não se faz mais necessário utilizar o mesmo comando, basta utilizar: ```docker-compose up ```

#### Rota

```GET /product```

**Requisição**

Parâmetros (*Via Header*)

* **X-USER-ID**: Id de usuário

Observação: Este Header não é obrigatório

Exemplo:

```curl -H "X-USER-ID: 1" -X GET 'localhost:8080/product'```

##### Respostas

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

```make run-tests```

Para o serviço de listagem de produtos com descontos, é necessário entrar no repositório ```/product-list``` e executar o seguinte comando:

```make run-tests```
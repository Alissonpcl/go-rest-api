# Go Rest API

Exemplo desenvolvido durante o curso [Creating Web Services with Go](https://app.pluralsight.com/library/courses/creating-web-services-go/table-of-contents) na [Pluralsight](https://app.pluralsight.com/).

Este projeto possui exemplos de:

1. Tratamento de requisições HTTP (routing, middleware, cors)
1. JSON Marshal/Unmarshall
1. SQL Database (com controle de pool e context)
1. Upload e Download de arquivos
1. Websockets
1. Templating (para criação de relatórios em HTML)

## Exemplo de requisições

**Para os comandos abaixo funcionarem é necessário que a aplicação esteja rodando e que o banco de dados tenha sido criado e populado com os script contigo no arquivo [products.sql](products.sql)**

**Websocket** <br/>
O Websocket pode ser testado diretamente em um browser utilizado a aba de Developer Tools.
No console digite o seguinte comando Javascript para iniciar uma conexão Websocket com a aplicação:

```let ws = new WebSocket("ws://localhost:5000/websocket")```

A partir deste momento o browser começará a receber retornos do servidor com uma lista de 10 produtos a cada 10 segundos.

É possível também enviar uma mensagem para o Websocket. Exemplo:

```ws.send(JSON.stringify({"name":"teste"}))```

**Demais métodos**

Endpoinp | Comando
-------- | -------
Listar todos os produtos | ```curl --location --request GET 'localhost:5000/api/products'```
Obter dados de um produto específico | ```curl --location --request GET 'localhost:5000/api/products/<id_produto>'```
Inserir produto | ```curl --location --request POST 'localhost:5000/api/products'--header 'Content-Type: application/json'--data-raw '{"manufacturer": "Alisson Lima", "sku": "t8A474iuv", "upc": "669100000000", "pricePerUnit": "369.46", "quantityOnHand": 5000, "productName": "balloon"}'```
Atualizar produto | ```curl --location --request PUT 'localhost:5000/api/products/12'--header 'Content-Type: application/json'--data-raw '{"productId": 12, "manufacturer": "Weimann, Waelchi and Kassulke", "sku": "t8A474iuv", "upc": "669100000000", "pricePerUnit": "369.46", "quantityOnHand": 5000, "productName": "balloon"}'```
Excluir produto | ```curl --location --request DELETE 'localhost:5000/api/products/191'--header 'Content-Type: application/json'--data-raw '{"manufacturer": "Alisson Lima", "sku": "t8A474iuv", "upc": "669100000000", "pricePerUnit": "369.46", "quantityOnHand": 5000, "productName": "balloon"}'```
Upload de arquivo | curl --location --request POST 'localhost:5000/api/receipts' --form 'receipt=@/Some/Path/File'
Download de arquivo | ```curl --location --request GET 'localhost:5000/api/receipts/<file_uploaded_name>'```


## Comandos disponíveis

**Executando a aplicação** <br/>
```go run main.go```

**Gerando a dist (arquivo executável)** <br/>
```go build github.com/alissonpcl/go-rest-api``` 

Será gera um arquivo *go-rest-api* (a extensão pode variar de acordo com SO).

 

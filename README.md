## Introduction 
Start project of GraphQl API with Golang and MongoDB



## Getting Started

1. Install modules
   ```sh
    $ make deps  
   ``` 
2. Build application
   ```sh
    $ make build 
   ``` 
3. Start application in development
    ```sh
    $ make run   
   ``` 
4. Start application in production
   ```sh
    $ make run-prod   
   ``` 

## Observations
**Schema-Changes lead you to re-generate**
* Very change in the graph/graphql/*.graphqls files should be finished with re-generating the GraphQL-Files by running:
 ```sh
    $ make generate  
   ``` 
**New Actions are not automatically added to the resolvers-implementation**
* New Actions inside the schema do not lead gqlgen to add them to your resolvers. You are responsible for implement them on your own (The interface containing all necessary functions can be found in `graph/generated/generated.go`).

**Models should live only inside your model-directory**

* When you define a new model inside *.graphls and re-generate your files, the schema will be generated and added to graph/model/models_gen.go.

## Features
1. Golang (Language) (https://go.dev/)
2. Graphql-Go - gqlgen (API) (https://github.com/99designs/gqlgen)
3. MongoDB (Database) (https://www.mongodb.com/pt-br)
4. Go Playground Validator (Validation) (https://github.com/go-playground/validator)

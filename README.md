# starter-pack-api

## Download the source and build

```
git clone 
cd starter-pack-api
go build main.go
```

## Launch tests

Pour tester seulement un package utiliser les commandes suivantes :

```
cd repository_with_test.go
go test -v -cover
```

Pour lancer tous les tests de tous les packages :

```
go test ./... -v -cover
```

## Launch the executable

```
./main.exe -configpath="Path to your config file"
```

## File structure

The project is divided into different sub-packages with the aim of simplifying the understanding of the code

## Updating the dependencies
- Create a branch
- Delete vendor, go.mod, go.sum
- ```go mod init starter-pack-api | go mod tidy | go mod vendor```
- Test it
- Commit and push

## Generate swagger
- use the following command ```go install github.com/swaggo/swag/cmd/swag@latest```
- go take the swag.exe from GOROOT/bin/ and put it in your current directory
- use the following command ```./swag.exe init```

## Get the swagger.json
- `Replace the needed information in the following URL
  ```http://yourhost:yourport/json/swagger_v(numéroversion).json?key=yourkey```
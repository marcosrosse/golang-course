To create a module in Golang, access the folder and run:
    ```
    go mod init [module name]
    ```

To build the project, run:
    ```
    go build
    ```

To import a module, run:
    ```
    go get [module name]
    ```

To delete a module, run:
    ```
    go mod tidy
    ```

To run a test, run:
    ```
    $ go test

    $ go test -v
    
    $ go test ./...
    
    $ go test --cover

    $ go test --coverprofile=coverage.out
    
    $ go tool cover --func=coverage.out

    $ go tool cover --html=coverage.out
    ```
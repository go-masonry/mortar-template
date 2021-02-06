# Mortar service template

This is a gRPC web service [template](https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template) using [Mortar](https://github.com/go-masonry/mortar)

For documentations and internals [read here](https://go-masonry.github.io)

## First steps

Once you download/clone this template, do the following

1. (_Optional_) When it's not yet a git repository

   Init Git and create a first commit, before any other changes

    ```sh
    git init
    git add -A
    git commit -m "initial commit"
    ```

2. (_Optional_) Download and install `protoc` compiler plugins if needed

   ```sh
   make go-install-deps
   ```

   > This will install [`protoc-gen-grpc-gateway`, `protoc-gen-openapiv2`, `protoc-gen-go`, `protoc-gen-go-grpc`].
   > For more information read [this](https://github.com/grpc-ecosystem/grpc-gateway#installation)

3. Start you service

    ```sh
    make run
    ```

4. Verify
    > from a different shell but in the same directory

    * **GRPC**

        ```sh
        grpcurl -plaintext \
                --import-path ./api \
                --import-path ./third_party/googleapis \
                -proto helloworld.proto \
                -d '{"name":"Mortar"}' \
                "localhost:5380" helloworld.Greeter/SayHello
        ```

        It's output should be

        ```json
        {
            "message": "Hello Mortar"
        }
        ```

    * **Public REST API**

        ```sh
        curl -XPOST -d '{"name": "Mortar"}' 'localhost:5381/v1/sayhello'
        ```

        It's output should be

        ```json
        {
            "message": "Hello Mortar"
        }
        ```

    * **Private API**
  
        ```sh
        curl localhost:5382/self/build
        ```

        And you should see something like this

        ```json
        {
        "git_commit": "5ebdaee",
        "version": "v1.2.3",
        "build_tag": "42",
        "build_time": "2021-01-30T15:47:15Z",
        "init_time": "2021-01-30T17:47:17.166625+02:00",
        "up_time": "1m24.145705728s",
        "hostname": "Tals-Mac-mini.lan"
        }
        ```

    > Server can be stopped with `CTRL+C`

## Mortar configurations

Everything related to Mortar setup and other Constructors are found in [app/mortar/*](app/mortar/).

Continue reading [here](app/mortar/README.md)

## Changing your API

This is a template project and as such it makes assumptions you will probably going to change.

By default this template comes as a `github.com/go-masonry/mortar-template` go module.

You will need to change:

1. **[GO Module](https://blog.golang.org/using-go-modules)**

   Delete `go.mod` and `go.sum` and recreate with the package name of your choice.

2. **Imports**

   By default this template API is compiled with the above import.
   Change `api/*.proto` files and recompile by running

    ```sh
    make gen-api
    ```

3. **Find & Replace/Update**

   `github.com/go-masonry/mortar-template` -> `<your-module-name>`

   * [`app/controllers/controller.go`](app/controllers/controller.go)
   * [`app/services/helloworld.go`](app/services/helloworld.go)
   * [`app/validations/helloworld.go`](app/validations/helloworld.go)
   * [`app/mortar/helloworld.go`](app/mortar/helloworld.go)
   * [`main.go`](main.go)

## Troubleshooting

1. _`protoc-gen-go`: program not found or is not executable_ means that `$GOBIN` is not in your `$PATH`

    Identify your `GOPATH`

    ```sh
    go env | grep GOPATH
    ```

    Set

    * For **fish shell**

        ```sh
        set -U fish_user_paths $GOPATH/bin $fish_user_paths
        ```

    * For **zsh**

        ```sh
        echo "export $GOPATH/bin:$PATH" >> ~/.zsh
        ```

2. Init Git and create a first commit.

    If not you will see something similar in your logs:

    ```sh
    git="fatal: not a git repository (or any of the parent directories): .git"
    ```
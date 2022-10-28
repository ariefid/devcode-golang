# DevCode GoLang

## About

The source code is for GetHired DevCode Golang by [Arief P](https://github.com/ariefid).

Todo application with user activities.

## Installation

-   Download or Clone this repository

    ```sh
    $ git clone https://github.com/ariefid/devcode-golang.git
    ```

-   Change directory to this repository directory.

    ```sh
    $ cd devcode-golang
    ```

-   Copy .env.example to .env:

    ```sh
    $ cp .env.example .env
    ```

-   Edit configuration in .env file (do not use root user! see: [https://stackoverflow.com/a/66910240](https://stackoverflow.com/a/66910240)):

    ```env
    MYSQL_PORT=3306
    MYSQL_USER=docker
    MYSQL_PASSWORD=docker
    MYSQL_DBNAME=todo
    ```

-   Run docker compose:

    ```sh
    $ docker compose up -d
    ```

-   Access [http://localhost:3030](http://localhost:3030) in your local computer.

## License

The source code is open-sourced software licensed under the [MIT license](./LICENSE).

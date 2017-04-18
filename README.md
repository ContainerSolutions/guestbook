Guestbook
---
stand-alone, single binary guestbook application in Go with PostgreSQL as its backend storage.

### Usage
For this application to work you need to pass the information (credentials, location, etc.) to
connect to PostgreSQL using the `postgres` flag. See the example below:

```
$ ./guestbook -postgres="postgres://user:pass@host/dbname?sslmode=disable"
```

For the full list of options for the connection string see the [libpq official documentation](https://godoc.org/github.com/lib/pq).

The application will be ready and listening on port 8080.

### Building
  - For OSX:

    ```
    $ make build-osx
    ```
  - For Linux:

    ```
    $ make build-osx
    ```

These tasks will create binary files for the mentioned platforms. Find the artifacts under `dist/`.

### Developing
#### Running Postgres
We run Postgres in a Docker container with little customisation. The only thing we customise is the password
which is `foobar`.

Run the container with
```
$ make pg
```

If you want to connect to the database using the psql tool, simply
```
$ make pg-client
```

#### Running the guestbook
There is a make task for that!
```
$ make run
```

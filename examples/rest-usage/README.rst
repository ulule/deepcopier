Introducing deepcopier, a Go library to make copying of structs a bit easier 
============================================================================

Context 
-------

We are currently refactoring our API_ at Ulule from our monolithic Python
stack with `django-tastypie`_ to a separate Go microservice_.

When working with models in Go, you don't want to expose all columns and
also implement more methods without writing a lot of code, because everyone
knows programmers are lazy ;)

deepcopier_ will help you in your daily job when you want to copy a struct into
another one (think resource) or from another one (think payload).

Installation
------------

Assuming you are already a Go developer, you have your environment up and ready,
so run this command in your shell:

::

    $ go get github.com/ulule/deepcopier

You are now ready to use this library.

Usage
-----

To demonstrate why you should use this library, we will build a dead simple REST
API in READ only.

We will use postgresql_ as database so I'm also assuming you
already have postgresql_ installed on your laptop :)

Let's create the databass!

::

    $ psql postgres
    psql (9.4.1)
    Type "help" for help.

    postgres=# create user dummy with password '';
    CREATE ROLE
    postgres=# create database dummy with owner dummy;
    CREATE DATABASE
    postgres=# \d
    No relations found.

We now have a perfectly capable database with no tables, let's jump to the
SQL schema.

.. code-block:: sql

    CREATE TABLE account (
        id serial PRIMARY KEY,
        first_name VARCHAR(50),
        last_name VARCHAR(50),
        username VARCHAR (50) UNIQUE NOT NULL,
        password VARCHAR (50) NOT NULL,
        email VARCHAR (355) UNIQUE NOT NULL,
        date_joined TIMESTAMP NOT NULL
    );

Transfer this schema to postgresql_.

::

    $ psql -U dummy
    psql (9.4.1)
    Type "help" for help.
    dummy=#     CREATE TABLE account(
    dummy(#         id serial PRIMARY KEY,
    dummy(#         first_name VARCHAR (50),
    dummy(#         last_name VARCHAR (50),
    dummy(#         username VARCHAR (50) UNIQUE NOT NULL,
    dummy(#         password VARCHAR (50) NOT NULL,
    dummy(#         email VARCHAR (355) UNIQUE NOT NULL,
    dummy(#         date_joined TIMESTAMP NOT NULL
    dummy(#     );
    CREATE TABLE
    dummy=# \d
                 List of relations
     Schema |      Name      |   Type   | Owner
    --------+----------------+----------+-------
     public | account        | table    | thoas
     public | account_id_seq | sequence | thoas
    (2 rows)

    dummy=#

First insertions incoming!

::

    dummy=# INSERT INTO account (username, first_name, last_name, password, email, date_joined) VALUES ('thoas', 'Florent', 'Messa', '8d56e93bcc8d63a171b5630282264341', 'foo@bar.com', '2015-07-31 15:10:10');

At this point, we have a schema in a great database, we need to setup our
REST API.

We will use:

* `go-json-rest`_ to handle requests
* gorm_ to manipulate the database as an ORM

In your shell, run this to install them

::

    $ go get -u github.com/jinzhu/gorm
    $ go get github.com/ant0ine/go-json-rest/rest

We will define a first attempt of our API to retrieve user information based
on its username.

We will rewrite our API three times so you need to focus.

.. code-block:: go

    // main.go
    package main

    import (
        "fmt"
        "github.com/ant0ine/go-json-rest/rest"
        "github.com/jinzhu/gorm"
        _ "github.com/lib/pq"
        "log"
        "net/http"
        "os"
        "time"
    )

    type Account struct {
        ID         uint `gorm:"primary_key"`
        FirstName  string
        LastName   string
        Username   string
        Password   string
        Email      string
        DateJoined time.Time
    }

    type Accounts struct {
        Db gorm.DB
    }

    func (a *Accounts) Detail(w rest.ResponseWriter, r *rest.Request) {
        account := &Account{}
        result := a.Db.First(&account, "username = ?", r.PathParam("username"))

        if result.RecordNotFound() {
            rest.NotFound(w, r)
            return
        }

        w.WriteJson(&account)
    }

    func main() {
        dsn := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
            os.Getenv("DATABASE_USER"),
            os.Getenv("DATABASE_NAME"))

        db, err := gorm.Open("postgres", dsn)

        fmt.Println(dsn)

        if err != nil {
            panic(err)
        }

        db.DB()
        db.DB().Ping()
        db.DB().SetMaxIdleConns(10)
        db.DB().SetMaxOpenConns(100)
        db.SingularTable(true)
        db.LogMode(true)

        api := rest.NewApi()

        api.Use(rest.DefaultDevStack...)

        accounts := &Accounts{Db: db}

        router, err := rest.MakeRouter(
            rest.Get("/users/:username", accounts.Detail),
        )
        if err != nil {
            log.Fatal(err)
        }
        api.SetApp(router)
        log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
    }

Let's start the server then

::

    $ DATABASE_USER=dummy DATABASE_NAME=dummy go run main.go

and retrieve the response.

::

    $ curl http://localhost:8080/users/thoas
    {
      "ID": 1,
      "Username": "thoas",
      "FirstName": "Florent",
      "LastName": "Messa",
      "Password": "8d56e93bcc8d63a171b5630282264341",
      "Email": "foo@bar.com",
      "DateJoined": "2015-07-31T15:10:10Z"
    }

Wait a minute? You are exposing the user's password... this not
what we are excepting... We want this specific format

.. code-block:: json

    {
      "id": 1,
      "username": "thoas",
      "first_name": "Florent",
      "last_name": "Messa",
      "name": "Florent Messa",
      "email": "foo@bar.com",
      "date_joined": "2015-07-31T15:10:10Z",
      "api_url": "http://localhost:8080/users/thoas"
    }

Implement a separate struct named ``AccountResource``

.. code-block:: go

    type AccountResource struct {
        ID         uint      `json:"id"`
        Username   string    `json:"username"`
        FirstName  string    `json:"first_name"`
        LastName   string    `json:"last_name"`
        Name       string    `json:"name"`
        Email      string    `json:"email"`
        DateJoined time.Time `json:"date_joined"`
    }

    func (a Account) Name() string {
        return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
    }

and rewrite ``Accounts.Detail`` to use deepcopier_

.. code-block:: go

    func (a *Accounts) Detail(w rest.ResponseWriter, r *rest.Request) {
        account := &Account{}
        result := a.Db.First(&account, "username = ?", r.PathParam("username"))

        if result.RecordNotFound() {
            rest.NotFound(w, r)
            return
        }

        resource := &AccountResource{}

        deepcopier.Copy(account).To(resource)

        w.WriteJson(&resource)
    }

We are good now, we can inspect our result

::

    $ curl http://localhost:8080/users/thoas
    {
      "id": 1,
      "username": "thoas",
      "first_name": "Florent",
      "last_name": "Messa",
      "name": "Florent Messa",
      "email": "foo@bar.com",
      "date_joined": "2015-07-31T15:10:10Z"
    }

Easy, right?

We will now rewrite for the last time ``Accounts.Detail`` to provide
some context to retrieve the base url in ``api_url`` attribute.

.. code-block:: go

    func (a *Accounts) Detail(w rest.ResponseWriter, r *rest.Request) {
        account := &Account{}
        result := a.Db.First(&account, "username = ?", r.PathParam("username"))

        if result.RecordNotFound() {
            rest.NotFound(w, r)
            return
        }

        resource := &AccountResource{}

        context := map[string]interface{}{"base_url": r.BaseUrl()}

        deepcopier.Copy(account).WithContext(context).To(resource)

        w.WriteJson(&resource)
    }

We need to update ``AccountResource`` to implement the ``ApiUrl`` new method

.. code-block:: go

    type AccountResource struct {
        ID         uint      `json:"id"`
        Username   string    `json:"username"`
        FirstName  string    `json:"first_name"`
        LastName   string    `json:"last_name"`
        Name       string    `json:"name"`
        Email      string    `json:"email"`
        DateJoined time.Time `json:"date_joined"`
        ApiUrl     string    `deepcopier:"context" json:"api_url"`
    }

    func (a Account) Name() string {
        return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
    }

    func (a Account) ApiUrl(context map[string]interface{}) string {
        return fmt.Sprintf("%s/users/%s", context["base_url"], a.Username)
    }

We have now the final result of what we excepted for the first time :)

::

    $ curl http://localhost:8080/users/thoas
    {
      "id": 1,
      "username": "thoas",
      "first_name": "Florent",
      "last_name": "Messa",
      "name": "Florent Messa",
      "email": "foo@bar.com",
      "date_joined": "2015-07-31T15:10:10Z",
      "api_url": "http://localhost:8080/users/thoas"
    }

If you have reached to the bottom you belong to the brave!

It has been a long introduction, hope your enjoy it!

Contributing to deepcopier
--------------------------

* Ping us on twitter `@oibafsellig <https://twitter.com/oibafsellig>`_, `@thoas <https://twitter.com/thoas>`_
* Fork the `project <https://github.com/ulule/deepcopier>`_
* Fix `bugs <https://github.com/ulule/deepcopier/issues>`_

Don't hesitate ;)


.. _API: http://developers.ulule.com/
.. _django-tastypie: https://github.com/django-tastypie/django-tastypie
.. _microservice: http://martinfowler.com/articles/microservices.html
.. _React.js: http://facebook.github.io/react/
.. _postgresql: http://www.postgresql.org/
.. _go-json-rest: https://github.com/ant0ine/go-json-rest
.. _gorm: https://github.com/jinzhu/gorm
.. _deepcopier: https://github.com/ulule/deepcopier

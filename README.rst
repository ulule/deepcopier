Deepcopier
==========

.. image:: https://secure.travis-ci.org/ulule/deepcopier.png?branch=master
    :alt: Build Status
    :target: http://travis-ci.org/ulule/deepcopier

This package is meant to make copying of structs to/from others structs a bit easier.


Installation
------------

::

    $ go get github.com/ulule/deepcopier

This package requires some dependencies:

* `oleiade/reflections <https://github.com/oleiade/reflections>`_: An awesome high level abstractions over reflect library

All dependencies will be installed for you.

Basic usage
-----------

.. code-block:: go

    // Deep copy instance1 into instance2
    Copy(instance1).To(instance2)

    // Deep copy instance1 into instance2 and passes the following context (which
    // is basically a map[string]interface{}) as first argument
    // to methods of instance2 that defined the struct tag "context".
    Copy(instance1).WithContext(map[string]interface{}{"foo": "bar"}).To(instance2)

    // Deep copy instance2 into instance1
    Copy(instance1).From(instance2)

    // Deep copy instance2 into instance1 and passes the following context (which
    // is basically a map[string]interface{}) as first argument
    // to methods of instance1 that defined the struct tag "context".
    Copy(instance1).WithContext(map[string]interface{}{"foo": "bar"}).From(instance2)

You should use the following struct tags:

* **field**: name of the field in the target instance
* **context**: method will take context (map[string]interface{}) as first argument
* **skip**: skip this field (do not process anything)

Example:

.. code-block:: go

    package main

    import (
        "fmt"
        "github.com/ulule/deepcopier"
    )

    // Model
    type User struct {
        Name string
    }

    func (u *User) MethodThatTakesContext(ctx map[string]interface{}) string {
        // do whatever you want
        return ""
    }

    // Resource
    type UserResource struct {
        DisplayName            string `deepcopier:"field:Name"`
        SkipMe                 string `deepcopier:"skip"`
        MethodThatTakesContext string `deepcopier:"context"`
    }

    func main() {
        user := &User{
            Name: "gilles",
        }

        resource := &UserResource{}

        deepcopier.Copy(user).To(resource)

        fmt.Println(resource.DisplayName)
    }

Looking for more information about the usage?

We wrote `an introduction article <https://github.com/ulule/deepcopier/blob/master/examples/rest-usage/README.rst>`_,
have a look and give us your feedback.

Contributing
------------

* Ping us on twitter `@oibafsellig <https://twitter.com/oibafsellig>`_, `@thoas <https://twitter.com/thoas>`_
* Fork the `project <https://github.com/ulule/deepcopier>`_
* Fix `bugs <https://github.com/ulule/deepcopier/issues>`_

Don't hesitate ;)

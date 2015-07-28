Deepcopier
==========

.. image:: https://secure.travis-ci.org/ulule/deepcopier.png?branch=master
    :alt: Build Status
    :target: http://travis-ci.org/ulule/deepcopier

This package is meant to make copying a bit more easy.

It will allow you to copy a struct from another one and also copy a struct to another one.


Installation
------------

::

    $ go get github.com/ulule/deepcopier

Usage
-----

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

* `field`: name of the field in the target instance
* `context`: method takes context (map[string]interface{}) as first argument
* `skip`: just skip this field (does not process anything)

Example:

.. code-block:: go

    // Model
    type User struct {
        Name   string
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

    user := &User{
        Name: "gilles",
    }

    resource := &UserResource{}

    Copy(user).To(resource)

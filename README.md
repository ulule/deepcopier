# Deepcopier

*Copy your struct like a boss.*

## Installation

```
$ go get github.com/ulule/deepcopier
```

## Usage

```go
// Deep copy instance1 into instance2
Copy(instance1).To(instance2)

// Deep copy instance1 into instance2 and passes the following context (which
// is basically a map[string]interface{}) as first argument
// to methods of instance2.
Copy(instance1).WithContext(map[string]interface{}{"foo": "bar"}).To(instance2)

// Deep copy instance2 into instance1
Copy(instance1).From(instance2)

// Deep copy instance2 into instance1 and passes the following context (which
// is basically a map[string]interface{}) as first argument
// to methods of instance1.
Copy(instance1).WithContext(map[string]interface{}{"foo": "bar"}).From(instance2)
```

The following struct tags are available:

* `field`: name of the field in the target instance
* `context`: auto-injects context as first argument to the given method
* `skip`: just skip this field (does not copy / do anything)

Example:

```go
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

Copy(&User).To(&UserResource)
```

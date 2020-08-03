## About Errors Package
Errors are one of the most important aspects of a programming language. The way you handle errors impacts the performance of the application.

Typical Error Handling in Golang:
```
Method 1:
    -  config, err := ioutil.ReadFile("config.json")
    if err!=nil{
        fmt.Print("Error:",err) // Regular logs which doesnot help in Realtime
    }
Method 2:
    - Panic (Use only which interrupts the normal operation of app)
Method 3:
    - Handle Gracefully (log the Errors & Return)
```
## Getting Started
Before getting started Ill try to explain the basic architecture of this project
<br><mark>Note: Not fully Implemented, since this project is just a demo</mark>
 #### <ins>3 Layer Architecture</ins>
| Layer | Purpose  | # Example File |
| ------- | --- | --- |
| Handler | Validates Request, gets required data from lower layers | [accounts.handler.go](https://github.com/chaitanya-apty/Go-logger-Implementation/blob/master/accounts/handler/accounts-handler.go) |
| Service | Middle man btw 1 & 3 layers, handles errors, transforms data | [accounts.service](https://github.com/chaitanya-apty/Go-logger-Implementation/blob/master/accounts/handler/accounts-service.go) |
| Repo | Queries Database/ Make other http requests(If in a micro-service) | [accounts.repo](https://github.com/chaitanya-apty/Go-logger-Implementation/blob/master/accounts/repo/accounts-repo.go) |

#### Advantages of this Pattern
-   Short story 😅, Independent of frameworks and testability
-   Long story 🥺 , [Read here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

### Handling errors gracefully
<b>Ill try to explain in 3 stages how Customized errors can be implemented with maintaing type checks</b>

<i>[You can jump here](https://github.com/chaitanya-apty/Go-logger-Implementation/blob/master/errors/errors.go) for the full code </i>
1.   `Create Custom defined Struct in order to wrap error`
```
    type Operation string // Contains the names of functions

    type ErrorType string // ex: NotFoundError

    const (
         NotFoundError ErrorType = "NOT_FOUND"
         UnAuthorizedError ErrorType = "UNAUTHORIZED"
         Unexpected ErrorType = "UNEXPECTED"
    )

    // Error - will contain the Error Object
    type Error struct {
        operations []Operation
        errorType ErrorType
        error error // Actual Error
    }
```
2. `Defining Some Helpers for getting and setting values`
```
    func NewError(operation Operation, errorType ErrorType, err error) *Error{
	return &Error{
		operations: []Operation{operation},
		errorType:  errorType,
		error:      err,
     }
    }

    func (e *Error) WithOperation(operation Operation) *Error{
	    e.operations = append(e.operations, operation)
	    return e
    }

    func (e *Error) Operations() []Operation{
	    return e.operations
    }
```
3. `Implement In Code`
    
    -   First, Catch and Create Error Object
    ```
    // accounts.repo (wrapping error with custom struct)
    const operation errors.Operation = "AccStore.GetUserById"
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Accounts ORDER BY id DESC")
    if err!=nil{
        return nil, errors.NewError(operation, errors.Unexpected, err, logrus.ErrorLevel)
    }
    ```
    -   Next, Pass the Error to Upper level with Operation
    ```
    //accounts.service
    const operation errors.Operation = "AccService.GetUser"
    User, err := AccStore.getUserById(id)
    if err != nil {
        return nil, err.WithOpetation(operation)
    }
    ```
    -   Log at handler layer / anywhere
    ```
    ex: If you log @ Handler layer when error is emitted,
    this will be the result

    code: log.Println(errors.Operation) // accounts.handler
    output: ["AccStore.GetUserById", "AccService.GetUser"] // clear stack info,where the error is emitted

    ```
## Final Thoughts
### In Golang, we have the liberty to handle the errors however we want.


## Authors
* **Chaitanya Kumar**
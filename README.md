# Queue-in-Go
This repository gives the implementation of queue, a kind of data structure, of which the elements are First-In/First-Out in Go.
The queue is implemented by a base slice, of which the size can be infinity, with methods: Offer(), Poll(), First(), Last(), Size(), IsEmpty();

Notes:
As for method like Poll(), First(), and Last(), the element returned is interface{} type. If you want to transform it into type int, type string and other types, write the following code in your program:
```go
switch v := element.(type) {
    case int:
        ...
    case string:
        ...
    ...
}
```

# The GOOD
I personally love Go and i use it for both scripting and writing actual structured software.

I like the way it has been structured and its features, from the ducktyping to the channels.

I like also the availability of tools for golang developers.

# The less good
Golang, as most awesome things, comes at its price. The complete lack of a toolkit to build UIs and the lack of constructors are the two things that come up first if I have to think about stuff i don't like in Go. 

This article is not about those things though.

# The horrible
What really chilled me to the bone where three (for now) "features"

## ""AWESOME"" SLICING

Let's talk about slices.

Slices are wonderful and extremely useful. If you know what you are doing.

If you come from python you probably have no idea on how slices internals work, because you don't need to.

Here is some python slices at work:
```python
a="AAAABBBB"
b=a[:4]
c=b+a
print(c)
# Yes this is python 3
#AAAAAAAABBBB
```
Easy. It just works. And if you are used to script in python this stuff is no surprise.

Let's take a look at Golang slices.

```go
//Let's create a new hash.
h := md5.New()
//Now let's hash "Foo" with it
d := h.Sum([]byte("Foo"))
/*
                          _             
__      ____ _ _ __ _ __ (_)_ __   __ _ 
\ \ /\ / / _` | '__| '_ \| | '_ \ / _` |
 \ V  V / (_| | |  | | | | | | | | (_| |
  \_/\_/ \__,_|_|  |_| |_|_|_| |_|\__, |
                                  |___/ 

This is where things break.
The := operator can fool you in a BAD way.
Since now "d" has been declared to the type 
that "Sum" returns, you don't need to know what 
that type is.
You would commonly expect the return type to be
a slice of bytes. And you would be wrong.

func Sum(data []byte) [Size]byte

The return type is an ARRAY of bytes.

And this is where the differences can be noted:
*/
d1 := d[:len(d)/2]
d2 := d[len(d)/2:]
/*
The two slices now represent the original array:

d: 466f6fd41d8cd98f00      b204e9800998ecf8427e,
d1:466f6fd41d8cd98f00, d2: b204e9800998ecf8427e
*/
d1[0] = 'B'
/*
Modifiying the d1 slice ALSO MODIFIED "d"
which now both start with 42 instead of 46

d:  426f6fd41d8cd98f00b204e9800998ecf8427e,
d1: 426f6fd41d8cd98f00, d2: b204e9800998ecf8427e

This is WEIRD. But there is WORSE
*/
d1 = append(d1, 'B')
/*
Appending to an array can modify also NEARBY slices

d:  426f6fd41d8cd98f00 wtf →4204e9800998ecf8427e,
d1: 426f6fd41d8cd98f00, d2: 4204e9800998ecf8427e
*/
d1 = append(d1, []byte("BBBBBBBBBBB")...)
/*
Unless you append too much, in which case a new array 
will be allocated for you.
d:  426f6fd41d8cd98f004204e9800998ecf8427e,
d1: 426f6fd41d8cd98f004242424242424242424242, d2: 4204e9800998ecf8427e
*/
```

IMHO this can be dangerous and you all gopher out there should always watch out for the type of the data you are handling.

Slices are safe to use, Arrays aren't.

## CLOSURE OF DOOM
__DISCLAIMER:__ Closures are a feature any language should have, and languages which don't have them are boring.

A Closure is basically a Lambda on steroids that can use variables accessible from its scope.

An example of closure:
```go
type Counter func() int

func CounterFactory(start int) Counter {
	return func() int {
		start++
		return start
	}
}
```
CounterFactory returns a function that, when called, increments the starting value and then returns it.

```go
counter:=CounterFactory(0)
counter() // Returns 1
counter() // Returns 2
```

Awesome, right?

Yes.

The problems start when you don't know well enough how Scoping is done in Go. 

If you use a closure in a range loop, for example:
```go
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d says %d\n", i, i)
			wg.Done()
		}()
	}
	wg.Wait()
}
```

Output:
```
Goroutine 10 says 10
Goroutine 10 says 10
Goroutine 10 says 10
Goroutine 10 says 10
Goroutine 10 says 10
Goroutine 10 says 10
Goroutine 10 says 10
Goroutine 10 says 10
Goroutine 10 says 10
Goroutine 10 says 10
```

### Why does it Happen?
Let's split up the scopes:
```
main{
	wg
	for{
	i
		{
			go closure
		}
	}
}
```
The "i" variable is declared in a scope that is outside the scope where the closure is built.

This means that the "i" variable is the same variable across every closure.

### How do I fix it?
Just move the "i" scope down.
```go
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i:=i //Yes, this is the right way to do it
		go func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d says %d\n", i, i)
			wg.Done()
		}()
	}
	wg.Wait()
}
```
If you don't trust me you can read this in [Effective Go](https://golang.org/doc/effective_go.html#channels)

# TL;DR
Go is a very powerful tool and using it without knowing it can cause big headaches.

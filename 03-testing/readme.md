![](/assets/microservicesWithGo.png)

# Testing in Go

test types
- Unit - most granular type of test and aimed at testing individual pieces and the smallest units possible. It tests at line level, individual function level
- Integration - tests code in its immediate surroundings. As soon as you are not using a mocked response it is an integration test.
- Functional - full running tests, or end-to-end tests.

We can visualise that in the context of a solution as follows, note that parts making up the integration test layer will themselves be subject to unit tests as the various sets in the functional test will have both unit and integration level test layers.

![](/03-testing/assets/gotesting.png)

With the types established we should now present the typical `perfect` scenario prescribed with the testing pyramid that sees 70% of total coverage at unit level, 20% at integration and 10% at end to end functional coverage. The model is great but we live in a imperfect world and sometimes we are forced to make tactical solutions, to ship software without adequate time for full testing coverage and other times we made be bolstering and retrofitting tests to existing software that is being reinvested in.

![](/03-testing/assets/gotesting_types.png)

## Unit testing

As we have said above the idea of the unit test is to test individual functions, statements and lines of code. A general rle of thumb is that we should have tests for every return of a function, that will typically catch `early return` for failure or bad scenarios as well as positive test cases and result returns.

In the example here, we're going to test the `user-dao.go` file from our domain layer. Specifically we're going to test the `GetUsers` function which has the following code

```go
// GetUser return the user or error
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	// implementation
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "Not found",
	}
}
```

We can observe that:
- We have two `returns` to test
- we have an error scenario where a user is not found.
- We have a return where a user was matched.

We can tell that the type of things we will want to check in the error situation are:
- an `error` was raised
- the error `Code` was as expected
- the error `Message` was correct
- the `http.StatusCode` is correct.

Just as for a valid situation we should ensure:
- the expected and actual `ID` do match
- Other values are as expected.

This fits the unit model in that we have a mocked database and this is a situation that can be mocked quite trivially and is isolated processing therefore it makes for a good unit test basis. So let's see how that looks. Note, we will use an additional import here as Go has no built-in assert library as found in some other languages but we have a 3rd party tool that offers the same job. If you object to a 3rd party addition for test code simplification you can achieve the same results with `if statements` and supply the testing object, the expected condition and the actual condition.

![](/03-testing/assets/unit-test.png)

## Running the tests

Now we can run the tests from the directory that contains them, which should be the same as the code itself, with `go test` or you can call the enclosing directory from the `$GOPATH` start point as the argument to `go test`.

eg: assume our code is in the `$GOPATH/src/github.com/username/mvc/domain` directory. If we pass this as the argument the `go` tool will find the correct file, one which mimics the name of the go file under test and with a `_test.go` suffix.

## Benchmarks

Here we're using a simple Bubble sort largely because its a notable inefficient way to sort. However, if we dig deeper we can see that actually a bubble sort is ne of the best operations and mot efficient if you have a very small number of elements, but the resource intensity of it is exponential, so as the elements number increases the performance drops off a cliff. We Can use the `testify` library again here with the `testing.B` for benchmarking.

As a comparison we're using a simple handcoded bubble sort up against the Go standard library `sort.Ints` and we can see that for cases of 10, 1000 & 100,000 numbers the go standard library is actually slower than a simple handcoded bubble sort, but as the number ramps up the standard library sort requires far fewer iterations to get a settled benchmark and it is significantly quicker than the one we have coded.

```go
package utils

// BubbleSort sorting algo
func BubbleSort(elements []int) []int {
	running := true

	for running {
		running = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				running = true
			}
		}
	}
	return elements
}
```

and we will test the code with the following routines.

```go
package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCaseScenario(t *testing.T) {
	// initialisation section
	// creates a slice where every element will require processing/swapping
	eles := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	// execution section
	eles = BubbleSort(eles)

	// validation section
	// tests return is not nil
	assert.NotNil(t, eles)

	// check length of return is equal to input
	assert.EqualValues(t, 9, len(eles))

	// checks individual elements against
	// expected sort ascending
	assert.EqualValues(t, 1, eles[0])
	assert.EqualValues(t, 2, eles[1])
	assert.EqualValues(t, 3, eles[2])
	assert.EqualValues(t, 4, eles[3])
	assert.EqualValues(t, 5, eles[4])
	assert.EqualValues(t, 6, eles[5])
	assert.EqualValues(t, 7, eles[6])
	assert.EqualValues(t, 8, eles[7])
	assert.EqualValues(t, 9, eles[8])

}

func TestBubbleSortBestCaseScenario(t *testing.T) {
	eles := []int{1, 2, 3, 4, 5}
	eles = BubbleSort(eles)

	assert.NotNil(t, eles)
	assert.EqualValues(t, 5, len(eles))

	// checks individual elements against
	assert.EqualValues(t, 1, eles[0])
	assert.EqualValues(t, 2, eles[1])
	assert.EqualValues(t, 3, eles[2])
	assert.EqualValues(t, 4, eles[3])
	assert.EqualValues(t, 5, eles[4])
}

func TestBubbleSortNilCaseScenario(t *testing.T) {
	eles := BubbleSort(nil)
	assert.Nil(t, eles)
}

func getElements(n int) []int {
	res := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		res[i] = j
		i++
	}
	return res
}

func BenchmarkBubbleSort1000(b *testing.B) {
	eles := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(eles)

	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	eles := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(eles)

	}
}

func BenchmarkSort1000(b *testing.B) { // we're still slower on 1000 elements
	eles := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(eles)

	}
}
```

When we ramp this up to 100,000

![](/03-testing/assets/benchmark.png)
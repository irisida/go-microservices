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
- We have two returns to test
- we have an error scenario where a user is not found.
- We have a return where a user was matched.

We can tell that the type of things we will want to check in the error situation are:
- an error was raised
- the error code was as expected
- the error message was correct
- the http.StatusCode is correct.

Just as for a valid situation we should ensure:
- the expected and actual ID do match
- Other values are as expected.

This fits the unit model in that we have a mocked database and this is a situation that can be mocked quite trivially and is isolated processing therefore it makes for a good unit test basis. So let's see how that looks. Note, we will use an additional import here as Go has no built-in assert library as found in some other languages but we have a 3rd party tool that offers the same job. If you object to a 3rd party addition for test code simplification you can achieve the same results with `if statements` and supply the testing object, the expected condition and the actual condition.

![](/03-testing/assets/unit-test.png)
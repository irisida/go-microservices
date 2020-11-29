![](/assets/microservicesWithGo.png)

# Testing in Go

test types
- Unit - most granular type of test and aimed at testing individual pieces and the smallest units possible. It tests at line level, individual function level
- Integration - tests code in its immediate surroundings. As soon as you are not using a mocked response it is an integration test.
- Functional - full running tests, or end-to-end tests.

We can visualise that in the context of a solution as follows, note that parts making up the integration test layer will themselves be subject to unit tests as the various sets in the functional test will have both unit and integration level test layers.

![](/03-testing/assets/gotesting.png)
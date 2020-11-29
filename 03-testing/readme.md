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

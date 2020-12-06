![](/assets/microservicesWithGo.png)

# Mocking

Mocking is a way of extending what we have done in the previous section. To successfully mock our system parts well we need to think about the structure of the project.
- service reshuffle to split / clarify
- adds code to controller but clarity is greater
- separates functional use, but allows for inclusion in the package.

The aim of mocking is to be able to follow the same API and the same functions and flow, sequence but by removing the need for an actual database, or actual other components. So, we can create a faked response of positive and negative use cases to test the intended flows, the exceptions and errors. 

![](04-mocking/assets/04-services.png)

![](04-mocking/assets/04-controller.png)

![](04-mocking/assets/04-dao.png)

![](04-mocking/assets/04-test.png)
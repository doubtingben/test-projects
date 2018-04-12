# INFURA Infrastructure Test

This is a take-home test for the INFURA infrastructure team that we would like 
you to attempt. You'll find a number of steps to complete below that will 
require some coding on your part; return any code and/or documents you wrote 
to us when you feel comfortable with the result.

Feel free to use the programming/scripting language you're most comfortable 
with (although Python and Go are preferred) and tools, libraries, or 
frameworks you believe are best suited to the tasks listed below. There is no 
strict time limit, but if possible please return this to us within one week.

Please note and let us know how long you worked on the test; this is for 
informational purposes and allows us to make adjustments and improvements for 
future applicants (feel free to provide feedback on the test too).


1. Register for an [INFURA API key](https://infura.io/register.html)
    1. You will have to use this key for subsequent requests to INFURA endpoints, 
    as briefly shown in the [Get Started](https://infura.io/#how-to) section of the site
2. Create an application that retrieves Ethereum Mainnet transaction and block 
data via the INFURA JSON-RPC API from 
[https://pmainnet.infura.io](https://pmainnet.infura.io/)
    1. See [the INFURA API docs](https://infura.io/docs/#supported-json-rpc-methods) 
    for a list of supported JSON-RPC methods
    2. See [the Ethereum docs](https://github.com/ethereum/wiki/wiki/JSON-RPC) for information on the 
    Ethereum API itself
3. Expose the retrieved transaction and block data via REST endpoints that your 
application provides
    1. To sanity check your results, feel free to use the similar functionality
    at the [INFURA Hub](https://hub.infura.io/mainnet)
4. Set up your application to run in a [Docker container](https://www.docker.com)
5. Create a load test for your application
6. Run some load test iterations and document the testing approach and the 
results obtained
    1. Specify some performance expectations given the load test results: 
    e.g., this application is able to support X requests per minute
7. Write up a short document describing the general setup of the components 
you've put together as well as instructions to run your application
8. **Bonus points**: add unit tests to cover most of the code you've written
9. Submit your application and load test code, as well as associated 
documentation, to the master branch of the Github repository 
we've set up for this purpose

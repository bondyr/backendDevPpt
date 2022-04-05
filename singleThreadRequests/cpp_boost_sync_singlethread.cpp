#include "../common/common.h"


int main(int argc, char** argv)
{
    try
    {
        req.set(http::field::host, host);
        req.set(http::field::user_agent, BOOST_BEAST_VERSION_STRING);
        auto [shouldDisplay, nrOfRequests] = processCmdLineArguments(argc, argv);
        // -------------------------------------------------------------

        cout << "\n=== Waiting for " << nrOfRequests << " responses.....\n\n";

        for (int i = 0; i < nrOfRequests; ++i) {
            makeRequest(shouldDisplay);
        }

        cout << "\n=== " << nrOfRequests << " responses received.\n";
    }
    catch(std::exception const& e)
    {
        std::cerr << "Error: " << e.what() << std::endl;
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}

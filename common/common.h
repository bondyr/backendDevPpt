#include <boost/beast/core.hpp>
#include <boost/beast/http.hpp>
#include <boost/beast/version.hpp>
#include <boost/asio/connect.hpp>
#include <boost/asio/ip/tcp.hpp>
#include <cstdlib>
#include <iostream>
#include <string>
#include <thread>
#include <vector>
#include <algorithm>

namespace beast = boost::beast;
namespace http = beast::http;
namespace net = boost::asio;

using tcp = net::ip::tcp;
using namespace std;

auto const host = "localhost";
auto const port = "3000";
auto const target = "/";
int version = 11;

// Set up an HTTP GET request message
http::request<http::string_body> req{http::verb::get, target, version};

// The io_context is required for all I/O
net::io_context ioc;

// These objects perform our I/O
tcp::resolver resolver(ioc);

// Performs an HTTP GET and prints the response
void makeRequest(bool shouldDisplay)
{
    beast::tcp_stream stream(ioc);

    // Look up the domain name
    auto const results = resolver.resolve(host, port);

    // Make the connection on the IP address we get from a lookup
    stream.connect(results);

    // Send the HTTP request to the remote host
    http::write(stream, req);

    // This buffer is used for reading and must be persisted
    beast::flat_buffer buffer;

    // Declare a container to hold the response
    http::response<http::string_body> res;

    // Receive the HTTP response
    // cout << "Waiting for response\n";
    http::read(stream, buffer, res);
    
    if (shouldDisplay) {
        cout << "Response: " << res.body() << "\n";
    }

    // Gracefully close the socket
    beast::error_code ec;
    stream.socket().shutdown(tcp::socket::shutdown_both, ec);
};

struct CmdArgs {
    bool shouldDisplay;
    int nrOfRequests;
};

CmdArgs processCmdLineArguments(int argc, char **argv)
{
    if (argc != 3)
    {
        std::cerr << "Usage: main shouldDisplay nrOfRequests\n"
                  << "Example:\n"
                  << "    main 1 100\n";
        throw EXIT_FAILURE;
    }

    const bool shouldDisplay = std::atoi(argv[1]);
    const auto nrOfRequests = std::atoi(argv[2]);

    return CmdArgs{shouldDisplay, nrOfRequests};
}

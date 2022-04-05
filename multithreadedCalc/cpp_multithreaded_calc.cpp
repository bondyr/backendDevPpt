#include <cstdlib>
#include <iostream>
#include <string>
#include <thread>
#include <vector>
#include <algorithm>

using namespace std;

void doSth()
{
    // expensive calculations
    auto i = 3;
    i *= 234423423;
}

int main(int argc, char **argv)
{
    try
    {
        auto const nrOfThreads = 10 * 1000;
        // auto const nrOfThreads = 100 * 1000;

        std::vector<std::thread> vec;
        vec.reserve(nrOfThreads);

        for (int i = 0; i < nrOfThreads; ++i) {
            vec.push_back(std::thread(doSth));
        }

        for (auto &th : vec)
        {
            if (th.joinable())
                th.join();
        }

        std::cout << "Processing done\n";
    }
    catch (std::exception const &e)
    {
        std::cerr << "Error: " << e.what() << std::endl;
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}

#include <primesieve.hpp>
#include <iostream>
#include <vector>

int main()
{
  // store the primes below 1000
  std::vector<int> primes;
  primesieve::generate_primes(1000, &primes);

  primesieve::iterator it;
  int prime = it.next_prime();

  // iterate over the primes below 10^6
  for (; prime < 1000000; prime = it.next_prime())
    std::cout << prime << std::endl;

  return 0;
}
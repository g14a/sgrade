#include <primesieve.hpp>
#include <iostream>
#include <vector>

using namespace std;

using namespace std;

bool isPrime(int n) { 
    if (n <= 1)  return false; 
    if (n <= 3)  return true; 
  
    if (n%2 == 0 || n%3 == 0) return false; 
  
    for (int i=5; i*i<=n; i=i+6) 
        if (n%i == 0 || n%(i+2) == 0) 
           return false; 

    return true;
}


int main()
{
  // store the primes below 1000
  std::vector<int> primes;

  for(int i=0;i<10000;i++) {
    if(isPrime(i)) {
      primes.push_back(i);
    }
  }

  // primesieve::generate_primes(1000, &primes);

  // primesieve::iterator it;
  // uint64_t prime = it.next_prime();

  // // iterate over the primes below 10^6
  // for (; prime < 1000000; prime = it.next_prime())
  //   std::cout << prime << std::endl;

  for(vector<int>::iterator it = primes.begin(); it != primes.end(); ++it) {
    cout << *it << endl;
  }

  return 0;
}
#include <primesieve.hpp>
#include <iostream>
#include <vector>
#include <cmath>
#include <string>
#include <fstream>
#include <omp.h>

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

bool IsPower(int num, int power) {
    float err = pow(double(num), 1/double(power));
    
    float t = abs(err-round(err));

    if (t<0.0000001) {
        return true;
    }

    return false;
}

int GetNthRoot(int a, int b) {
    int s = pow((float)(a), 1/(float)(b));

    return (int)round(s);
}

void PerfectPowers(vector<int> primes, int lower, int upper, int maxPower, ofstream &file) {
    vector<int> primebuf;
    primesieve::iterator it;
    int prime = it.next_prime();

    for (prime = lower; prime < upper; prime = it.next_prime()) {
        primebuf.push_back(prime);

        for (int index=0; index < primebuf.size(); index++) {
            int prime = primebuf[index];
            int sum = 0;

            for (int i=index; i < primebuf.size(); i++) {
                sum += primebuf[i];
            }

        #pragma omp parallel for
            for (int power = 2; power <= maxPower; power++) {
                if(IsPower(sum, power)) {
                    string s = to_string(prime) + ":" + to_string(prime) + " = " + to_string(sum) + " = " + to_string(GetNthRoot(sum, power)) + "**" + to_string(power) + "\n";
                    file << s;
                }
            }
        }
    }
}

int main(int argc, char **argv) {

    int lowerBound = atoi(argv[1]);
    int upperBound = atoi(argv[2]);
    int maxPower = atoi(argv[3]);
    string filename = argv[4];

    std::vector<int> primes;
    primesieve::generate_primes(lowerBound, upperBound, &primes);

    ofstream file (filename, ofstream::out);

    long double start = omp_get_wtime();

    PerfectPowers(primes, lowerBound, upperBound, maxPower, file);

    cout << omp_get_wtime() - start << endl;
    
    return 0;
}

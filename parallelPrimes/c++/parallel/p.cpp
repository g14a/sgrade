#include <primesieve.hpp>
#include <iostream>
#include <vector>
#include <cmath>
#include <string>

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

void PerfectPowers(vector<int> primes, int maxPower) {
    vector<int> primebuf;

    for (vector<int>::iterator it = primes.begin(); it != primes.end(); ++it) {
        primebuf.push_back(*it);
       // int dep = *it;

        for (int index=0; index < primebuf.size(); index++) {
            int prime = primebuf[index];
            int sum = 0;
           // int t = dep;

            for (int i=index; i < primebuf.size(); i++) {
                sum += primebuf[i];
            }

            for (int power = 2; power <= maxPower; power++) {

                // cout << "Sum = " << sum << " and power = " << power << endl;
                if(IsPower(sum, power)) {

                    string s = to_string(prime) + ":" + to_string(*it) + " = " + to_string(sum) + " = " + to_string(GetNthRoot(sum, power)) + "**" + to_string(power);
                    cout << s << endl;
                }
            }
        }
    }
}

int main()
{
  // store the primes below 1000
  std::vector<int> primes;
  primesieve::generate_primes(10000, &primes);

  PerfectPowers(primes, 5);
  
  return 0;
}
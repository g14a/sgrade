#include<iostream>
#include<cmath>

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

    cout << err << endl;

    int t = abs(err-round(err));

    cout << round(err) << endl;
    cout << t << endl;

    if (t<0.0000001) {
        return true;
    }

    return false;
}

int main() {
    if (IsPower(81, 3)) {
        cout << "Yes" << endl;
    } else {
        cout << "NO" << endl;
    }

    return 0;
}
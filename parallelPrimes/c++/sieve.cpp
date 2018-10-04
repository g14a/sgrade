#include<iostream>
#include<cstdlib>
#include<vector>

using namespace std;

void sieve(int n) {

    vector<bool> prime;
    
    for(int i = 0; i*i <= n; i++) {
        if(prime[i] == true) {
            for(int j=i*2;j<=n;j+=i) {
                prime.push_back(false);
            }
        }
    }

    for(auto p=prime.begin();p!=prime.end();p++) {
        if(prime[*p]) {
            cout << *p << " ";
        }
    }
}

int main() {
    int n=30;
    cout << "int n\n";

    sieve(n);
    return 0;
}
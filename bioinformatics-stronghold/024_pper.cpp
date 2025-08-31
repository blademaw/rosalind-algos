#include <bits/stdc++.h>

using namespace std;

long long mod_factorial(long long k) {
  if (k == 0) {
    return 1;
  }
  return (k * mod_factorial(k-1)) % 1'000'000;
}


long long mod_ncr(long long n, long long r) {
  double sum = 1;
  for (int i = 1; i <= r; i++){
    sum = sum * (n - r + i) / i;
  }
  return (long long)sum % 1'000'000;
}


int main (int argc, char *argv[]) {
  int n, k; cin >> n >> k;

  cout << (mod_ncr(n, k) * mod_factorial(k)) % 1'000'000 << "\n";

  return 0;
}

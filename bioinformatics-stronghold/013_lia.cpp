#include <cmath>
#include <iostream>

double nCr(int n, int r) {
    if (r > n || r < 0) return 0;
    if (r == 0 || r == n) return 1;
    if (r > n - r) r = n - r;
    double result = 1;
    for (int i = 1; i <= r; i++) {
        result = result * (n - i + 1) / i;
    }
    return result;
}

int main (int argc, char *argv[]) {
  int k, n;
  std::cin >> k >> n;
  int total{(int)std::pow(2, k)};

  double prob{0};
  for (int i=n; i <= total; i++) {
    prob += nCr(total, i) * std::pow(1/4., i) * std::pow(3/4., total - i);
  }

  std::cout << prob << "\n";
  
  return 0;
}

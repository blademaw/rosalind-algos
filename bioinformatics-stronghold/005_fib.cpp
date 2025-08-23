#include <iostream>

int main (int argc, char *argv[]) {
  int n, k;
  std::cin >> n >> k;

  long long fs[40] {0};
  fs[0] = 1;
  fs[1] = 1;

  for (int i = 2; i < n; i++) {
    fs[i] = fs[i-1] + k*fs[i-2];
  }

  std::cout << fs[n-1] << "\n";

  return 0;
}

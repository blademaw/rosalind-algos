#include <iostream>

int main (int argc, char *argv[]) {
  int k, m, n;
  std::cin >> k >> m >> n;

  float s = k + m + n;
  float sp = s - 1.0;

  float p = (k/s)*((k-1)/sp + m/sp + n/sp) + (m/s)*(k/sp + 0.75*(m-1)/sp + 0.5*n/sp) + (n/s)*(k/sp + 0.5*m/sp);

  std::cout << p << "\n";

  return 0;
}

#include <iostream>
#include <vector>

int main (int argc, char *argv[]) {
  int n, m;
  std::cin >> n >> m;

  std::vector<long long> dp(n);
  dp[0] = 1; dp[1] = 1;

  for (int i=2; i < n; i++) {
    if (i < m) {
      dp[i] = dp[i-1] + dp[i-2];
    } else {
      long long s{0};
      for (int j{1}; j < m; j++) {
        s += dp[i-j-1];
      }
      dp[i] = s;
    }
  }

  std::cout << dp[n-1] << "\n";

  return 0;
}

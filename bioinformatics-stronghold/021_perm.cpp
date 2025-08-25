#include <iostream>
#include <vector>

void permutations(std::vector<std::vector<int>>& sols, int k, std::vector<int>& a) {
  if (k == 1) {
    sols.push_back(a);
  } else {
    permutations(sols, k-1, a);
    for (int i{}; i < k - 1; i++) {
      if (k % 2 == 0) {
        std::swap(a[i], a[k-1]);
      } else {
        std::swap(a[0], a[k-1]);
      }
      permutations(sols, k-1, a);
    }
  }
}

void print_arr(const std::vector<int>& a) {
  for (int i{}; i < a.size(); i++) {
    std::cout << a[i];
    if (i < a.size() - 1) std::cout << " ";
  }
  std::cout << "\n";
}

int main (int argc, char *argv[]) {
  int n; std::cin >> n;
  std::vector<int> ns(n);
  for (int i{}; i < n; i++) {
    ns[i] = i+1;
  }

  std::vector<std::vector<int>> sols;
  permutations(sols, n, ns);

  std::cout << sols.size() << "\n";
  for (const auto& s : sols) {
    print_arr(s);
  }
  // heaps_algorithm(n, ns);

  return 0;
}

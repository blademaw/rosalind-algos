#include <iostream>
#include <vector>

int main (int argc, char *argv[]) {
  std::string s, t;
  std::cin >> s >> t;

  std::vector<int> res{};

  for (int i{}; i <= s.length() - t.length(); i++) {
    if (s.substr(i, t.length()) == t) {
      res.push_back(i+1);
    }
  }

  for (int i{}; i < res.size(); i++) {
    if (i == res.size()-1) {
      std::cout << res[i] << "\n";
    } else {
      std::cout << res[i] << " ";
    }
  }

  return 0;
}

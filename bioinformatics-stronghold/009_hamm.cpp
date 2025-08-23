#include <iostream>

int main (int argc, char *argv[]) {
  std::string s1, s2;
  std::cin >> s1 >> s2;

  int d{0};
  for (int i{}; i < s1.length(); i++) {
    if (s1[i] != s2[i]) {
      d++;
    }
  }

  std::cout << d << "\n";

  return 0;
}

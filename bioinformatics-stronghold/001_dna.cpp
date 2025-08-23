#include <iostream>
#include <string>

int main (int argc, char *argv[]) {
  std::string s;
  std::cin >> s;

  int a {}, c {}, g {}, t{};
  for (char ch : s) {
    switch (ch) {
      case 'A': a++; break;
      case 'C': c++; break;
      case 'G': g++; break;
      case 'T': t++; break;
    }
  }

  std::cout << a << " " << c << " " << g << " " << t << "\n";
  
  return 0;
}

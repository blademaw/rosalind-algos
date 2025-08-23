#include <iostream>

int main (int argc, char *argv[]) {
  std::string s;
  std::cin >> s;

  for (int i = 0; i < s.length(); i++) {
    if (s[i] == 'T') {
      s[i] = 'U';
    }
  }
  std::cout << s << "\n";
  return 0;
}

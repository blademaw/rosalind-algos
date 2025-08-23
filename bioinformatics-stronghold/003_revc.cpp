#include <iostream>

void rev_str(std::string &s) {
  for (int i=0; i < (int)s.length()/2; i++) {
    char tmp_c = s[i];
    s[i] = s[s.length()-i-1];
    s[s.length()-i-1] = tmp_c;
  }
}

void complement_str(std::string &s) {
  for (int i=0; i < s.length(); i++) {
    switch (s[i]) {
      case 'A': s[i] = 'T'; break;
      case 'T': s[i] = 'A'; break;
      case 'G': s[i] = 'C'; break;
      case 'C': s[i] = 'G'; break;
    }
  }
}

int main (int argc, char *argv[]) {
  std::string s;
  std::cin >> s;

  rev_str(s);
  complement_str(s);

  std::cout << s << "\n";
  return 0;
}

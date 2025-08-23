#include <iostream>
#include "utils.cpp"

int main (int argc, char *argv[]) {
  std::string s;
  std::cin >> s;

  double res{0};
  for (auto& c : s) {
    res += mono_mass_table[c];
  }
  
  printf("%.10f\n", res);

  return 0;
}

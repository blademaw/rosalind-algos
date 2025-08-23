#include <iostream>
#include <string>
#include <unordered_map>
#include "utils.cpp"

int main (int argc, char *argv[]) {
  std::string s;
  std::cin >> s;

  std::unordered_map<char, int> codonCounts;
  for (auto [codon, c] : codon_table) {
    codonCounts[c] += 1;
  }

  int res{1};
  for (auto c : s) {
    res = (res * codonCounts[c]) % 1'000'000;
  }
  res = (res * codonCounts['*']) % 1'000'000; // add in stop codon

  std::cout << res << "\n";

  return 0;
}

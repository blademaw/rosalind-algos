#include <iostream>
#include <string>
#include <unordered_map>

std::unordered_map<std::string, char> codon_table = {
  {"UUU", 'F'}, {"CUU", 'L'}, {"AUU", 'I'}, {"GUU", 'V'},
  {"UUC", 'F'}, {"CUC", 'L'}, {"AUC", 'I'}, {"GUC", 'V'},
  {"UUA", 'L'}, {"CUA", 'L'}, {"AUA", 'I'}, {"GUA", 'V'},
  {"UUG", 'L'}, {"CUG", 'L'}, {"AUG", 'M'}, {"GUG", 'V'},
  {"UCU", 'S'}, {"CCU", 'P'}, {"ACU", 'T'}, {"GCU", 'A'},
  {"UCC", 'S'}, {"CCC", 'P'}, {"ACC", 'T'}, {"GCC", 'A'},
  {"UCA", 'S'}, {"CCA", 'P'}, {"ACA", 'T'}, {"GCA", 'A'},
  {"UCG", 'S'}, {"CCG", 'P'}, {"ACG", 'T'}, {"GCG", 'A'},
  {"UAU", 'Y'}, {"CAU", 'H'}, {"AAU", 'N'}, {"GAU", 'D'},
  {"UAC", 'Y'}, {"CAC", 'H'}, {"AAC", 'N'}, {"GAC", 'D'},
  {"UAA", '*'}, {"CAA", 'Q'}, {"AAA", 'K'}, {"GAA", 'E'},
  {"UAG", '*'}, {"CAG", 'Q'}, {"AAG", 'K'}, {"GAG", 'E'},
  {"UGU", 'C'}, {"CGU", 'R'}, {"AGU", 'S'}, {"GGU", 'G'},
  {"UGC", 'C'}, {"CGC", 'R'}, {"AGC", 'S'}, {"GGC", 'G'},
  {"UGA", '*'}, {"CGA", 'R'}, {"AGA", 'R'}, {"GGA", 'G'},
  {"UGG", 'W'}, {"CGG", 'R'}, {"AGG", 'R'}, {"GGG", 'G'}
};

int main (int argc, char *argv[]) {
  std::string s;
  std::cin >> s;

  std::string res;
  int i{}, j{};
  while (i <= s.length()) {
    char c = codon_table[s.substr(i, 3)];
    if (c == '*') {
      break;
    } else {
      res.push_back(c);
      i += 3;
      j += 1;
    }
  }

  std::cout << res << "\n";
  return 0;
}

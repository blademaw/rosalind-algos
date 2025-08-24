#include <iostream>
#include <vector>
#include "utils.cpp"

int main (int argc, char *argv[]) {
  auto [names, seqs] = parse_rosalind_fasta_input(std::cin);
  std::vector<std::vector<int>> counts(4, std::vector<int>(seqs[0].length()));
  std::vector<int> max_counts(seqs[0].length(), 0);
  std::vector<char> max_chars(seqs[0].length(), '_');

  for (int i{}; i < names.size(); i++) {
    for (int j{}; j < seqs[i].length(); j++) {
      int cur_count{0};
      switch (seqs[i][j]) {
        case 'A': cur_count = ++counts[0][j]; break;
        case 'C': cur_count = ++counts[1][j]; break;
        case 'G': cur_count = ++counts[2][j]; break;
        case 'T': cur_count = ++counts[3][j]; break;
        default: std::cout << "Non-ACGT character in sequence; exiting...\n"; exit(1);
      }
      if (cur_count > max_counts[j]) {
        max_counts[j] = cur_count;
        max_chars[j] = seqs[i][j];
      }
    }
  }

  for (const auto& c : max_chars) {
    std::cout << c;
  }
  std::cout << "\n";

  char labels[4] = {'A', 'C', 'G', 'T'};
  for (int i{}; i < 4; i++) {
    std::cout << labels[i] << ": ";
    for (const auto& c : counts[i]) {
      std::cout << c << " ";
    }
    std::cout << "\n";
  }

  return 0;
}

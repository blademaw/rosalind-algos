#include <iostream>
#include <string_view>
#include <utility>
#include "utils.cpp"
#include "trie.cpp"

int main (int argc, char *argv[]) {
  auto [names, seqs] = parse_rosalind_fasta_input(std::cin);
  std::vector<std::pair<size_t, size_t>> pairs;

  std::string_view s = seqs[0];
  for (size_t i = 0; i < s.length(); i++) {
    for (size_t len = 4; len <= 12; len++) {
      if (i + len > s.length()) {
        break;
      }

      // PERF: this could be faster with a read-only version of reverse_complement
      if (s.substr(i, len) == reverse_complement((std::string)s.substr(i, len))) {
        pairs.push_back({i, len});
      }
    }
  }

  for (const auto& [a, b] : pairs) {
    std::cout << a+1 << " " << b << "\n";
  }

  return 0;
}

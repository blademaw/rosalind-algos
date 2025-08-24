#include <iostream>
#include "utils.cpp"

int main (int argc, char *argv[]) {
  auto [_, seqs] = parse_rosalind_fasta_input(std::cin);
  auto s = seqs[0];

  for (int i{1}; i < seqs.size(); i++) {
    bool found_all = false;
    
    while (!found_all) {
      auto pos = s.find(seqs[i]);
      if (pos != std::string::npos) {
        s = s.erase(pos, seqs[i].length());
      } else {
        found_all = true;
      }
    }
  }

  std::cout << build_protein_str(dna_to_rna(s), false) << "\n";

  return 0;
}

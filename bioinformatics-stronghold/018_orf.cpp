#include <iostream>
#include <unordered_set>
#include "utils.cpp"

std::string build_protein_str(const std::string& s) {
  std::string res = "M";
  char codon;
  int j{0};
  while (j < s.length()-3+1) {
    codon = codon_table[s.substr(j, 3)];
    if (codon == '*') {
      return res;
    }
    res += codon;
    j += 3;
  }

  return "";
}

int main (int argc, char *argv[]) {
  auto [names, seqs] = parse_rosalind_fasta_input(std::cin);
  auto seq_rev = reverse_complement(seqs[0]);
  auto seq_rna = dna_to_rna(seqs[0]);
  auto seq_rna_r = dna_to_rna(seq_rev);

  std::unordered_set<std::string> candidates;
  for (int i{}; i < seq_rna.length()-3+1; i++) {
    if (seq_rna.substr(i, 3) == "AUG") {
      auto candidate = build_protein_str(seq_rna.substr(i+3));
      if (candidate != "") {
        candidates.insert(candidate);
      }
    }
  }
  for (int i{}; i < seq_rna_r.length()-3+1; i++) {
    if (seq_rna_r.substr(i, 3) == "AUG") {
      auto candidate = build_protein_str(seq_rna_r.substr(i+3));
      if (candidate != "") {
        candidates.insert(candidate);
      }
    }
  }

  for (const auto& candidate : candidates) {
    std::cout << candidate << "\n";
  }

  return 0;
}

#pragma once
#include <iostream>
#include <istream>
#include <string>
#include <tuple>
#include <unordered_map>
#include <vector>

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

std::unordered_map<char, float> mono_mass_table = {
  {'A', 71.03711},
  {'C', 103.00919},
  {'D', 115.02694},
  {'E', 129.04259},
  {'F', 147.06841},
  {'G', 57.02146},
  {'H', 137.05891},
  {'I', 113.08406},
  {'K', 128.09496},
  {'L', 113.08406},
  {'M', 131.04049},
  {'N', 114.04293},
  {'P', 97.05276},
  {'Q', 128.05858},
  {'R', 156.10111},
  {'S', 87.03203},
  {'T', 101.04768},
  {'V', 99.06841},
  {'W', 186.07931},
  {'Y', 163.06333}
};


// Convert a DNA sequence into RNA
std::string dna_to_rna(std::string s) {
  for (int i = 0; i < s.length(); i++) {
    if (s[i] == 'T') {
      s[i] = 'U';
    }
  }
  return s;
}

// Reverse a string
std::string reverse_str(std::string s) {
  for (int i=0; i < (int)s.length()/2; i++) {
    char tmp_c = s[i];
    s[i] = s[s.length()-i-1];
    s[s.length()-i-1] = tmp_c;
  }
  return s;
}

// Get the reverse complement of a DNA sequence
std::string reverse_complement(std::string s) {
  for (int i=0; i < s.length(); i++) {
    switch (s[i]) {
      case 'A': s[i] = 'T'; break;
      case 'T': s[i] = 'A'; break;
      case 'G': s[i] = 'C'; break;
      case 'C': s[i] = 'G'; break;
    }
  }

  s = reverse_str(s);

  return s;
}


// Get a list of names and corresponding sequences from a Rosalind-style FASTA
// input
std::tuple<std::vector<std::string>, std::vector<std::string>>
parse_rosalind_fasta_input(std::istream& in_stream) {
  std::string line{}, cur{};
  std::vector<std::string> names{}, seqs{};

  // parse all FASTA strs
  while (std::getline(in_stream, line)) {
    if (line[0] == '>') {
      if (cur.size() > 0) {
        seqs.push_back(cur);
      }

      names.push_back(std::string(line.begin() + 1, line.end()));
      cur = "";
    } else {
      cur += line;
    }
  }
  seqs.push_back(cur);

  return std::make_tuple(names, seqs);
}


// Build a protein string from an RNA string
std::string build_protein_str(const std::string& s, bool start_with_M = true) {
  std::string res;
  if (start_with_M) {
    res = "M";
  }

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

  return res;
}

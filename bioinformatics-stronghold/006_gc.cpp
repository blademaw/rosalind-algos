#include <iostream>
#include <string>
#include <vector>

struct Seq {
  std::string label;
  std::string seq;
};


double gc_content(std::string& s) {
  int t{0};
  for (char c : s) {
    if (c == 'C' || c == 'G') {
      t++;
    }
  }
  return (double)t/s.length();
}


int main (int argc, char *argv[]) {
  std::vector<Seq> seqs;
  std::string line;
  Seq cur;

  // parse all FASTA strs
  while (std::getline(std::cin, line)) {
    if (line[0] == '>') {
      if (cur.seq.size() > 0) {
        seqs.push_back(cur);
      }

      cur = {std::string(line.begin() + 1, line.end()), ""};
    } else {
      cur.seq += line;
    }
  }
  seqs.push_back(cur);

  // set temp max as first
  Seq* max_seq = &seqs[0];
  double max_seq_gc = gc_content(max_seq->seq);

  // find maximum
  for (Seq& s : seqs) {
    double gc = gc_content(s.seq);
    if (gc > max_seq_gc) {
      max_seq = &s;
      max_seq_gc = gc;
    }
  }
  std::cout << max_seq->label << "\n" << max_seq_gc*100.0 << "\n";

  return 0;
}

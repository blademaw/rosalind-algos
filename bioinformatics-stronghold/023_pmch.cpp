#include <bits/stdc++.h>
#include "utils.cpp"
#include "bigint.h"

using namespace std;

bigint factorial(bigint n) {
  if (n <= (bigint)1) {
    return 1;
  }
  return n * factorial(n - (bigint)1);
}

int main (int argc, char *argv[]) {
  auto [names, seqs] = parse_rosalind_fasta_input(cin);
  int a_count{0}, c_count{0};

  for (const auto& c : seqs[0]) {
    if (c == 'A') {
      a_count++;
    } else if (c == 'C') {
      c_count++;
    }
  }

  cout << factorial(a_count) * factorial(c_count) << "\n";

  return 0;
}

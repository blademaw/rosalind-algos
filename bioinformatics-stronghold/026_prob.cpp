#include <bits/stdc++.h>

using namespace std;

int main (int argc, char *argv[]) {
  string s; cin >> s;
  vector<float> gcs;
  float gc; while (cin >> gc) {
    gcs.push_back(gc);
  }

  array<int, 4> counts {0};
  for (const auto& c : s) {
    switch (c) {
      case 'A': counts[0]++; break;
      case 'C': counts[1]++; break;
      case 'G': counts[2]++; break;
      case 'T': counts[3]++; break;
    }
  }

  // P(A) = P(T) = (1-gc)/2; P(G) = P(C) = gc/2
  // P(Sequence) = P(x1)*P(x2)*...*P(xn) if x is a symbol; all P(x) are given above
  // => log10(P(Sequence)) = log10(P(x1)) + log10(P(x2)) + ... + log10(P(xn))
  // = count(A)*log10(P(A)) + count(C)*log10(P(C)) + ...
  // = (count(A) + count(T))*log10(P(A)) + (count(C) + count(G))*log10(P(C))
  // since P(A)=P(T) and P(C)=P(G).

  for (const auto& gc : gcs) {
    printf("%.3f ", (counts[0] + counts[3]) * log10((1. - gc)/2.) + (counts[1] + counts[2]) * log10(gc/2.));
  }
  cout << "\n";

  return 0;
}

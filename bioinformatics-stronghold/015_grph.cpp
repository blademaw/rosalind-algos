#include <algorithm>
#include <climits>
#include <iostream>
#include <tuple>
#include <vector>

int main (int argc, char *argv[]) {
  std::string line, cur{};
  std::vector<std::string> names{}, seqs{};

  // parse all FASTA strs
  while (std::getline(std::cin, line)) {
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

  std::vector<std::tuple<int, int>> edges{};

  for (int i{}; i < names.size(); i++) {
    for (int j{}; j < names.size(); j++) {
      if (i != j &&
        seqs[i] != seqs[j] && 
        seqs[i].substr(seqs[i].length()-3) == seqs[j].substr(0, 3)) {
        edges.push_back(std::make_tuple(i, j));
      }
    }
  }

  for (const auto& [u, v] : edges) {
    std::cout << names[u] << " " << names[v] << "\n";
  }

  return 0;
}

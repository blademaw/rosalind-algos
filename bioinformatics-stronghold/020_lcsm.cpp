#include <iostream>
#include <memory>
#include <numeric>
#include <string>
#include "utils.cpp"

// Solution using suffix tries (quite slow -- ~24 seconds for the full dataset).

struct Node {
  char c;
  Node* parent;
  std::vector<std::unique_ptr<Node>> children;
  std::vector<bool> reachable; // whether we can reach this node from i-th string
  
  Node(char character = '\0', Node* par = nullptr)
    : c(character), parent(par) {}
};

// Either traverse to a valid node given character, i-th string out of n
// strings, or insert a new node in the suffix trie
Node* traverse_suffix_trie(Node& node, char c, int i, int n) {
  for (const auto& child : node.children) {
    if (child->c == c) {
      return child.get();
    }
  }

  auto new_child = std::make_unique<Node>(c, &node);
  // initialise reachable array to true for i-th string
  std::vector<bool> reachable(n, false);
  new_child->reachable = std::move(reachable);

  Node* child_ptr = new_child.get();
  node.children.push_back(std::move(new_child));

  return child_ptr;
}

// Extend the suffix trie with suffix s of i-th string out of n strings
void extend_suffix_trie(Node& root, const std::string& s, int i, int n) {
  Node* cur_node = &root;

  for (const auto& c : s) {
    cur_node = traverse_suffix_trie(*cur_node, c, i, n);
    cur_node->reachable[i] = true;
    // std::cout << "After traversing for " << c << ", sizes are: " << cur_node->children.size() << ", " << cur_node->reachable.size() << "\n";
  }
}

void dfs_longest_common_substring(const Node* node, const Node*& best_node, int& max_depth, int n, int cur_depth = 0) {
  if (std::accumulate(node->reachable.begin(), node->reachable.end(), 0) == n) {
    if (cur_depth > max_depth) {
      max_depth = cur_depth;
      best_node = node;
    }

    for (const auto& child : node->children) {
      dfs_longest_common_substring(child.get(), best_node, max_depth, n, cur_depth + 1);
    }
  }
}

std::string reconstruct_suffix(const Node* node) {
  std::string s{};
  auto cur_node = node;
  while (cur_node->parent != nullptr) {
    s += cur_node->c;
    cur_node = cur_node->parent;
  }

  return reverse_str(s);
}

int main (int argc, char *argv[]) {
  auto [names, seqs] = parse_rosalind_fasta_input(std::cin);
  int n = names.size();

  Node root;
  root.reachable = std::move(std::vector<bool>(n, true)); // think of as empty string (always reachable)

  std::cout << "Extending suffix trie...\n";
  for (int i{}; i < n; i++) {
    for (int j{}; j < seqs[i].length(); j++) {
      extend_suffix_trie(root, seqs[i].substr(j), i, n);
    }
  }

  const Node* best_node = nullptr;
  int max_depth = 0;
  std::cout << "Finding longest common substring...\n";
  dfs_longest_common_substring(&root, best_node, max_depth, n);

  if (best_node == nullptr) {
    std::cout << "No common substring found.\n";
  } else {
    std::cout << reconstruct_suffix(best_node) << "\n";
    // std::cout << "Substring is " << reconstruct_suffix(best_node) << " @ depth " << max_depth << "\n";
  }

  return 0;
}

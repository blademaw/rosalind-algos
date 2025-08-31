#include <bits/stdc++.h>

using namespace std;

// number of connections required is just the number of disjoint/independent
// components minus one (|V| vertices requires |V|-1 edges to be connected)
// TODO: but what about creating cycles? I guess we assume this can't happen?

void dfs(int u, vector<vector<int>>& edges, vector<bool>& visited) {
  if (visited[u]) {
    return;
  }
  visited[u] = true;
  for (const auto& v : edges[u]) {
    dfs(v, edges, visited);
  }
  return;
}

int main (int argc, char *argv[]) {
  int n; cin >> n;
  vector<vector<int>> edges(n);
  for (auto& u : edges) {
    u.reserve(2*n);
  }

  int u, v;
  while (cin >> u >> v) {
    edges[u-1].push_back(v-1);
    edges[v-1].push_back(u-1);
  }

  // traverse (DFS) repeatedly until we visit every node
  vector<bool> visited(n, false);
  int components = 0;

  for (int u = 0; u < n; u++) {
    if (visited[u]) {
    } else {
      // explore from u
      dfs(u, edges, visited);
      components++;
    }
  }

  cout << components - 1 << "\n";

  return 0;
}

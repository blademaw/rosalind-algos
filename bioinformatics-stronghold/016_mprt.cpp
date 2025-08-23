#include <iostream>
#include <curl/curl.h>
#include <string>
#include <vector>
#include <fstream>

/* NOTE: This problem took me five submissions to solve since IDs that Rosalind
 * provides have changed in the UniProt database since 2012 (when the question
 * was released).
 */

static size_t WriteCallback(void *contents, size_t size, size_t nmemb, std::string *userp)
{
    size_t totalSize = size * nmemb;
    userp->append((char*)contents, totalSize);
    return totalSize;
}

int main (int argc, char *argv[]) {
  // std::ofstream file; // optionally save results to temp file
  // file.open("016_temp.txt");
  CURL *curl;
  CURLcode res;
  std::string url, response;
  std::vector<std::string> names{}, seqs{};

  while (std::getline(std::cin, url)) {
    names.push_back(url);
    url = url.substr(0, url.find_first_of('_'));
    url = "http://www.uniprot.org/uniprot/" + url + ".fasta";
    // url = "https://rest.uniprot.org/unisave/" + url + "?format=fasta&versions=36";
    // url = "https://rest.uniprot.org/unisave/" + url + "?format=fasta";
    std::cout << "Getting url " << url << "...";

    response.clear();

    curl = curl_easy_init();
    if (curl) {
      curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
      curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);
      curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, WriteCallback);
      curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);

      res = curl_easy_perform(curl);
      if (res != CURLE_OK) {
        std::cout << "Curl failed.\n";
        exit(1);
      } else {
        // file << response << "\n";

        std::cout << "OK.\n";
        std::string clean{};
        for (const auto& c : response.substr(response.find_first_of('\n')+1)) {
          if (c != '\n') {
            clean += c;
          }
        }
        std::cout << clean << "\n";
        seqs.push_back(clean);
      }

      curl_easy_cleanup(curl);
    }
  }
  // file.close();
  
  std::vector<int> idx{};
  for (int i{}; i < names.size(); i++) {
    idx.clear();
    auto s = seqs[i];

    for (int j{}; j < s.length()-4+1; j++) {
      if (s[j] == 'N' &&
          s[j+1] != 'P' &&
          (s[j+2] == 'S' || s[j+2] == 'T') &&
          s[j+3] != 'P') {
        idx.push_back(j);
      }
    }

    if (idx.size() > 0) {
      std::cout << names[i] << "\n";
      for (const auto& i : idx) {
        std::cout << i+1 << " ";
      }
      std::cout << "\n";
    }
  }
  
  return 0;
}

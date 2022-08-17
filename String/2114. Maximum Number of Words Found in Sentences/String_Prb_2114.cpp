#include <iostream>
#include <cstring>
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    int mostWordsFound(vector<string>& sentences) {
        int size = sentences.size();
        int maxI = INT_MIN;
        for(int i=0; i<size; ++i){
            int words = count(sentences[i].begin(),sentences[i].end(),' ');
            maxI = max(words,maxI);
        } 
        return maxI + 1;
    }
};

int main(){
	std::vector<std::string> operations = {"alice and bob love leetcode","i think so too","this is great thanks very much"};
	std::cout << "Output: " << (new Solution)->mostWordsFound(operations) << std::endl;
	return 0;
}
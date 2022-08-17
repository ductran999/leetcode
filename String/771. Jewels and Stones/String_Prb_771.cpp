#include <iostream>
#include <cstring>
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    int numJewelsInStones(string jewels, string stones) {
        int countJewel = 0;
        for(int i = 0; i < stones.length(); ++i){
            for(int j = 0; j < jewels.length(); ++j){
                if(stones[i] == jewels[j]){
                    ++countJewel;
                    continue;
                } 
            }
        }
        return countJewel;
    }
};

int main(){
	std::string jewels = "aA";
    std::string stones = "aAAbbbb";
	std::cout << "Output: " << (new Solution)->numJewelsInStones(jewels, stones) << std::endl;
	return 0;
}
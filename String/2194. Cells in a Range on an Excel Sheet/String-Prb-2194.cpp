#include <iostream>
#include <cstring>
#include <vector>

class Solution {
public:
    std::vector<std::string> cellsInRangestd(std::string s) {
        std::vector<std::string> ans;
        for(char i = s[0]; i <= s[3]; ++i){
            std::string tmp;
            tmp.push_back(i);
            for(char j = s[1]; j <= s[4]; ++j){
                tmp.push_back (j);
                ans.push_back (tmp);
                tmp.pop_back();
            }
            tmp.pop_back();
        }
        return ans;
    }
};

int main(){
	std::string s = "K1:L2";
    std::vector<std::string> ans =  (new Solution)->cellsInRangestd(s);
	for(auto tmp : ans){
        std::cout << tmp << ", ";
    }

	return 0;
}
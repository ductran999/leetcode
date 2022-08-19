#include <iostream>
#include <cstring>
#include <vector>

class Solution {
public:
    std::string restoreString(std::string s, std::vector<int>& indices) {
        int size = indices.size();
        std::string ans;
        for(int i = 0; i < size; ++i) ans.push_back('a');
        for(int i = 0; i < size; ++i){
            ans[indices[i]] = s[i];
        }
        return ans;
    }
};

int main(){
	std::string s = "codeleet"; std::vector<int> indices {4,5,6,7,0,2,1,3};
	std::cout << (new Solution)->restoreString(s, indices) << std::endl;
	return 0;
}
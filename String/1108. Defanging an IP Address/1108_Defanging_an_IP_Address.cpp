#include <iostream>
#include <cstring>

class Solution {
public:
    std::string defangIPaddr(std::string address) {
        std::string ans = "";
        for(char ad : address){
            if(ad == '.'){
               ans += "[.]"; 
                continue;
            }
            ans+=ad;
        }
        return ans;
    }
};

int main(){
	std::string input; std::cin >> input;
	std::cout << (new Solution)->defangIPaddr(input) << std::endl;
	return 0;
}
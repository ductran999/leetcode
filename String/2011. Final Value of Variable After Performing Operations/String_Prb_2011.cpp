#include <iostream>
#include <cstring>
#include <vector>

class Solution {
public:
    int finalValueAfterOperations(std::vector<std::string>& operations) {
        int x=0;
        for(int i = 0; i < operations.size(); i++){
            if(operations[i][0]  == '-' or operations[i][2] == '-') --x;
            else ++x;
        }
        return x;
    }
};
int main(){
	std::vector<std::string> operations = {"--X", "++X", "X++"};
	std::cout << "Output: " << (new Solution)->finalValueAfterOperations(operations) << std::endl;
	return 0;
}
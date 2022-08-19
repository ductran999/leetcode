#include <iostream>
#include <cstring>
#include <vector>
using namespace std;

class Solution {
public:
    int balancedStringSplit(string s) {
        int cntL = 0, cntR = 0, ans = 0;
        for(int i = 0; i < s.length(); ++i){
            if(s[i] == 'L') ++cntL;
            else if(s[i] == 'R') ++cntR;
            if(cntL == cntR){
                ++ans;
            }
        }
        return ans;
    }
};

int main(){
	std::string s = "RLRRLLRLRL";
    int maximumBalanceString =  (new Solution)->balancedStringSplit(s);
	cout << maximumBalanceString << endl;

	return 0;
}
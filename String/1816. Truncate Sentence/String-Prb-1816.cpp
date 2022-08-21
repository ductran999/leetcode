#include <iostream>
#include <cstring>

using namespace std;

class Solution {
public:
    string truncateSentence(string s, int k) {
        string ans = "";
        s+=" ";
        for(int i = 0; i < s.length(); ++i){
            if(s[i] == ' ') k--;
            if(k == 0) break;
            ans.push_back(s[i]);
        }
        return ans;
    }
};

int main(){
    string s = "Hello how are you Contestant"; int k = 4;
    cout <<  (new Solution) -> truncateSentence(s, k) << endl;
    return 0;
}
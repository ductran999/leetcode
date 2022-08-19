#include <iostream>
#include <cstring>
#include <vector>
using namespace std;

class Solution {
public:
    string sortSentence(string s) {
        vector<string> v(10,"x");
        s+=" ";
        string tmp = "";
        for(int i = 0; i < s.length(); ++i){
            if(s[i] == ' '){
                int idx = tmp[tmp.size()-1] - '0';
                tmp.pop_back();
                v[idx] = tmp;
                tmp.clear();
            }else tmp.push_back(s[i]);
        }
        string ans = "";
        for(int i = 1; i < 10; i++){
            if(v[i] != "x") ans += (v[i] + " ");
            else break;
        }
        ans.pop_back();
        return ans;
    }
};

int main(){
    string input = "is2 sentence4 This1 a3";
    cout << (new Solution) ->sortSentence (input) << endl;
	return 0;
}
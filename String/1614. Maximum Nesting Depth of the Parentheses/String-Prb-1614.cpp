#include <iostream>
#include <cstring>
#include <stack>

using namespace std;

class Solution {
public:
    int maxDepth(string s) {
        if(s.empty()) return 0;
        stack<int> st;
        st.push(0);
        int maxDepth = 0;
        for(int i = 0; i < s.length(); ++i){
            if(s[i] == '('){
                st.push(st.top()+1);
            }else if(s[i] == ')'){
                maxDepth = max(st.top(), maxDepth); st.pop();
            }
        }
        return maxDepth;
    }
};

int main()
{
    string input = "(1+(2*3)+((8)/4))+1";
    cout << (new Solution)->maxDepth(input) << endl;

    return 0;
}
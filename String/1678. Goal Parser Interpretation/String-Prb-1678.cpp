#include <iostream>
#include <cstring>
#include <stack>

using namespace std;

class Solution {
public:
    string interpret(string command) {
        string ans ="";
        for (int i = 0; i < command.length(); ++i) {
            if(command[i] == 'G') 
            {
                ans+='G';
                continue;
            }
            if(command[i] == '(' and command[i+1] == ')') {
                ans+= 'o';
                i += 1;
                continue;
            }
            ans += "al";
            i += 3;
        }
        return ans;
    }
};

int main(){
	string command = "G()(al)";
	cout << "Output: " << (new Solution) -> interpret(command) << endl;
	return 0;
}
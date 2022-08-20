#include <iostream>
#include <cstring>
#include <vector>
#include <algorithm>
#include <map> 
#include <set>

using namespace std;

class Solution {
public:
    
    vector<string> morse = {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};
    
    string enCode(string input){
        string tmp = "";
        for(int i = 0; i < input.length(); ++i){
            int idx = input[i]-'a';
            string charEncode = morse[idx];
            tmp += charEncode;
        }
        return tmp;
    }

    int uniqueMorseRepresentations(vector<string>& words) {
        set<string> ans;
        for(int i = 0; i < words.size(); ++i){
            string tmp = enCode(words[i]);
            ans.insert(tmp);
        }
        return ans.size();
    }
};

int main(){
	std::vector<string> words = {"gin","zen","gig","msg"};
	cout << (new Solution) -> uniqueMorseRepresentations(words) << endl;
	return 0;
}
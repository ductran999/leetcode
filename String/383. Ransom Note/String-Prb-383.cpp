#include <iostream>
#include <vector>
using namespace std;

class Solution {
public: 
    bool canConstruct(string ransomNote, string magazine) {
        int n = ransomNote.size();
        int m = magazine.size();
        if(n > m) return false;
        int freq[26] = {0};
        for(int i = 0; i < m; ++i) freq[magazine[i]-'a']++;
        for(int i = 0; i < n; ++i){
            if(freq[ransomNote[i] - 'a'] > 0) 
                freq[ransomNote[i]-'a']--; 
            else return false;
        }
        return true;
    }
};
int main() {
    string ransomNote = "aa";
    string magazine = "aaabbb";
    cout << (new Solution)->canConstruct(ransomNote, magazine) << endl;
    return 0;
}
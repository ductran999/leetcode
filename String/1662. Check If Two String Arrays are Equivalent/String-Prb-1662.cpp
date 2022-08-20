#include <iostream>
#include <cstring>
#include <vector>
#include <algorithm>
#include <map>
#include <set>

using namespace std;

class Solution {
public:
    bool arrayStringsAreEqual(vector<string>& word1, vector<string>& word2) {
        string w1 = ""; string w2 = "";
        for(int i = 0; i < word1.size(); ++i) w1 += word1[i];
        for(int i = 0; i < word2.size(); ++i) w2 += word2[i];
        return w1 == w2;
    }
};

int main()
{
    vector<string> word1 = {"ac", "b"}, word2 = {"ab", "c"};
    cout << (new Solution) -> arrayStringsAreEqual(word1, word2) << endl;

    return 0;
}
#include <iostream>
#include <cstring>
#include <vector>
#include <algorithm>
#include <map>
#include <set>
#include <bits/stdc++.h>

using namespace std;

class Solution
{
public:
    int countConsistentStrings(string allowed, vector<string> &words)
    {
        unordered_set<char> uSet;
        for (int i = 0; i < allowed.length(); ++i)
        {
            uSet.insert(allowed[i]);
        }
        int cnt = 0;
        for (int i = 0; i < words.size(); ++i)
        {
            bool valid = true;
            for (int j = 0; j < words[i].size(); ++j)
            {
                if (uSet.find(words[i][j]) == uSet.end())
                {
                    valid = false;
                    break;
                }
            }
            cnt += valid;
        }
        return cnt;
    }
};

int main()
{
    string allow = "abc";
    vector<string> words = {"a", "b", "c", "ab", "ac", "bc", "abc"};
    cout << (new Solution)->countConsistentStrings(allow, words) << endl;

    return 0;
}
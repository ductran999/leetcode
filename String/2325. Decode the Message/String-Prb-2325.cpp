#include <iostream>
#include <cstring>
#include <vector>
#include <map>
using namespace std;

class Solution {
public:
    string decodeMessage(string key, string message) {
        int val = 0;
        map<char, int> m;
        for (int i = 0; i < key.length(); i++)
        {
            if (key[i] != ' ' && (m.find(key[i]) == m.end()))
            {
                m.insert({key[i], val++});
            }
        }
        string deCode = "";
        for(int i = 0; i < message.length(); ++i){
            if(message[i] != ' '){
                 deCode += m.at(message[i]) + 'a';
            } else deCode += ' ';
        }
        return deCode;
    }
};

int main()
{
    string key = "the quick brown fox jumps over the lazy dog";
    string messages = "vkbs bs t suepuv";
    cout << (new Solution)->decodeMessage(key, messages) << endl;

    return 0;
}
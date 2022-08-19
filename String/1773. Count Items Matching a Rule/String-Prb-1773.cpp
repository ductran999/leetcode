#include <iostream>
#include <cstring>
#include <vector>
using namespace std;

class Solution {
public:
    int countMatches(vector<vector<string>>& items, string ruleKey, string ruleValue) {
        int cnt = 0, cols = 0;
        if(ruleKey == "type") cols = 0;
        else if (ruleKey == "color") cols = 1;
        else cols = 2;
        for(int i = 0; i < items.size(); ++i){
            if(items[i][cols] == ruleValue) ++cnt;
        }
        return cnt;
    }
};

int main(){
    vector<vector<string>> items = {{"phone","blue","pixel"},{"computer","silver","lenovo"},{"phone","gold","iphone"}};
    std::string ruleValue = "silver";
	std::string ruleKey = "color";
    int numberOfValidItem =  (new Solution)-> countMatches(items, ruleKey, ruleValue);
	cout << numberOfValidItem << endl;

	return 0;
}
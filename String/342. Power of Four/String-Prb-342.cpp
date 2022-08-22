#include <iostream>
#include <cstring>

using namespace std;

class Solution {
public:    
    bool isPowerOfFour(int n) {
        if(n == 1) return true;
        if(n <= 0) return false;
        while(n >= 4){
            if(n % 4 == 0) n/=4;
            else return false;
        }
        return n == 1;
    }
};

int main(){
    int n = 16;
    cout <<  (new Solution) -> isPowerOfFour(n) << endl;
    return 0;
}
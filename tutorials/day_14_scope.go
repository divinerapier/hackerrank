package tutorials

/*
#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

class Difference {
    private:
    vector<int> elements;

  	public:
  	int maximumDifference;
	// Add your code here
    Difference(vector<int>& v) {
        elements = v;
        max = 0;
        min = 101;
    }


    void computeDifference() {
        for (auto it = elements.begin(); it != elements.end(); ++it) {
            if (*it < min) {
                min = *it;
            }
            if (*it > max) {
                max = *it;
            }
        }
        maximumDifference = max - min;
    }

private:
    int max ;
    int min ;
    }; // End of Difference class

int main() {
    int N;
    cin >> N;

    vector<int> a;

    for (int i = 0; i < N; i++) {
        int e;
        cin >> e;

        a.push_back(e);
    }

    Difference d(a);

    d.computeDifference();

    cout << d.maximumDifference;

    return 0;
}

*/

#include <bits/stdc++.h>
using namespace std;

void solve (vector<int>& v, int lim) {
    sort (v.begin(), v.end());
    int start = 0;
    int end = v.size()-1;
    int ans = 0;
    while (start <= end) {
        if (v[end] > lim) {
            cout << "-1" << endl;
            return;
        }
        // else if (start == end && v[start] < lim) {
        //     ans++;
        //     start++;
        //     end--;
        // }
        if (v[start] + v[end] <= lim) {
            ans++;
            start++;
            end--;
        }
        else {
            ans++;
            end--;
        }
        // cout << v[start]+v[end] << endl;
    }
    cout << ans << endl;
}


int main () {
    vector<int> v;
    int n; cin >> n;
    for (int i=0; i<n; i++) {
        int d; cin >> d;
        v.push_back(d);
    }
    int lim; cin >> lim;
    solve (v, lim);
    return 0;
}
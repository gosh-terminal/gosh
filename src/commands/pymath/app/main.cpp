#include "lib/pymathlib.h"
using namespace std;
int main() {
    string input;
    cout << "Welcome to pymath available calculators include algebra, chemistry, geometry, physics\nWhich do you want?: ";
    cin >> input;
    if (input == "algebra") {
        pymath::algebra();
    } else if (input == "geometry") {
        pymath::geometry();
    } else if (input == "physics") {
        pymath::physics();
    } else if (input == "chemistry") {
        pymath::chemistry();
    }
}
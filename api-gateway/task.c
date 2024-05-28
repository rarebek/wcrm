#include <iostream>
#include <cmath>  

int main() {
    double x, y, a, b;
    

    std::cout << "X uchun qiymat kiriting: ";
    std::cin >> x;
    std::cout << "y uchun qiymat kiriting: ";
    std::cin >> y;
    std::cout << "a uchun qiymat kiriting: ";
    std::cin >> a;
    std::cout << "b uchun qiymat kiriting: ";
    std::cin >> b;
    
    // a)
    double partA = sqrt(2 * x - y) + sqrt(2 * y - x);
    std::cout << "A misol natijasi: " << partA << std::endl;

    // b)
    double partB = (sin(x) * sin(x) - cos(x) * cos(x)) / cos((x + y) / 4);
    std::cout << "B misol natijasi: " << partB << std::endl;
    
    // c)
    double partC = 1 / (5 * M_PI) + 1 / (1 + pow((x - b) / a, 3));
    std::cout << "C misol natijasi: " << partC << std::endl;

    return 0;
}

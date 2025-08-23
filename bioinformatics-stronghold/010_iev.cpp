#include <iostream>

int main (int argc, char *argv[]) {
  double a{}, b{}, c{}, d{}, e{}, f{};
  std::cin >> a >> b >> c >> d >> e >> f;

  double res{};
  res += a*2.0*1.0;
  res += b*2.0*1.0;
  res += c*2.0*1.0;
  res += d*2.0*0.75;
  res += e*2.0*0.5;

  std::cout << res << "\n";

  return 0;
}

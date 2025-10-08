#include "../ImageBlur.hpp"
#include <iostream>
#include <vector>

int main() {
  const int width = 16;
  const int height = 16;
  std::vector<unsigned char> img(width * height * RGB_CHANNELS, 200);
  for (int i = 0; i < width * height * RGB_CHANNELS; i++) {
  img[i] = i % 255;
  }

  ImageBlur blur(25);
  blur.blur_image(img.data(), height, width, 0, height);

  for (int i = 0; i < height; i++) {
    for (int j = 0; j < width; j++) {
      std::cout << static_cast<int>(img[i * width + j])  << " ";
    }
    std::cout << std::endl;
  }

  std::cout << "Blur complete" << std::endl;
  return 0;
}

#define RGB_CHANNELS 3


// Extremely fast Image Blur class for RGB IMAGES.
class ImageBlur {
private:
  int blur_strength;
  void blur_image_helper(unsigned char *imgBuf, int height, int width, int start,
                         int end, int blur_strength);
  int blur_1d_v_simd(int x, int width, int *horizontal_sums[],
                     unsigned char *out_row, int blur_strength);
  int blur_1d_3c(int x, int width, unsigned char *cur_src, int src_width,
                 int *temp_out, int blur_strength);
  int reflect_index(int i, int n);

public:
  ImageBlur(int blur_strength = 25) : blur_strength(blur_strength){};

  // The image buf to blur.
  // Height of the image.
  // Width of the image.
  // The position for start from on the height.
  // The position to end  bluring from on the height.
  void blur_image(unsigned char *imgBuf, int height, int width, int start, int end);
};

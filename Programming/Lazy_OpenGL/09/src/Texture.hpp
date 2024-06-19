#ifndef TEXTURE_HPP
#define TEXTURE_HPP

#include "LOpenGL.hpp"
#include "Rect.hpp"
#include <iostream>

class Texture
{
 public:
  Texture();
  ~Texture();

  bool loadFromPixels32(GLuint* pixels,
      GLuint imgWidth, GLuint imgHeight, GLuint texWidth, GLuint texHeight);
  bool loadFromFile(std::string path);
  void free();
  void render(GLfloat x, GLfloat y, Rect<GLfloat>* clip = nullptr);
  GLuint getID() { return _ID; }
  GLuint texWidth() { return _texWidth; }
  GLuint texHeight() { return _texHeight; }
  GLuint imgWidth() { return _imgWidth; }
  GLuint imgHeight() { return _imgHeight; }

  bool lock();
  bool unlock();
  GLuint* getPixelData32() { return _pixels; }
  GLuint getPixel32(GLuint x, GLuint y) { return _pixels[y * _texWidth + x]; }
  void setPixel32(GLuint x, GLuint y, GLuint val) {
    _pixels[y * _texWidth + x] = val;
  }

 private:
  GLuint powerOfTwo(GLuint num);
  GLuint _ID;
  GLuint* _pixels;
  GLuint _texWidth;
  GLuint _texHeight;
  GLuint _imgWidth;
  GLuint _imgHeight;
};



#endif // TEXTURE_HPP

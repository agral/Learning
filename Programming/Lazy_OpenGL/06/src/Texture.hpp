#ifndef TEXTURE_HPP
#define TEXTURE_HPP

#include "LOpenGL.hpp"
#include <iostream>

class Texture
{
 public:
  Texture();
  ~Texture();

  bool loadFromPixels32(GLuint* pixels, GLuint width, GLuint height);
  bool loadFromFile(std::string path);
  void free();
  void render(GLfloat x, GLfloat y);
  GLuint getID() { return _ID; }
  GLuint width() { return _width; }
  GLuint height() { return _height; }

 private:
  GLuint _ID;
  GLuint _width;
  GLuint _height;
};



#endif // TEXTURE_HPP

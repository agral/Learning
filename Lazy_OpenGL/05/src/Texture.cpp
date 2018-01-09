#include "Texture.hpp"

Texture::Texture() :
  _ID(0),
  _width(0),
  _height(0)
{
}

Texture::~Texture()
{
  free();
}


bool Texture::loadFromPixels32(GLuint* pixels, GLuint width, GLuint height)
{
  free(); // Frees the texture if it exists.

  _width = width;
  _height = height;
  glGenTextures(1, &_ID);
  glBindTexture(GL_TEXTURE_2D, _ID);
  glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, width, height, 0, GL_RGBA,
      GL_UNSIGNED_BYTE, pixels);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
  glBindTexture(GL_TEXTURE_2D, 0);

  GLenum error = glGetError();
  if (error != GL_NO_ERROR)
  {
    std::cerr << "Error loading texture from " << pixels << "pixels."
        << std::endl;
    std::cerr << gluErrorString(error) << std::endl;
    return false;
  }

  return true;
}

void Texture::free()
{
  if (_ID != 0)
  {
    glDeleteTextures(1, &_ID);
    _ID = 0;
  }

  _width = 0;
  _height = 0;
}

void Texture::render(GLfloat x, GLfloat y)
{
  if (_ID != 0)
  {
    glLoadIdentity();
    glTranslatef(x, y, 0.f);
    glBindTexture(GL_TEXTURE_2D, _ID);

    // Renders a textured quad:
    glBegin(GL_QUADS);
        glTexCoord2f(0.f, 0.f);
        glVertex2f(0.f, 0.f);
        glTexCoord2f(1.f, 0.f);
        glVertex2f(_width, 0.f);
        glTexCoord2f(1.f, 1.f);
        glVertex2f(_width, _height);
        glTexCoord2f(0.f, 1.f);
        glVertex2f(0.f, _height);
    glEnd();
  }
}

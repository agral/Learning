#include "Texture.hpp"
#include <IL/il.h>

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


bool Texture::loadFromFile(std::string path)
{
  bool isSuccessful = false;
  ILuint imgId = 0;
  ilGenImages(1, &imgId);
  ilBindImage(imgId);
  ILboolean success = ilLoadImage(path.c_str());
  if (success == IL_TRUE)
  {
    success = ilConvertImage(IL_RGBA, IL_UNSIGNED_BYTE);
    if (success == IL_TRUE)
    {
      isSuccessful = loadFromPixels32(
          (GLuint*)ilGetData(),
          (GLuint)ilGetInteger(IL_IMAGE_WIDTH),
          (GLuint)ilGetInteger(IL_IMAGE_HEIGHT)
      );
    }

    ilDeleteImages(1, &imgId);
  }

  if (!isSuccessful)
  {
    std::cout << "Unable to load texture: " << path << std::endl;
  }

  return isSuccessful;
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


void Texture::render(GLfloat x, GLfloat y, Rect<GLfloat>* clip)
{
  if (_ID != 0)
  {
    glLoadIdentity();

    GLfloat top = 0.f;
    GLfloat bottom = 1.f;
    GLfloat left = 0.f;
    GLfloat right = 1.f;
    GLfloat width = _width;
    GLfloat height = _height;

    // Handles the clipping rect, if provided:
    if (clip != nullptr)
    {
      left = clip->x / _width;
      right = (clip->x + clip->w) / _width;
      top = clip->y / _height;
      bottom = (clip->y +clip->h) / _height;

      width = clip->w;
      height = clip->h;
    }

    glTranslatef(x, y, 0.f);
    glBindTexture(GL_TEXTURE_2D, _ID);

    // Renders a textured quad:
    glBegin(GL_QUADS);
        glTexCoord2f(left, top);
        glVertex2f(0.f, 0.f);
        glTexCoord2f(right, top);
        glVertex2f(width, 0.f);
        glTexCoord2f(right, bottom);
        glVertex2f(width, height);
        glTexCoord2f(left, bottom);
        glVertex2f(0.f, height);
    glEnd();
  }
}

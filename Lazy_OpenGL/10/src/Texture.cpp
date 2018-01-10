#include "Texture.hpp"
#include <IL/il.h>
#include <IL/ilu.h>
#include <cstring>

Texture::Texture() :
  _ID(0),
  _pixels(nullptr),
  _texWidth(0),
  _texHeight(0),
  _imgWidth(0),
  _imgHeight(0)
{
}


Texture::~Texture()
{
  free();
}


bool Texture::loadFromPixels32(GLuint* pixels,
    GLuint imgWidth, GLuint imgHeight, GLuint texWidth, GLuint texHeight)
{
  free(); // Frees the texture if it exists.

  _imgWidth = imgWidth;
  _imgHeight = imgHeight;
  _texWidth = texWidth;
  _texHeight = texHeight;

  glGenTextures(1, &_ID);
  glBindTexture(GL_TEXTURE_2D, _ID);
  glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, _texWidth, _texHeight, 0, GL_RGBA,
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
      GLuint imgWidth = (GLuint)ilGetInteger(IL_IMAGE_WIDTH);
      GLuint imgHeight = (GLuint)ilGetInteger(IL_IMAGE_HEIGHT);
      GLuint texWidth = powerOfTwo(imgWidth);
      GLuint texHeight = powerOfTwo(imgHeight);

      if (imgWidth != texWidth || imgHeight != texHeight)
      {
        iluImageParameter(ILU_PLACEMENT, ILU_UPPER_LEFT);
        iluEnlargeCanvas((int)texWidth, (int)texHeight, 1);
      }

      isSuccessful = loadFromPixels32((GLuint*)ilGetData(),
            imgWidth, imgHeight, texWidth, texHeight);
    }

    ilDeleteImages(1, &imgId);
  }

  if (!isSuccessful)
  {
    std::cout << "Unable to load texture: " << path << std::endl;
  }

  return isSuccessful;
}


bool Texture::loadPixelsFromFile(std::string path)
{
  free();
  bool pixelsLoaded = false;

  ILuint id = 0;
  ilGenImages(1, &id);
  ilBindImage(id);

  ILboolean success = ilLoadImage(path.c_str());
  if (success == IL_TRUE)
  {
    success = ilConvertImage(IL_RGBA, IL_UNSIGNED_BYTE);
    if (success == IL_TRUE)
    {
      GLuint imgWidth = (GLuint)ilGetInteger(IL_IMAGE_WIDTH);
      GLuint imgHeight = (GLuint)ilGetInteger(IL_IMAGE_HEIGHT);
      GLuint texWidth = powerOfTwo(imgWidth);
      GLuint texHeight = powerOfTwo(imgHeight);

      if (imgWidth != texWidth || imgHeight != texHeight)
      {
        iluImageParameter(ILU_PLACEMENT, ILU_UPPER_LEFT);
        iluEnlargeCanvas((int)texWidth, (int)texHeight, 1);
      }

      GLuint size = texWidth * texHeight;
      _pixels = new GLuint[size];
      _imgWidth = imgWidth;
      _imgHeight = imgHeight;
      _texWidth = texWidth;
      _texHeight = texHeight;
      memcpy(_pixels, ilGetData(), size*4);
      pixelsLoaded = true;
    }

    ilDeleteImages(1, &id);
  }

  if (!pixelsLoaded)
  {
    std::cerr << "Unable to load: " << path << std::endl;
  }

  return pixelsLoaded;
}


bool Texture::loadTextureFromPixels32()
{
  bool isSuccessful = true;

  if (_ID == 0 && _pixels != nullptr)
  {
    glGenTextures(1, &_ID);
    glBindTexture(GL_TEXTURE_2D, _ID);
    glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, _texWidth, _texHeight,
        0, GL_RGBA, GL_UNSIGNED_BYTE, _pixels);
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
    glBindTexture(GL_TEXTURE_2D, 0);

    GLenum error = glGetError();
    if (error != GL_NO_ERROR)
    {
      std::cerr << "Error loading texture from pixels:" << std::endl;
      std::cerr << gluErrorString(error) << std::endl;
      isSuccessful = false;
    }
    else
    {
      delete [] _pixels;
      _pixels = nullptr;
    }
  }
  else
  {
    std::cerr << "Can not load texture from current pixels." << std::endl;
    if (_ID != 0)
    {
      std::cerr << "A texture is already loaded." << std::endl;
    }
    else if (_pixels == nullptr)
    {
      std::cerr << "There are no pixels to create the texture from." << std::endl;
    }

    isSuccessful = false;
  }

  return isSuccessful;
}


bool Texture::loadFromFileWithColorKey(std::string path, GLuint key)
{
  bool isSuccessful = false;
  if (loadPixelsFromFile(path))
  {
    for (GLuint i = 0; i < _texWidth * _texHeight; ++i)
    {
      if (_pixels[i] == key)
      {
        _pixels[i] = 0x00ffffff; // switches pixel to transparent-white.
      }
    }

    isSuccessful = loadTextureFromPixels32();
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

  if (_pixels != nullptr)
  {
    delete [] _pixels;
    _pixels = nullptr;
  }

  _imgWidth = 0;
  _imgHeight = 0;
  _texWidth = 0;
  _texHeight = 0;
}


void Texture::render(GLfloat x, GLfloat y, Rect<GLfloat>* clip)
{
  if (_ID != 0)
  {
    glLoadIdentity();

    GLfloat texTop = 0.f;
    GLfloat texBottom = (GLfloat)_imgHeight / (GLfloat)_texHeight;
    GLfloat texLeft = 0.f;
    GLfloat texRight = (GLfloat)_imgWidth / (GLfloat)_texWidth;
    GLfloat quadWidth = _imgWidth;
    GLfloat quadHeight = _imgHeight;

    // Handles the clipping rect, if provided:
    if (clip != nullptr)
    {
      texLeft = clip->x / _texWidth;
      texRight = (clip->x + clip->w) / _texWidth;
      texTop = clip->y / _texHeight;
      texBottom = (clip->y +clip->h) / _texHeight;

      quadWidth = clip->w;
      quadHeight = clip->h;
    }

    glTranslatef(x, y, 0.f);
    glBindTexture(GL_TEXTURE_2D, _ID);

    // Renders a textured quad:
    glBegin(GL_QUADS);
        glTexCoord2f(texLeft, texTop);
        glVertex2f(0.f, 0.f);
        glTexCoord2f(texRight, texTop);
        glVertex2f(quadWidth, 0.f);
        glTexCoord2f(texRight, texBottom);
        glVertex2f(quadWidth, quadHeight);
        glTexCoord2f(texLeft, texBottom);
        glVertex2f(0.f, quadHeight);
    glEnd();
  }
}


bool Texture::lock()
{
  bool isSuccessful = false;

  if (_pixels == nullptr && _ID != 0)
  {
    GLuint size = _texWidth * _texHeight;
    _pixels = new GLuint[size];
    glBindTexture(GL_TEXTURE_2D, _ID);
    glGetTexImage(GL_TEXTURE_2D, 0, GL_RGBA, GL_UNSIGNED_BYTE, _pixels);
    glBindTexture(GL_TEXTURE_2D, 0);
    isSuccessful = true;
  }

  return isSuccessful;
}


bool Texture::unlock()
{
  bool isSuccessful = false;

  if (_pixels != nullptr && _ID != 0)
  {
    glBindTexture(GL_TEXTURE_2D, _ID);
    glTexSubImage2D(GL_TEXTURE_2D, 0, 0, 0, _texWidth, _texHeight,
        GL_RGBA, GL_UNSIGNED_BYTE, _pixels);
    delete [] _pixels;
    _pixels = nullptr;
    glBindTexture(GL_TEXTURE_2D, 0);
    isSuccessful = true;
  }

  return isSuccessful;
}


///////////////////////////////////////////////////////////////////////////////
GLuint Texture::powerOfTwo(GLuint num)
{
  if (num != 0)
  {
    --num;
    num |= (num >> 1);
    num |= (num >> 2);
    num |= (num >> 4);
    num |= (num >> 8);
    num |= (num >> 16);
    ++num;
  }
  return num;
}

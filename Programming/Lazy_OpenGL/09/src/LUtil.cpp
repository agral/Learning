#include "LUtil.hpp"
#include "Config/config.hpp"
#include "Texture.hpp"

#include <IL/il.h>
#include <IL/ilu.h>

Texture gTexture;

bool initGL()
{
  glViewport(0.f, 0.f, Config::SCREEN_WIDTH, Config::SCREEN_HEIGHT);

  // Initializes the Projection Matrix:
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  glOrtho(0.0f, Config::SCREEN_WIDTH, Config::SCREEN_HEIGHT, 0.0f, 1.0f, -1.0f);

  // Initializes the ModelView Matrix:
  glMatrixMode(GL_MODELVIEW);
  glLoadIdentity();

  // Initializes the clearing color to full-opaque black:
  glClearColor(0.f, 0.f, 0.f, 1.f);

  // Enable texturing:
  glEnable(GL_TEXTURE_2D);

  GLenum error = glGetError();
  if (error != GL_NO_ERROR)
  {
    std::cerr << "Error while initializing OpenGL." << std::endl;
    std::cerr << gluErrorString(error) << std::endl;
    return false;
  }

  ilInit();
  iluInit();
  ilClearColour(255, 255, 255, 0);

  ILenum ilError = ilGetError();
  if (ilError != IL_NO_ERROR)
  {
    std::cout << "Error while initializing DevIL." << std::endl;
    std::cout << iluErrorString(ilError) << std::endl;
  }

  return true;
}


bool loadMedia()
{
  if (!gTexture.loadFromFile("res/circle.png"))
  {
    std::cerr << "Unable to load texture from file." << std::endl;
    return false;
  }

  gTexture.lock();
  GLuint keyedColor = 0xffff00ff;
  GLuint* pixels = gTexture.getPixelData32();
  for (GLuint i = 0; i < gTexture.texWidth() * gTexture.texHeight(); ++i)
  {
    if (pixels[i] == keyedColor)
    {
      pixels[i] = 0x00000000;
    }
  }

  // Applies diagonal lines:
  for (GLuint y = 0; y < gTexture.imgHeight(); ++y)
  {
    for (GLuint x = 0; x < gTexture.imgWidth(); ++x)
    {
      if (y % 10 != x % 10)
      {
        gTexture.setPixel32(x, y, 0x00000000);
      }
    }
  }
  gTexture.unlock();

  return true;
}

void update()
{
  // purposefully empty - the program displays a static image.
}


void render()
{
  glClear(GL_COLOR_BUFFER_BIT);

  gTexture.render((Config::SCREEN_WIDTH - gTexture.imgWidth()) / 2.f,
      (Config::SCREEN_HEIGHT - gTexture.imgHeight()) / 2.f);

  glutSwapBuffers();
}

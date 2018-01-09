#include "LUtil.hpp"
#include "Config/config.hpp"
#include "Texture.hpp"

Texture gCheckerboardTexture;

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

  return true;
}


bool loadMedia()
{
  const int width = 256;
  const int height = 256;
  const int pixelCount = width * height;
  GLuint checkerboard[pixelCount];

  for (int i = 0; i < pixelCount; ++i)
  {
    if ((i / 256 & 16) ^ (i % 256 & 16))
    {
      checkerboard[i] = 0xFFD28B26; // Solarized-Blue in ABGR hex notation.
    }
    else
    {
      checkerboard[i] = 0xFF164BCB; // Solarized-Orange in ABGR hex notation.
    }
  }
  std::cout << std::endl;

  if (!gCheckerboardTexture.loadFromPixels32(
        checkerboard, width, height))
  {
    std::cerr << "Unable to load the checkerboard texture." << std::endl;
    return false;
  }

  return true;
}

void update()
{
  // purposefully empty - the program displays a static image.
}


void render()
{
  glClear(GL_COLOR_BUFFER_BIT);

  GLfloat x = (Config::SCREEN_WIDTH - gCheckerboardTexture.width()) / 2.f;
  GLfloat y = (Config::SCREEN_HEIGHT - gCheckerboardTexture.height()) / 2.f;

  gCheckerboardTexture.render(x, y);

  glutSwapBuffers();
}

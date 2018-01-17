#include "LUtil.hpp"
#include "Config/config.hpp"
#include "Texture.hpp"

#include <IL/il.h>
#include <IL/ilu.h>

Texture gTexture;
Rect<GLfloat> gStretchRect = {
    0.f, 0.f, Config::SCREEN_WIDTH, Config::SCREEN_HEIGHT};
GLenum gFiltering = GL_LINEAR;

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

  // Enable blending:
  glEnable(GL_BLEND);
  glDisable(GL_DEPTH_TEST);
  glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA);

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
    return false;
  }

  return true;
}


bool loadMedia()
{
  if (!gTexture.loadFromFile("res/opengl_logo.png"))
  {
    std::cerr << "Unable to load texture from file." << std::endl;
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

  gTexture.render(0.f, 0.f, nullptr, &gStretchRect);

  glutSwapBuffers();
}


void handleKeys(unsigned char key, int, int)
{
  if (key == 'q')
  {
    glBindTexture(GL_TEXTURE_2D, gTexture.getID());
    if (gFiltering != GL_LINEAR)
    {
      glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);
      glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
      gFiltering = GL_LINEAR;
    }
    else
    {
      glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST);
      glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
      gFiltering = GL_NEAREST;
    }

    glBindTexture(GL_TEXTURE_2D, 0);
  }
}


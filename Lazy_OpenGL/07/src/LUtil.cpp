#include "LUtil.hpp"
#include "Config/config.hpp"
#include "Texture.hpp"
#include "Rect.hpp"

#include <IL/il.h>
#include <IL/ilu.h>

Texture gSpritesTexture;
Rect<GLfloat> gClips[4];

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
  gClips[0].x = 0.f;
  gClips[0].y = 0.f;
  gClips[0].w = 128.f;
  gClips[0].h = 128.f;

  gClips[1].x = 128.f;
  gClips[1].y = 0.f;
  gClips[1].w = 128.f;
  gClips[1].h = 128.f;

  gClips[2].x = 0.f;
  gClips[2].y = 128.f;
  gClips[2].w = 128.f;
  gClips[2].h = 128.f;

  gClips[3].x = 128.f;
  gClips[3].y = 128.f;
  gClips[3].w = 128.f;
  gClips[3].h = 128.f;


  if (!gSpritesTexture.loadFromFile("res/spritesheet.png"))
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

  gSpritesTexture.render(0.f, 0.f, &gClips[0]);
  gSpritesTexture.render(Config::SCREEN_WIDTH - gClips[1].w, 0.f, &gClips[1]);
  gSpritesTexture.render(0.f, Config::SCREEN_HEIGHT - gClips[2].h, &gClips[2]);
  gSpritesTexture.render(Config::SCREEN_WIDTH - gClips[3].w,
      Config::SCREEN_HEIGHT - gClips[3].h, &gClips[3]);

  glutSwapBuffers();
}

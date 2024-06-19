#include "LUtil.hpp"
#include "Config/config.hpp"

int gColorMode = Config::COLOR_MODE_CYAN;
GLfloat gProjectionScale = 1.f;

bool initGL()
{
  // Initializes the Projection Matrix:
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  glOrtho(0.0f, Config::SCREEN_WIDTH, Config::SCREEN_HEIGHT, 0.0f, 1.0f, -1.0f);

  // Initializes the ModelView Matrix:
  glMatrixMode(GL_MODELVIEW);
  glLoadIdentity();

  // Initializes the clearing color to full-opaque black:
  glClearColor(0.f, 0.f, 0.f, 1.f);

  GLenum error = glGetError();
  if(error != GL_NO_ERROR)
  {
    std::cerr << "Error while initializing OpenGL." << std::endl;
    std::cerr << gluErrorString(error) << std::endl;
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

  // Resets the modelview's matrix:
  glMatrixMode(GL_MODELVIEW);
  glLoadIdentity();

  // Moves to the center of the screen:
  glTranslatef(Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f, 0.f);

  if(gColorMode == Config::COLOR_MODE_CYAN)
  {
    glBegin(GL_QUADS);
        glColor3f(0.f, 1.f, 1.f);
        glVertex2f(-50.f, -50.f);
        glVertex2f( 50.f, -50.f);
        glVertex2f( 50.f,  50.f);
        glVertex2f(-50.f,  50.f);
    glEnd();
  }
  else
  {
    glBegin(GL_QUADS);
        glColor3f(1.f, 0.f, 0.f);
        glVertex2f(-50.f, -50.f);
        glColor3f(1.f, 1.f, 0.f);
        glVertex2f( 50.f, -50.f);
        glColor3f(0.f, 1.f, 0.f);
        glVertex2f( 50.f,  50.f);
        glColor3f(0.f, 0.f, 1.f);
        glVertex2f(-50.f,  50.f);
    glEnd();
  }

  glutSwapBuffers();
}


void handleKeys(unsigned char key, int, int)
{
  if(key == 'q')
  {
    if(gColorMode == Config::COLOR_MODE_CYAN)
    {
      gColorMode = Config::COLOR_MODE_MULTI;
    }
    else
    {
      gColorMode = Config::COLOR_MODE_CYAN;
    }
  }
  else if(key == 'e')
  {
    if(gProjectionScale == 1.f)
    {
      gProjectionScale = 2.f;
    }
    else if(gProjectionScale == 2.f)
    {
      gProjectionScale = 0.5f;
    }
    else if(gProjectionScale == 0.5f)
    {
      gProjectionScale = 1.f;
    }

    glMatrixMode(GL_PROJECTION);
    glLoadIdentity();
    glOrtho(0.0f, Config::SCREEN_WIDTH * gProjectionScale,
        Config::SCREEN_HEIGHT * gProjectionScale, 0.0f, 1.0f, -1.0f );
  }
}

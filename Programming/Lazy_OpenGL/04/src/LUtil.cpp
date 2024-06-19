#include "LUtil.hpp"
#include "Config/config.hpp"

GLfloat gCameraX = 0.f;
GLfloat gCameraY = 0.f;

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

  glPushMatrix();

  // Initializes the clearing color to full-opaque black:
  glClearColor(0.f, 0.f, 0.f, 1.f);

  GLenum error = glGetError();
  if (error != GL_NO_ERROR)
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

  glMatrixMode(GL_MODELVIEW);
  glPopMatrix();
  glPushMatrix();

  glTranslatef(Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f, 0.f);

  // Red quad:
  glBegin(GL_QUADS);
      glColor3f(1.f, 0.f, 0.f);
      glVertex2f(-Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
      glVertex2f( Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
      glVertex2f( Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
      glVertex2f(-Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
  glEnd();

  glTranslatef(Config::SCREEN_WIDTH, 0.f, 0.f);

  // Green quad:
  glBegin(GL_QUADS);
      glColor3f(0.f, 1.f, 0.f);
      glVertex2f(-Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
      glVertex2f( Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
      glVertex2f( Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
      glVertex2f(-Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
  glEnd();

  glTranslatef(0.f, Config::SCREEN_HEIGHT, 0.f);

  // Blue quad:
  glBegin(GL_QUADS);
      glColor3f(0.f, 0.f, 1.f);
      glVertex2f(-Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
      glVertex2f( Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
      glVertex2f( Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
      glVertex2f(-Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
  glEnd();

  glTranslatef(-Config::SCREEN_WIDTH, 0.f, 0.f);

  // Yellow quad:
  glBegin(GL_QUADS);
      glColor3f(1.f, 1.f, 0.f);
      glVertex2f(-Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
      glVertex2f( Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
      glVertex2f( Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
      glVertex2f(-Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
  glEnd();

  glutSwapBuffers();
}


void handleKeys(unsigned char key, int, int)
{
  if (key == 'k')
  {
    gCameraY -= 16.f;
  }
  else if (key == 'j')
  {
    gCameraY += 16.f;
  }
  else if (key == 'h')
  {
    gCameraX -= 16.f;
  }
  else if (key == 'l')
  {
    gCameraX += 16.f;
  }

  glMatrixMode(GL_MODELVIEW);
  glPopMatrix();
  glLoadIdentity();

  glTranslatef(-gCameraX, -gCameraY, 0.f);
  glPushMatrix();
}

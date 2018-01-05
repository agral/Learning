#include "LUtil.hpp"

bool initGL()
{
  // Initializes the Projection Matrix:
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();

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

  // Renders a white rectangle:
  glBegin(GL_QUADS);
      glVertex2f(-0.5f, -0.5f);
      glVertex2f(-0.5f,  0.5f);
      glVertex2f( 0.5f,  0.5f);
      glVertex2f( 0.5f, -0.5f);
  glEnd();

  glutSwapBuffers();
}

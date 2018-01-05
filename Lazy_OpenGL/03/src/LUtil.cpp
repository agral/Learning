#include "LUtil.hpp"
#include "Config/config.hpp"

Config::ViewportMode gViewportMode = Config::ViewportMode::Full;

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
  glLoadIdentity();

  // Moves to the center of the screen:
  glTranslatef(Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f, 0.f);

  // Full-view:
  if (gViewportMode == Config::ViewportMode::Full)
  {
    glViewport(0.f, 0.f, Config::SCREEN_WIDTH, Config::SCREEN_HEIGHT);

    glBegin(GL_QUADS);
        glColor3f(1.f, 0.f, 0.f);
        glVertex2f(-Config::SCREEN_WIDTH / 2.f, -Config::SCREEN_HEIGHT / 2.f);
        glVertex2f( Config::SCREEN_WIDTH / 2.f, -Config::SCREEN_HEIGHT / 2.f);
        glVertex2f( Config::SCREEN_WIDTH / 2.f,  Config::SCREEN_HEIGHT / 2.f);
        glVertex2f(-Config::SCREEN_WIDTH / 2.f,  Config::SCREEN_HEIGHT / 2.f);
    glEnd();
  }
  else if (gViewportMode == Config::ViewportMode::HalfCenter)
  {
    glViewport(Config::SCREEN_WIDTH / 4.f, Config::SCREEN_HEIGHT / 4.f,
        Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f);
    glBegin(GL_QUADS);
        glColor3f(0.f, 1.f, 0.f);
        glVertex2f(-Config::SCREEN_WIDTH / 2.f, -Config::SCREEN_HEIGHT / 2.f);
        glVertex2f( Config::SCREEN_WIDTH / 2.f, -Config::SCREEN_HEIGHT / 2.f);
        glVertex2f( Config::SCREEN_WIDTH / 2.f,  Config::SCREEN_HEIGHT / 2.f);
        glVertex2f(-Config::SCREEN_WIDTH / 2.f,  Config::SCREEN_HEIGHT / 2.f);
    glEnd();
  }
  else if (gViewportMode == Config::ViewportMode::HalfTop)
  {
    glViewport(Config::SCREEN_WIDTH / 4.f, Config::SCREEN_HEIGHT / 2.f,
        Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f);
    glBegin(GL_QUADS);
        glColor3f(0.f, 0.f, 1.f);
        glVertex2f(-Config::SCREEN_WIDTH / 2.f, -Config::SCREEN_HEIGHT / 2.f);
        glVertex2f( Config::SCREEN_WIDTH / 2.f, -Config::SCREEN_HEIGHT / 2.f);
        glVertex2f( Config::SCREEN_WIDTH / 2.f,  Config::SCREEN_HEIGHT / 2.f);
        glVertex2f(-Config::SCREEN_WIDTH / 2.f,  Config::SCREEN_HEIGHT / 2.f);
    glEnd();
  }
  else if (gViewportMode == Config::ViewportMode::Quad)
  {
    // Top-left blue quad:
    glViewport(0.f, Config::SCREEN_HEIGHT / 2.f,
        Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f);
    glBegin(GL_QUADS);
        glColor3f(0.f, 0.f, 1.f);
        glVertex2f(-Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
        glVertex2f( Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
        glVertex2f( Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
        glVertex2f(-Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
    glEnd();

    // Top-right yellow quad:
    glViewport(Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f,
        Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f);
    glBegin(GL_QUADS);
        glColor3f(1.f, 1.f, 0.f);
        glVertex2f(-Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
        glVertex2f( Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
        glVertex2f( Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
        glVertex2f(-Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
    glEnd();

    // Bottom-left red quad:
    glViewport(0.f, 0.f,
        Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f);
    glBegin(GL_QUADS);
        glColor3f(1.f, 0.f, 0.f);
        glVertex2f(-Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
        glVertex2f( Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
        glVertex2f( Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
        glVertex2f(-Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
    glEnd();

    // Bottom-right green quad:
    glViewport(Config::SCREEN_WIDTH / 2.f, 0.f,
        Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f);
    glBegin(GL_QUADS);
        glColor3f(0.f, 1.f, 0.f);
        glVertex2f(-Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
        glVertex2f( Config::SCREEN_WIDTH / 4.f, -Config::SCREEN_HEIGHT / 4.f);
        glVertex2f( Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
        glVertex2f(-Config::SCREEN_WIDTH / 4.f,  Config::SCREEN_HEIGHT / 4.f);
    glEnd();
  }
  else if (gViewportMode == Config::ViewportMode::Radar)
  {
    // Full-sized quad:
    glViewport(0.f, 0.f, Config::SCREEN_WIDTH, Config::SCREEN_HEIGHT);
    glBegin(GL_QUADS);
        glColor3f(1.f, 1.f, 1.f);
        glVertex2f(-Config::SCREEN_WIDTH / 8.f, -Config::SCREEN_HEIGHT / 8.f);
        glVertex2f( Config::SCREEN_WIDTH / 8.f, -Config::SCREEN_HEIGHT / 8.f);
        glVertex2f( Config::SCREEN_WIDTH / 8.f,  Config::SCREEN_HEIGHT / 8.f);
        glVertex2f(-Config::SCREEN_WIDTH / 8.f,  Config::SCREEN_HEIGHT / 8.f);
        glColor3f(0.f, 0.f, 0.f);
        glVertex2f(-Config::SCREEN_WIDTH / 16.f, -Config::SCREEN_HEIGHT / 16.f);
        glVertex2f( Config::SCREEN_WIDTH / 16.f, -Config::SCREEN_HEIGHT / 16.f);
        glVertex2f( Config::SCREEN_WIDTH / 16.f,  Config::SCREEN_HEIGHT / 16.f);
        glVertex2f(-Config::SCREEN_WIDTH / 16.f,  Config::SCREEN_HEIGHT / 16.f);
    glEnd();

    // Radar quad:
    glViewport(Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f,
        Config::SCREEN_WIDTH / 2.f, Config::SCREEN_HEIGHT / 2.f);
    glBegin(GL_QUADS);
        glColor3f(1.f, 1.f, 1.f);
        glVertex2f(-Config::SCREEN_WIDTH / 8.f, -Config::SCREEN_HEIGHT / 8.f);
        glVertex2f( Config::SCREEN_WIDTH / 8.f, -Config::SCREEN_HEIGHT / 8.f);
        glVertex2f( Config::SCREEN_WIDTH / 8.f,  Config::SCREEN_HEIGHT / 8.f);
        glVertex2f(-Config::SCREEN_WIDTH / 8.f,  Config::SCREEN_HEIGHT / 8.f);
        glColor3f(0.f, 0.f, 0.f);
        glVertex2f(-Config::SCREEN_WIDTH / 16.f, -Config::SCREEN_HEIGHT / 16.f);
        glVertex2f( Config::SCREEN_WIDTH / 16.f, -Config::SCREEN_HEIGHT / 16.f);
        glVertex2f( Config::SCREEN_WIDTH / 16.f,  Config::SCREEN_HEIGHT / 16.f);
        glVertex2f(-Config::SCREEN_WIDTH / 16.f,  Config::SCREEN_HEIGHT / 16.f);
    glEnd();
  }


  glutSwapBuffers();
}


void handleKeys(unsigned char key, int, int)
{
  if(key == 'q')
  {
    if (gViewportMode == Config::ViewportMode::Full)
    {
      gViewportMode = Config::ViewportMode::HalfCenter;
    }
    else if (gViewportMode == Config::ViewportMode::HalfCenter)
    {
      gViewportMode = Config::ViewportMode::HalfTop;
    }
    else if (gViewportMode == Config::ViewportMode::HalfTop)
    {
      gViewportMode = Config::ViewportMode::Quad;
    }
    else if (gViewportMode == Config::ViewportMode::Quad)
    {
      gViewportMode = Config::ViewportMode::Radar;
    }
    else
    {
      gViewportMode = Config::ViewportMode::Full;
    }
  }
}

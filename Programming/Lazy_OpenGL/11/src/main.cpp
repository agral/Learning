#include "LUtil.hpp"
#include "Config/config.hpp"
#include "Rect.hpp"

#include <iostream>

void runMainLoop(int val);

int main(int argc, char **argv)
{
  glutInit(&argc, argv);
  glutInitContextVersion(2, 1);
  glutInitDisplayMode(GLUT_DOUBLE);
  glutInitWindowSize(Config::SCREEN_WIDTH, Config::SCREEN_HEIGHT);
  glutCreateWindow(Config::WINDOW_CAPTION);

  if(!initGL())
  {
    std::cerr << "Unable to initialize OpenGL!" << std::endl;
    return 1;
  }

  if (!loadMedia())
  {
    std::cerr << "Unable to load the media." << std::endl;
    return 1;
  }

  glutDisplayFunc(render);
  glutKeyboardFunc(handleKeys);
  glutTimerFunc(1000 / Config::FPS, runMainLoop, 0);
  glutMainLoop();

  return 0;
}

void runMainLoop(int val)
{
  update();
  render();
  glutTimerFunc(1000 / Config::FPS, runMainLoop, val);
}

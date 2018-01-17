#ifndef LUTIL_HPP
#define LUTIL_HPP

#include "LOpenGL.hpp"
#include <iostream>

bool initGL();
bool loadMedia();
void update();
void render();
void handleKeys(unsigned char key, int x, int y);

#endif // LUTIL_HPP


--[[
Assignments work as expected in any language.
Lua allows multiple assignments, where instead of (var) <- value the programmer
may use a list of variables and a list of values.

In a multiple assignment the list of values is first evaluated and only then
the actual assignment takes place.

The multiple assignments are usually used in functions. It is common to see
functions returning more than one argument.
--]]

x, y = 13.17, "Mary had a little lamb"
print( string.format( "Before swapping: x = %s, y = %s", x, y ) )
y, x = x, y
print( string.format( "After swapping: x = %s, y = %s", x, y ) )

print( "a, b, c = 0, 1" )
a, b, c = 0, 1
print( string.format( "After assignment: a = %s, b = %s, c = %s", a, b, c ) )

print( "a, b = 3, 4, 5" )
a, b = 3, 4, 5
print( string.format( "After assignment: a = %s, b = %s", a, b ) )
print( "The redundant value 5 has been ignored." )

print( "Common mistake: assigning a value to several variables." )
a, b, c = 0
print( "a, b, c = 0" )
print( string.format( "After assignment: a = %s, b = %s, c = %s", a, b, c ) )
print( "The correct way to initialize:" )
a, b, c = 0, 0, 0
print( "a, b, c = 0, 0, 0" )
print( string.format( "After assignment: a = %s, b = %s, c = %s", a, b, c ) )

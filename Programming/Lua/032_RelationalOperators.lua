--[[
    The following operators are supported in Lua:
         <    >   <=   >=   ==   ~=
    All these operators always result in either true or false.
    ~= is the negation of equality operator (==).

    If the values have different types, they are considered as different values.
    Otherwise they are compared according to their types.
    Therefore, "10" == 10 yields false.

    nil is equal only to itself.

    Tables, userdata and functions are compared by reference.
    This means that two such values are equal iff they are same object.
--]]


print( string.format( "\"10\" == 10 --> %s", "10" == 10 ) )
a = {}; a.x = 1; a.y = 0
b = {}; b.x = 1; b.y = 0
c = a

print( string.format( "a == b --> %s", a == b ) )
print( string.format( "a == c --> %s", a == c ) )
print( string.format( "b == c --> %s", b == c ) )


-- beware when comparing values with different types:
print( string.format( "2 < 15 --> %s", 2 < 15 ) )
print( string.format( "\"2\" < \"15\"--> %s", "2" < "15" ) )
-- for strings the alphabetical order matters!

print()
print( "Error is raised when strings and numbers are compared." )
print( "Trying to compare: \"2\" < 15 ..." )
print( string.format( "\"2\" < 15--> %s", "2" < 15 ) )

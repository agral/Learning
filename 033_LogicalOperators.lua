--[[
The logical operators are: [and], [or] and [not].
For these logical operators values [nil] and [false] are falsy
and all the other values are truthy (including 0 and "", an empty string).

The operator [and] returns its first argument if it is falsy,
otherwise it returns its second argument.

The operator [or] returns its first argument if it is truthy,
otherwise it returns its second argument.

Both [and] and [or] use short-cut evaluation - that is, they evaluate their
second operand only when necessary.
--]]

local function printAndResult( a, b )
    print( string.format( "%s and %s --> %s", a, b, ( a and b ) ) )
end
local function printOrResult( a, b )
    print( string.format( "%s or %s --> %s", a, b, ( a or b ) ) )
end

printAndResult( 4, 5 )
printAndResult( nil, 13 )
printAndResult( false, 13 )
printAndResult( true, 13 )
printOrResult( 4, 5 )
printOrResult( false, 5 )
printOrResult( 0, 13 )

--[[
USEFUL IDIOMS:
x = x or v    is equivalent to "if not x then x = v"
as in "if x is not already set (or false), sets x to v".

a and b or c: equivalent of C's a ? b : c
for example:
  max = (x > y) and x or y.
This will not work if b is falsy, so make sure that it is truthy.
--]]

print( "x = 5, y = 10, max = ( x > y ) and x or y" )
local x = 5
local y = 10
local max = ( x > y ) and x or y
print( string.format( "--> max is %d", max ) )

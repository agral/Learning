--[[
    Strings are sequence of characters, as usual.
    Any binary data may be stored in a string (they are 8-bits wide)
    Strings are immutable.
--]]
a = "one string"
b = string.gsub( a, "one", "another" )
print( a )
print( b )

print()
print( "Strings may also be delimited by double square brackets [[ ]]" )
s = [[
<html>
  <head>
    <title>Hello</title>
  </head>
  <body>
    <a href="http://tio.run/>TIO</a>
  </body>
</html>
]]
print( s )


-- Lua automatically converts between numbers and strings at runtime.
print( "10" + 1 ) --> 11
print( "10 + 1" ) --> 10 + 1
--print( "lua" + 1 ) --> ERROR: can not convert "lua" to number.

--[[
    Despite those automatic conversions, a comparison between different types
    results in false value, even if values are equal (their types are not).
--]]
print( "10" == 10 ) --> false
print( tonumber("10") == 10 ) --> true
print( "10" == tostring(10) ) --> true

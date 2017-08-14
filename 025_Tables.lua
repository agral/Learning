--[[
    Tables in Lua are neither values nor variables; they are objects.
    They behave very similarly to dicts in Python, in a sense that
    they allow the programmer to store a value under a liberally-chosen key.
    They are also similar to Java arrays (but not e.g. C arrays).
--]]

a = {}             -- this creates a table and stores its reference in `a'
k = "x"
a[ k ] = 10        -- creates new entry: key="x", value=10
a[ 20 ] = "great"  -- creates a new entry, key=20, value="great"

print( a[ "x" ] )
k = 20
print( a[ k ] )
a[ "x" ] = a[ "x" ] + 1
print( a[ "x" ] )

-- It is possible to use a.name instead of a[ "name" ]:
print( a.x )

-- Iterating over the elements of the array:
print()
a = {}
for i = 1,10 do
    a[ i ] = i * 2
end
for i,line in ipairs( a ) do
    print( line )
end
print()

-- It is valid to access a non-initialized variable.
-- A special nil value is returned in such case.
print( "print( b ):" )
print( b )
print()

print( "b = 10" )
b = 10
print( "print( b ):" )
print( b )
print()

-- Usually it is not necessary to delete global variables.
-- In any case, it is possible to delete it by assigning nil to it:
print( "b = nil")
b = nil
print( "print( b ):" )
print( b )
print()

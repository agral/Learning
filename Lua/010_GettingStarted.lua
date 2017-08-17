print( "Hello World" )

-- defines a factorial function
function Factorial( n )
    if n == 0 then
        return 1
    else
        return n * Factorial( n - 1 )
    end
end


print( "Enter a number, please: ")
a = io.read( "*number" )
print( string.format( "%d! is %d", a, Factorial( a ) ) )

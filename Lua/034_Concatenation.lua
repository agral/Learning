--[[
Lua does the string concatenation by [..] operator.
Yes, that is two dots, not "+" sign.
"Hello world " .. "of Lua" is valid, while
"Goodbye cruel " + "world" is not.

Strings in Lua are immutables, so the concatenation operator always creates
a new string, leaving its operands unmodified.
--]]

local printConcat = function( a, b )
    print( string.format( "l=[%s], r=[%s], concat=[%s]", a, b, a .. b ) )
end

printConcat( "Hello world ", "of Lua" )
printConcat( 0, 1 )

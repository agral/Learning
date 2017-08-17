print( [[
Operators precedence in lua is as follows, from higher to lower priority:
+----------+----------------------+
| Priority | Operator(s)          |
+==========+======================+
| 1        | ^                    |
| 2        | not, unary -         |
| 3        | *, /                 |
| 4        | +, -                 |
| 5        | ..                   |
| 6        | <, >, <=, >=, ~=, == |
| 7        | and                  |
| 8        | or                   |
+----------+----------------------+
Operators of equal priority are left-associative,
EXCEPT FOR ^ and .. operators, which are right-associative.
]] )
--[[

This means that the following expressions are equivalent:
    a + i < b / 2 + 1   <==>    ( a + i ) < ( ( b / 2 ) + 1 )
    5 + x ^ 2 / 8       <==>    5 + ( ( x ^ 2 ) / 8 )
    a < y and y <= z    <==>    ( a < y ) and ( y <= z )
    -x ^ 2              <==>    - ( x ^ 2 )
    x ^ y ^ z           <==>    x ^ ( y ^ z )

The following code tests the precedence of operators.
--]]


function printResult( codeAsString )
    local f = assert( loadstring(
            "return function() return " .. codeAsString .. " end" ) )()

    print( string.format( "%s ==> [%s]", codeAsString, f() ) )
end

printResult( "3 + 4 < 5 / 2 + 1" )
printResult( "5 + 4 ^ 2  * 8" )
printResult( "5 < 10 and 10 <= 15" )
printResult( "-3 ^ 2" )
print()
printResult( "4 ^ 3 ^ 2" )
printResult( "( 4 ^ 3 ) ^ 2" )
printResult( "4 ^ ( 3 ^ 2 )" )

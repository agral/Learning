-- If - then - else blocks work as expected. The syntax is straightforward:

local function print_array( arr )
    for _,v in ipairs( arr ) do
        io.write( v )
        io.write("\t")
    end
    print()
end

local numbers = { 3, -1, 4, -1, 5, -9, 2 }
print( "The number array: ")
print_array( numbers )

for i,v in ipairs( numbers ) do
    if v < 0 then
        numbers[ i ] = 0
    end
end
print( "After converting negative numbers to 0:" )
print_array( numbers )

-- elseif keyword is used for multiple conditions checking:
local function calculate( operand, lhs, rhs )
    local result
    if operand == "+" then
        result = lhs + rhs
    elseif operand == "-" then
        result = lhs - rhs
    elseif operand == "*" then
        result = lhs * rhs
    elseif operand == "/" then
        result = lhs / rhs
    else
        result = string.format( "Error: Unknown operand \"%s\".", operand )
    end

    return result
end

local function calculate_and_print( operand, lhs, rhs )
    result = calculate( operand, lhs, rhs )
    print( string.format( "%s %s %s == %s", lhs, operand, rhs, result ) )
end

calculate_and_print( "*", 6, 7 )
calculate_and_print( "/", 65536, 256 )
calculate_and_print( "garbage", 11, 22 )


-------------------------------------------------------------------------------
-- while loop: the condition is evaluated and if it is true, the body
-- of the loop is executed repeatedly until it turns false.
local i = 1 -- remember, the arrays in Lua are 1-indexed.
local accumulator = 0
while numbers[ i ] do
    accumulator = accumulator + numbers[ i ]
    print( string.format(
            "numbers[ %d ] is %d, and the accumulated sum is %d.",
            i, numbers[ i ], accumulator
    ) )
    i = i + 1
end

-------------------------------------------------------------------------------
-- repeat-until loop does exactly the same, except that the test is done
-- after the body, so the body executes at least once:
local input
repeat
    print( "To go further, write: :q! below." )
    input = io.read()
until input == ":q!"
print( "That's how you exit vim!" )

-------------------------------------------------------------------------------
-- numeric for-loop has the following syntax:
--     for var = v1, v2, step do
--         instruction(s)
--     end
-- All the three expressions are evaluated once (before the loop starts).
-- The control variable var is declared automatically as local one
-- and it DOES NOT exist outside the loop body.
-- Also, the value of control variable var should not be modified
-- by a programmer. This results in an "unpredictable" (undefined?) behavior.
-- As every other loop, it may be break'ed and continue'ed.
print( "x\tx^2\tx^3\tx^4" )
print( "-----------------------------" )
for i = 10, 1, -1 do
    print( string.format( "%d\t%d\t%d\t%d", i, i*i, i*i*i, i*i*i*i ) )
end

print( "Here is why you should not use the for-loop control variable:" )
for i = 1, 10 do print( i ) end
max = i; print( string.format( "%s", max ) ) -- oops, max is a global 'i'!

print( "Here is what it should look like - finding a value in a list:" )
local haystack = { 2, 7, 1, 8, 2, 8, 1, 8, 4, 5, 9, 0, 4, 5 }
print_array( haystack )

local function linear_search( needle, haystack )
    local found = nil
    for i = 1, #haystack do -- by the way, # is unary operator "length-of".
        if haystack[ i ] == needle then
            found = i
            break
        end
    end

    return found
end

local needles = { 3, 5, 7 }
for i = 1, #needles do
    print( string.format( "linear_search(%d, haystack) --> %s",
            needles[ i ], linear_search( needles[ i ], haystack )
    ) )
end

-------------------------------------------------------------------------------
-- generic for-loop allows traversing all the values returned by an iterator
-- (iterator function). It utilizes pairs() or ipairs() to iterate:
--     for i, val in ipairs( table ) do instructions end
--     for k in pairs( table ) do instructions end
-- In the first example: i gets an index; val gets the value under this index.
-- In the second example, k holds the keys of the table. No order is guaranteed.

local days = { "Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
        "Saturday", "Sunday" }
print( "Days of the week:" )
print_array( days )
local revDays = {}
for i, v in ipairs( days ) do
    revDays[ v ] = i
end
print( "Reverse-days of the week:" )
for k in pairs( revDays ) do
    print( string.format( "%s -> %s", k, revDays[ k ] ) )
end
print( string.format( "Third day of the week is %s", days[ 3 ] ) )
print( string.format( "And Friday is %d. day of a week.",
        revDays["Friday"] ) )

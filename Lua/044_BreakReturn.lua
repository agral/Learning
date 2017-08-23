-- break and return instructions allow to jump out from an enclosing block.
-- break instruction breaks execution of an inner loop
-- (for, repeat and while-loops are valid). It can NOT be used outside of these
-- loops. The program will continue execution immediately after the broken loop.
--
-- return statement returns (optional) results from a function.
-- Every lua function is implicitly terminated by an empty return statement,
-- so there is no need to explicitly using it when a function does not return
-- any meaningful value.
--
-- Attention: break and return instructions can only appear as a last statement
-- of a block - that is, just before an end/else/until statement.
-- A bit of a hack allows you to use it anywhere you want:
-- enclose a return/break instruction in a do..end block, like this:

print( "Breaking in the middle of a chunk!" )
local i = 1
while i < 10 do
    print( string.format( "[%d], header", i ) )

    -- the docs state that this should not compile (runtime error).
    -- It works nevertheless!
    if i == 3 then
        break
        print( "You'll never reach me!" )
    end

    print( string.format( "[%d], body", i ) )
    print( string.format( "[%d], footer", i ) )

    i = i + 1
end


-- return'ing from a function in a middle of its body is disallowed:
local function foo()
    print( "foo() begins..." )


    print( "foo() is running..." )

    -- return -- invalid, return must be at the end of a function!
    do return end -- perfectly valid, return stays right before end ;-)

    print( "foo() will return after this statement!" )
    return 42
end

foo()

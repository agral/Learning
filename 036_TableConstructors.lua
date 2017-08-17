--[[
The tables are initialized by a table constructor, {}.
The simplest constructor is an empty constructor, {}.
It is possible to initialize arrays by passing things inside curly braces.
Example:
]]--

function print_table( table )
    print()
    print( "Index\t|\tValue\t\t|\tType" )
    print( "-------------------------------------------------" )
    for n,val in pairs( table ) do
        print( string.format( "%s\t|\t%s\t\t|\t%s", n, val, type( val ) ) )
    end
end

weekdays = { "Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
             "Saturday", "Sunday" }

--Tables are 1-indexed (and not 0-indexed). Thus, weekdays[ 4 ]
--returns Thursday, as one would expect.

print( string.format( "The 4th day of the week is %s", weekdays[ 4 ] ) )

--The following syntax is used in order to initialize a table
--that will be used as a record:
-- concise syntax:
r1 = { x = 3, y = 4 }

-- equivalent syntax:
r2 = {}
r2.x = 3
r2.y = 4

-- Adding and removing fields is always possible, and instances of different
-- types may freely coexist in any array:
w = { x = 0, y = 0, label = "Console" }
x = { math.sin( 0 ), math.sin( 30 * math.pi / 180 ),
        math.sin( 60 * math.pi / 180) }

print_table( w )

-- Comma after last entry in constructor is valid:
arr = { x = 10, y = 45, "one", "two", "three", }
print_table( arr )

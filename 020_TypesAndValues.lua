function PrintType( val )
    print( string.format( "type( %s ) --> %s", val, type( val ) ) )
end

PrintType( "Hello world" )
PrintType( 10.4 * 3 )
PrintType( print )
PrintType( type )
PrintType( nil )

-- result of type() is always a string
PrintType( type( type( X ) ) )


function AssignAndPrint( val )
    print( string.format( "val is %s, type( val ) is %s", val, type( val ) ) )
    print( "Assigning val to a")
    local a = val
    PrintType( a )
    print()
end


AssignAndPrint( 10 )
AssignAndPrint( "We the people" )
AssignAndPrint( print )

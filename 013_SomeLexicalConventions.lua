--[[
    Identifiers in Lua can be any string composed of:
      - letters       [A-Za-z]
      - digits        [0-9]
      - underscores   _
    provided that it does not begin with a digit.

    For instance, the following identifiers are valid:
        i    j    i10    _ij
        aSomewhatLongName    _INPUT

    The concept of what a letter is in Lua is locale-dependent.
    This means that with a proper locale some exotic identifiers
    could be used (such as índice or ação).

    DO NOT DO IT.

    Stick to the standard [A-Za-z] range, guaranteed to be supported
    in any locale.
--]]

--[[
    The following words are reserved - they can not be used as identifiers:
    and       break     do        else      elseif
    end       false     for       function  if
    in        local     nil       not       or
    repeat    return    then      true      until     while

    -- for example, the following chunk is not valid:
    nil = 5

    However, Lua is case-sensitive.
    While and is a reserved word, And & AND are two other
    different identifiers.
--]]

print( "Nil = 4; NIL = 2" )
Nil = 4; NIL = 2
print( "print( Nil, NIL, nil ) )" )
print( Nil, NIL, nil )

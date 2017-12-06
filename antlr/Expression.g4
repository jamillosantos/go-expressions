grammar Expression;

expression
   : multiplyingExpression ((PLUS | MINUS) multiplyingExpression)*
   ;

multiplyingExpression
   : powExpression ((TIMES | DIV | MOD) powExpression)*
   ;

powExpression
   : signedAtom (POW signedAtom)*
   ;

signedAtom
   : operator=(PLUS | MINUS) signedAtom
   | function
   | atom
   | binaryOp
   ;

binaryOp
   : atom relop atom
   ;

atom
   : scientific
   | variable
   | constant
   | LPAREN expression RPAREN
   | str
   ;

str
   : QUOTED_STRING
   ;

scientific
   : SCIENTIFIC_NUMBER
   ;

constant
   : PI
   | EULER
   | I
   ;

variable
   : VARIABLE
   ;

function
   : fname=VARIABLE LPAREN expression (COMMA expression)* RPAREN
   ;

relop
   : EQ
   | NOT_EQ
   | GT
   | LT
   | OR
   | AND
   | XOR
   ;

LPAREN
   : '('
   ;


RPAREN
   : ')'
   ;


PLUS
   : '+'
   ;


MINUS
   : '-'
   ;


TIMES
   : '*'
   ;


DIV
   : '/'
   ;

MOD
   : '%'
   ;


GT
   : '>'
   ;


LT
   : '<'
   ;


EQ
   : '=='
   ;

NOT_EQ
   : '!='
   ;

OR
   : '||'
   ;

AND
   : '&&'
   ;

XOR
   : 'xor'
   ;


COMMA
   : ','
   ;


POINT
   : '.'
   ;


POW
   : '^'
   ;


PI
   : 'pi'
   ;


EULER
   : E2
   ;


I
   : 'i'
   ;


VARIABLE
   : VALID_ID_START VALID_ID_CHAR*
   ;

QUOTED_STRING
   : QUOTE ( ESCAPED_QUOTE | ~('\n'|'\r') )*? QUOTE
   ;

QUOTE
   : '"'
   ;

fragment ESCAPED_QUOTE
   : '\\"'
   ;

fragment VALID_ID_START
   : ('a' .. 'z') | ('A' .. 'Z') | '_'
   ;


fragment VALID_ID_CHAR
   : VALID_ID_START | ('0' .. '9')
   ;


SCIENTIFIC_NUMBER
   : NUMBER ((E1 | E2) SIGN? NUMBER)?
   ;


fragment NUMBER
   : ('0' .. '9') + ('.' ('0' .. '9') +)?
   ;


fragment E1
   : 'E'
   ;


fragment E2
   : 'e'
   ;


fragment SIGN
   : ('+' | '-')
   ;


WS
   : [ \r\n\t] + -> skip
   ;

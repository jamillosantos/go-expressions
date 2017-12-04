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
   | binaryOp
   | atom
   ;

binaryOp
   : atom relop atom
   ;

atom
   : scientific
   | variable
   | constant
   | LPAREN expression RPAREN
   | string
   ;

string
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
   : funcname LPAREN expression (COMMA expression)* RPAREN
   ;

funcname
   : COS
   | TAN
   | SIN
   | ACOS
   | ATAN
   | ASIN
   | LOG
   | LN
   | SQRT
   | IF
   ;

relop
   : EQ
   | GT
   | LT
   | OR
   | AND
   | XOR
   ;


COS
   : 'cos'
   ;


SIN
   : 'sin'
   ;


TAN
   : 'tan'
   ;


ACOS
   : 'acos'
   ;


ASIN
   : 'asin'
   ;


ATAN
   : 'atan'
   ;


LN
   : 'ln'
   ;


LOG
   : 'log'
   ;


SQRT
   : 'sqrt'
   ;

IF
   : 'if'
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

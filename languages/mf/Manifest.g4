grammar Manifest;

mf: (section)* EOF;

section
  : Key COLON SPACE value
  ;

value
 : OTHERS (SEMI VERSION ASSIGN STRING_LITERAL)? (SEMI configAssign)*
 ;

configAssign
  : assignKey SEQUAL assignValue
  ;

assignKey: OTHERS;
assignValue: OTHERS;

VERSION: 'version';

Key: 'Manifest-Version'
  | 'Bundle-Activator'
  | 'Created-By'
  | 'Import-Package'
  | 'Export-Package'
  | 'Bundle-Version'
  | 'Bundle-Name'
  | 'Bundle-Description'
  | 'Bundle-DocURL'
  | 'Bundle-ManifestVersion'
  | 'Bundle-Vendor'
  | 'Bundle-SymbolicName'
  | 'Implementation-Version'
  | 'Implementation-Title'
  | 'Ant-Version'
  | 'Spring-Version'
  | Uppercase Letter* '-' Uppercase Letter*
  ;

//IDENTIFIER: Letter LetterOrDigit*;
OTHERS:  ValueText (SPACE? ValueText)*;
ValueText
  : ('(' | '.' |')' | '-'| LetterOrDigit+)+
//  | Version
  ;

//Version
//  : Digits ('.' LetterOrDigit)+
//  ;


// Separators

COLON:              ':';
LPAREN:             '(';
RPAREN:             ')';
LBRACE:             '{';
RBRACE:             '}';
LBRACK:             '[';
RBRACK:             ']';
SEMI:               ';';
COMMA:              ',';
DOT:                '.';
ASSIGN:             '=';
GT:                 '>';
LT:                 '<';
BANG:               '!';
TILDE:              '~';
QUESTION:           '?';
EQUAL:              '==';
LE:                 '<=';
GE:                 '>=';
NOTEQUAL:           '!=';
AND:                '&&';
OR:                 '||';
INC:                '++';
DEC:                '--';
ADD:                '+';
SUB:                '-';
MUL:                '*';
DIV:                '/';
BITAND:             '&';
BITOR:              '|';
CARET:              '^';
MOD:                '%';
DQUOTE:             '"';
SEQUAL:             ':=';

Uppercase:          [A-Z];

//LINE_COMMENT :       ';' ~('\n'|'\r')*  ->  channel(HIDDEN);
SPACE:               [ \t];
NL :                 ('\r'? '\n' ' '? ) -> skip;

STRING_LITERAL:     '"' (~["\\\r\n] | EscapeSequence)* '"';

fragment EscapeSequence
    : '\\' [btnfr"'\\]
    | '\\' ([0-3]? [0-7])? [0-7]
    | '\\' 'u'+ HexDigit HexDigit HexDigit HexDigit
    ;
fragment HexDigit
    : [0-9a-fA-F]
    ;

fragment Letter: [a-zA-Z];
fragment LetterOrDigit
    : Letter
    | [0-9]
    ;
fragment Digits
    : [0-9] ([0-9_]* [0-9])?
    ;

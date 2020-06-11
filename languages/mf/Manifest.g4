grammar Manifest;

mf: (section)* EOF;

section
//  : 'SHA1-Digest' COLON SPACE ValueText
  : Key COLON SPACE value '='?
  ;

value
 : pkg (COMMA pkg)*
 | STRING_LITERAL
 ;

pkg
  : OTHERS (SEMI configAssign)*
  ;

configAssign: assignKey (SEQUAL | ASSIGN) assignValue;

assignKey: OTHERS;
assignValue: STRING_LITERAL | OTHERS;

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
  | 'Name'
  | 'Can-Redefine-Classes'
  | 'Package'
  | 'SHA1-Digest'
  | 'ant' '-' Uppercase Letter* ('-' Uppercase Letter*)*
  | Uppercase Letter* '-' Uppercase Letter* ('-' Uppercase Letter*)*
  ;

OTHERS:  ValueText (SPACE? ValueText)*;
ValueText
  : ('(' | '.' |')' | '-' | '$' | '_' | '/' | '%' | '+' | LetterOrDigit+)+
  ;

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

SPACE:               [ \t];
NL :                 ('\r'? '\n' ' '? ) -> skip;

STRING_LITERAL:     '"' (~["\\\r\n] | EscapeSequence)* '"';
IDENTIFIER: Letter LetterOrDigit*;

fragment EscapeSequence
    : '\\' [btnfr"'\\]
    | '\\' ([0-3]? [0-7])? [0-7]
    | '\\' 'u'+ HexDigit HexDigit HexDigit HexDigit
    | '/' | '+'
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

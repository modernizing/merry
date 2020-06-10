grammar Manifest;

mf: (LINE_COMMENT | section)* EOF;

section
  : Key Colon SPACE? value
  ;

value
 : ValueText (ValueText | SPACE)*
 ;

Key: Uppercase Letter* '-' Uppercase Letter*
  ;

ValueText
  :  ('.' | LetterOrDigit+)+
  | Symbol
  ;

QualifiedName
    : IDENTIFIER ('.' IDENTIFIER)+
    ;

IDENTIFIER: Letter LetterOrDigit*;

fragment Letter: [a-zA-Z];
fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

Colon:         ':';

Uppercase:     [A-Z];
Symbol:        '(' | ')' | '-' | '=' | '[' | ']' | '"' | ',' | ';';

LINE_COMMENT : ';' ~('\n'|'\r')*  ->  channel(HIDDEN);
SPACE:         [ \t];
NL :           '\r'? '\n' -> skip;

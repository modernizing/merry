grammar Manifest;

mf: (LINE_COMMENT | section)* EOF;

section
  : Key Colon SPACE? value
  ;

value
 : ValueText (ValueText | SPACE)*
 ;

Key
  : IsImport
  | IsExport
  | Uppercase Letter* '-' Uppercase Letter*
  ;

IsImport: 'Import-Package';
IsExport: 'Export-Package';

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

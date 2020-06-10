grammar Manifest;

mf: (LINE_COMMENT | section)* EOF;

section
  : isImport Colon Space? value
  | Key Colon Space? value
  ;

isImport: 'Import-Package';

value: ValueText (ValueText | Space)*;

Key: Uppercase Letter* '-' Uppercase Letter*;

ValueText
  : LetterOrDigit
  | Symbol
  ;

fragment Letter: [a-zA-Z];
fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

Uppercase: [A-Z];
Lowercase: [a-z];

Symbol: '(' | ')' | '.' | '-' | ';' | '=' | '[' | ']' | '"' | ',';

Colon: ':';

LINE_COMMENT : ';' ~('\n'|'\r')*  ->  channel(HIDDEN);

Space:  [ \t];

NewLine : '\r'? '\n' -> skip;

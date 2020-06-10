grammar Manifest;

mf: (LINE_COMMENT | section)* EOF;

section: Key Colon Space? value;

value: ValueText (ValueText | Space)*;

Key: Uppercase (Letter | Symbol)*;

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

Symbol: '(' | ')' | '.' | '-';

Colon: ':';

LINE_COMMENT : ';' ~('\n'|'\r')*  ->  channel(HIDDEN);

Space:  [ \t];

NewLine : '\r'? '\n' -> skip;

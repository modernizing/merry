grammar Manifest;

mf: (LINE_COMMENT | section)* EOF;

section: Key Colon value;

value
  : Value
  ;

Value : ([a-z] | [A-Z] | [0-9] | '(' | ')' | '.' | '*')+;

Key: [A-Z] ([a-z] | '-' | [A-Z])+;

Colon: ':';

LINE_COMMENT : ';' ~('\n'|'\r')*  ->  channel(HIDDEN);

WS  :   (('\r')? '\n' |  ' ' | '\t')+  -> skip;


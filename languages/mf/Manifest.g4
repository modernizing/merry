grammar Manifest;

mf: (LINE_COMMENT | section)* EOF;

section: START_HEAD Colon key_values;

key_values
  : Version
  ;

Version : [0-9]+ ('.' | [0-9]+)+;

START_HEAD: [A-Z] ([a-z] | '-' | [A-Z])+;

HEAD_TEXT: 'a' .. 'z' | '-';

Colon: ':';

LINE_COMMENT : ';' ~('\n'|'\r')*  ->  channel(HIDDEN);

WS  :   (('\r')? '\n' |  ' ' | '\t')+  -> skip;


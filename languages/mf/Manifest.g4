grammar Manifest;

mf: (LINE_COMMENT | section)* EOF;

section : section_header ':' key_values;

section_header: START_HEAD HEAD_TEXT+;

key_values: TEXT;

START_HEAD: 'A' .. 'Z';

HEAD_TEXT: 'a' .. 'z' | '-';

TEXT: ( 'a' .. 'z' | 'A' .. 'Z' | '_' | '0' .. '9' | '/' | '\\' | ':' | '*' | '.' | ',' | '@' | ' ')+;

LINE_COMMENT : ';' ~('\n'|'\r')*  ->  channel(HIDDEN);

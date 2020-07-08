%{
package parser

import (
	"fmt"
)

func (l *lexer) Lex(lval *yySymType) int {
	const (
		S_INIT = iota
		S_COMMENTS
	)
	
	c := l.current
	currentState := 0

	if l.empty {
		c, l.empty = l.getc(), false
	}

%}

D                       [0-9]
L                       [a-zA-Z_:]

%x S_COMMENTS
%yyc c
%yyn c = l.getc()
%yyt currentState

%%
	l.buf.Reset()

\/\/[^\r\n]*\n           /* single-line comments */

"/*" 					currentState = S_COMMENTS
<S_COMMENTS>"*/" 		currentState = S_INIT
<S_COMMENTS>.|\n         /* ignore chars within multi-line comments */

[\t ]                /* whitespace */
\|[\t \n]*

\-> lval.str = l.token(); return RE_OP
\<\- lval.str = l.token(); return LE_OP
\-\- lval.str = l.token(); return UE_OP
, lval.str = l.token(); return COMMA
= lval.str = l.token(); return EQ
# lval.str = l.token(); return HASH


[\n\r] return NEWLINE;
\0    return EOF;

{L}({L}|{D})* lval.str = l.token(); return STRING
\"{L}({L}|{D}|[\t ])*\" lval.str = l.token()[1:len(l.token())-1]; return QSTRING
`[^`]*` lval.str = l.token()[1:len(l.token())-1]; return BTICKSTR
%%

  l.empty = true
  return int(c)
}




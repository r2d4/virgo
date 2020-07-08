%{
package parser

//import "fmt"
import "matt-rickard.com/virgo/pkg/types/v1"

// func init() {
//   yyDebug = 99
// }

%}

%union{
  str string
  str_list []string
  edge_stmt v1.EdgeStmt
  d_list []v1.Directive
  kv_list []v1.KV
}

%token<str> 
  LE_OP
  UE_OP
  RE_OP
  COMMA
  QUOTE
  NEWLINE
  STRING "string"
  QSTRING
  EQ
  LINE
  BTICKSTR
  EOF
  PIPE
  HASH

%type<str> op value
%type<edge_stmt> edge_stmt
%type<str_list> node_name_list
%type<d_list> directive_stmt
%type<kv_list> kv_list

%start start

%%

start: stmt
  ;

stmt: stmt stmt
  | stmt_start stmt stmt_end
  | vertex_stmt
  | edge_stmt
  | directive_stmt
  ;

vertex_stmt: 
  value EQ BTICKSTR { yylex.(*lexer).AddVertexDef($1, $3) }
  | value EQ value
  { 
    yylex.(*lexer).graph.Vertices[$1] = yylex.(*lexer).graph.Vertices[$3]
  }
  ;

edge_stmt: {}
  | node_name_list op node_name_list 
  {
    $$ = v1.EdgeStmt{LHS:$1, Op: $2, RHS: $3}
    yylex.(*lexer).AddEdges($$.Op, $1, $3)
  }
  | edge_stmt op node_name_list 
  {
    $$ = v1.EdgeStmt{LHS:$1.RHS, Op: $2, RHS:$3}
    yylex.(*lexer).AddEdges($$.Op, $1.RHS, $3)
  }
  ;

directive_stmt: HASH STRING kv_list
  {
    yylex.(*lexer).graph.Directives = append(yylex.(*lexer).graph.Directives, v1.Directive{$2, $3})
  }
  | HASH STRING
  {
      yylex.(*lexer).graph.Directives = append(yylex.(*lexer).graph.Directives, v1.Directive{Name: $2})
  }
  ;

kv_list: kv_list kv_list { $$ = append($1, $2...)}
  | value EQ value { $$ = []v1.KV{{$1, $3}} }


node_name_list:
  node_name_list node_name_list { $$ = append($1, $2...) }
  | node_name_list COMMA value { $$ = append($1, $3) }
  | value { $$ = []string{$1} }
  ;

value: STRING | QSTRING;
op: RE_OP | LE_OP | UE_OP;
stmt_start: /* empty */ {} | NEWLINE;
stmt_end: NEWLINE | EOF { return 1 };
// Code generated by goyacc - DO NOT EDIT.

package parser

import __yyfmt__ "fmt"

//import "fmt"
import "matt-rickard.com/virgo/pkg/types/v1"

// func init() {
//   yyDebug = 99
// }

type yySymType struct {
	yys       int
	str       string
	str_list  []string
	edge_stmt v1.EdgeStmt
	d_list    []v1.Directive
	kv_list   []v1.KV
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57360
	yyEofCode = 57344
	BTICKSTR  = 57356
	COMMA     = 57349
	EOF       = 57357
	EQ        = 57354
	HASH      = 57359
	LE_OP     = 57346
	LINE      = 57355
	NEWLINE   = 57351
	PIPE      = 57358
	QSTRING   = 57353
	QUOTE     = 57350
	RE_OP     = 57348
	STRING    = 57352
	UE_OP     = 57347
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -28
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		57352: 0,  // STRING (36x)
		57353: 1,  // QSTRING (35x)
		57346: 2,  // LE_OP (27x)
		57348: 3,  // RE_OP (27x)
		57347: 4,  // UE_OP (27x)
		57357: 5,  // EOF (25x)
		57359: 6,  // HASH (25x)
		57351: 7,  // NEWLINE (25x)
		57344: 8,  // $end (23x)
		57370: 9,  // value (17x)
		57364: 10, // node_name_list (11x)
		57349: 11, // COMMA (9x)
		57361: 12, // directive_stmt (5x)
		57362: 13, // edge_stmt (5x)
		57367: 14, // stmt (5x)
		57369: 15, // stmt_start (5x)
		57371: 16, // vertex_stmt (5x)
		57354: 17, // EQ (4x)
		57363: 18, // kv_list (3x)
		57365: 19, // op (2x)
		57356: 20, // BTICKSTR (1x)
		57366: 21, // start (1x)
		57368: 22, // stmt_end (1x)
		57360: 23, // $default (0x)
		57345: 24, // error (0x)
		57355: 25, // LINE (0x)
		57358: 26, // PIPE (0x)
		57350: 27, // QUOTE (0x)
	}

	yySymNames = []string{
		"STRING",
		"QSTRING",
		"LE_OP",
		"RE_OP",
		"UE_OP",
		"EOF",
		"HASH",
		"NEWLINE",
		"$end",
		"value",
		"node_name_list",
		"COMMA",
		"directive_stmt",
		"edge_stmt",
		"stmt",
		"stmt_start",
		"vertex_stmt",
		"EQ",
		"kv_list",
		"op",
		"BTICKSTR",
		"start",
		"stmt_end",
		"$default",
		"error",
		"LINE",
		"PIPE",
		"QUOTE",
	}

	yyTokenLiteralStrings = map[int]string{
		57352: "string",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {21, 1},
		2:  {14, 2},
		3:  {14, 3},
		4:  {14, 1},
		5:  {14, 1},
		6:  {14, 1},
		7:  {16, 3},
		8:  {16, 3},
		9:  {13, 0},
		10: {13, 3},
		11: {13, 3},
		12: {12, 3},
		13: {12, 2},
		14: {18, 2},
		15: {18, 3},
		16: {10, 2},
		17: {10, 3},
		18: {10, 1},
		19: {9, 1},
		20: {9, 1},
		21: {19, 1},
		22: {19, 1},
		23: {19, 1},
		24: {15, 0},
		25: {15, 1},
		26: {22, 1},
		27: {22, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [38][]uint8{
		// 0
		{38, 39, 19, 19, 19, 4, 37, 40, 19, 35, 36, 12: 34, 33, 30, 31, 32, 21: 29},
		{8: 28},
		{38, 39, 19, 19, 19, 4, 37, 40, 27, 35, 36, 12: 34, 33, 62, 31, 32},
		{38, 39, 19, 19, 19, 19, 37, 40, 9: 35, 36, 12: 34, 33, 61, 31, 32},
		{24, 24, 24, 24, 24, 24, 24, 24, 24},
		// 5
		{23, 23, 52, 51, 53, 23, 23, 23, 23, 19: 59},
		{22, 22, 22, 22, 22, 22, 22, 22, 22},
		{10, 10, 10, 10, 10, 11: 10, 17: 56},
		{38, 39, 52, 51, 53, 9: 50, 48, 49, 19: 47},
		{41},
		// 10
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 11: 9, 17: 9},
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 11: 8, 17: 8},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{38, 39, 15, 15, 15, 15, 15, 15, 15, 43, 18: 42},
		{38, 39, 16, 16, 16, 16, 16, 16, 16, 43, 18: 46},
		// 15
		{17: 44},
		{38, 39, 9: 45},
		{13, 13, 13, 13, 13, 13, 13, 13, 13},
		{38, 39, 14, 14, 14, 14, 14, 14, 14, 43, 18: 46},
		{38, 39, 9: 50, 55},
		// 20
		{38, 39, 12, 12, 12, 12, 12, 12, 12, 50, 48, 49},
		{38, 39, 9: 54},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 11: 10},
		{7, 7},
		{6, 6},
		// 25
		{5, 5},
		{11, 11, 11, 11, 11, 11, 11, 11, 11, 11: 11},
		{38, 39, 18, 18, 18, 18, 18, 18, 18, 50, 48, 49},
		{38, 39, 9: 58, 20: 57},
		{21, 21, 21, 21, 21, 21, 21, 21, 21},
		// 30
		{20, 20, 20, 20, 20, 20, 20, 20, 20},
		{38, 39, 9: 50, 60},
		{38, 39, 17, 17, 17, 17, 17, 17, 17, 50, 48, 49},
		{38, 39, 19, 19, 19, 65, 37, 64, 9: 35, 36, 12: 34, 33, 62, 31, 32, 22: 63},
		{38, 39, 26, 26, 26, 26, 37, 40, 26, 35, 36, 12: 34, 33, 62, 31, 32},
		// 35
		{25, 25, 25, 25, 25, 25, 25, 25, 25},
		{3, 3, 3, 3, 3, 3, 3, 3, 2},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 24

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 7:
		{
			yylex.(*lexer).AddVertexDef(yyS[yypt-2].str, yyS[yypt-0].str)
		}
	case 8:
		{
			yylex.(*lexer).graph.Vertices[yyS[yypt-2].str] = yylex.(*lexer).graph.Vertices[yyS[yypt-0].str]
		}
	case 10:
		{
			yyVAL.edge_stmt = v1.EdgeStmt{LHS: yyS[yypt-2].str_list, Op: yyS[yypt-1].str, RHS: yyS[yypt-0].str_list}
			yylex.(*lexer).AddEdges(yyVAL.edge_stmt.Op, yyS[yypt-2].str_list, yyS[yypt-0].str_list)
		}
	case 11:
		{
			yyVAL.edge_stmt = v1.EdgeStmt{LHS: yyS[yypt-2].edge_stmt.RHS, Op: yyS[yypt-1].str, RHS: yyS[yypt-0].str_list}
			yylex.(*lexer).AddEdges(yyVAL.edge_stmt.Op, yyS[yypt-2].edge_stmt.RHS, yyS[yypt-0].str_list)
		}
	case 12:
		{
			yylex.(*lexer).graph.Directives = append(yylex.(*lexer).graph.Directives, v1.Directive{yyS[yypt-1].str, yyS[yypt-0].kv_list})
		}
	case 13:
		{
			yylex.(*lexer).graph.Directives = append(yylex.(*lexer).graph.Directives, v1.Directive{Name: yyS[yypt-0].str})
		}
	case 14:
		{
			yyVAL.kv_list = append(yyS[yypt-1].kv_list, yyS[yypt-0].kv_list...)
		}
	case 15:
		{
			yyVAL.kv_list = []v1.KV{{yyS[yypt-2].str, yyS[yypt-0].str}}
		}
	case 16:
		{
			yyVAL.str_list = append(yyS[yypt-1].str_list, yyS[yypt-0].str_list...)
		}
	case 17:
		{
			yyVAL.str_list = append(yyS[yypt-2].str_list, yyS[yypt-0].str)
		}
	case 18:
		{
			yyVAL.str_list = []string{yyS[yypt-0].str}
		}
	case 27:
		{
			return 1
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}

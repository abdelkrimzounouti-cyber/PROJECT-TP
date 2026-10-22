// Code generated from Pcg.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // Pcg
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// we don't actually need this

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PcgParser struct {
	*antlr.BaseParser
}

var PcgParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pcgParserInit() {
	staticData := &PcgParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'program'", "'begin'", "'end'", "'write'", "'read'",
		"'else'", "'if'", "'var'", "'while'", "'then'", "'do'", "'writeln'",
		"'true'", "'false'", "'not'", "'and'", "'or'", "'<-'", "'<='", "'>='",
		"'=='", "'<>'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "':'", "'<'",
		"'>'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT", "PROGRAM", "BEGIN",
		"END", "WRITE", "READ", "ELSE", "IF", "VAR", "WHILE", "THEN", "DO",
		"WRITELN", "TRUE", "FALSE", "NOT", "AND", "OR", "ASSIGN", "LE", "GE",
		"EQ", "NE", "ADD", "SUB", "MUL", "DIV", "LPAR", "RPAR", "COLON", "LT",
		"GT", "NUMBER", "ID",
	}
	staticData.RuleNames = []string{
		"prog", "optdecls", "decls", "decl", "linst", "insts", "inst", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 163, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 31, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 40, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 54, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 107, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 131, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 158, 8, 7, 10, 7, 12, 7,
		161, 9, 7, 1, 7, 0, 1, 14, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 3, 1, 0, 28,
		29, 1, 0, 26, 27, 2, 0, 22, 25, 33, 34, 174, 0, 16, 1, 0, 0, 0, 2, 30,
		1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 53, 1, 0, 0, 0, 10,
		55, 1, 0, 0, 0, 12, 106, 1, 0, 0, 0, 14, 130, 1, 0, 0, 0, 16, 17, 5, 4,
		0, 0, 17, 18, 5, 36, 0, 0, 18, 19, 3, 2, 1, 0, 19, 20, 5, 5, 0, 0, 20,
		21, 3, 10, 5, 0, 21, 22, 5, 6, 0, 0, 22, 23, 5, 0, 0, 1, 23, 24, 6, 0,
		-1, 0, 24, 1, 1, 0, 0, 0, 25, 26, 5, 11, 0, 0, 26, 27, 3, 4, 2, 0, 27,
		28, 6, 1, -1, 0, 28, 31, 1, 0, 0, 0, 29, 31, 6, 1, -1, 0, 30, 25, 1, 0,
		0, 0, 30, 29, 1, 0, 0, 0, 31, 3, 1, 0, 0, 0, 32, 33, 3, 6, 3, 0, 33, 34,
		6, 2, -1, 0, 34, 40, 1, 0, 0, 0, 35, 36, 3, 6, 3, 0, 36, 37, 3, 4, 2, 0,
		37, 38, 6, 2, -1, 0, 38, 40, 1, 0, 0, 0, 39, 32, 1, 0, 0, 0, 39, 35, 1,
		0, 0, 0, 40, 5, 1, 0, 0, 0, 41, 42, 5, 36, 0, 0, 42, 43, 5, 32, 0, 0, 43,
		44, 5, 36, 0, 0, 44, 45, 6, 3, -1, 0, 45, 7, 1, 0, 0, 0, 46, 47, 3, 12,
		6, 0, 47, 48, 6, 4, -1, 0, 48, 54, 1, 0, 0, 0, 49, 50, 3, 12, 6, 0, 50,
		51, 3, 8, 4, 0, 51, 52, 6, 4, -1, 0, 52, 54, 1, 0, 0, 0, 53, 46, 1, 0,
		0, 0, 53, 49, 1, 0, 0, 0, 54, 9, 1, 0, 0, 0, 55, 56, 3, 8, 4, 0, 56, 57,
		6, 5, -1, 0, 57, 11, 1, 0, 0, 0, 58, 59, 5, 36, 0, 0, 59, 60, 5, 21, 0,
		0, 60, 61, 3, 14, 7, 0, 61, 62, 6, 6, -1, 0, 62, 107, 1, 0, 0, 0, 63, 64,
		5, 7, 0, 0, 64, 65, 5, 30, 0, 0, 65, 66, 3, 14, 7, 0, 66, 67, 5, 31, 0,
		0, 67, 68, 6, 6, -1, 0, 68, 107, 1, 0, 0, 0, 69, 70, 5, 15, 0, 0, 70, 71,
		5, 30, 0, 0, 71, 72, 3, 14, 7, 0, 72, 73, 5, 31, 0, 0, 73, 74, 6, 6, -1,
		0, 74, 107, 1, 0, 0, 0, 75, 76, 5, 8, 0, 0, 76, 77, 5, 30, 0, 0, 77, 78,
		5, 36, 0, 0, 78, 79, 5, 31, 0, 0, 79, 107, 6, 6, -1, 0, 80, 81, 5, 10,
		0, 0, 81, 82, 3, 14, 7, 0, 82, 83, 5, 13, 0, 0, 83, 84, 3, 10, 5, 0, 84,
		85, 5, 9, 0, 0, 85, 86, 3, 10, 5, 0, 86, 87, 5, 6, 0, 0, 87, 88, 5, 10,
		0, 0, 88, 89, 6, 6, -1, 0, 89, 107, 1, 0, 0, 0, 90, 91, 5, 10, 0, 0, 91,
		92, 3, 14, 7, 0, 92, 93, 5, 13, 0, 0, 93, 94, 3, 10, 5, 0, 94, 95, 5, 6,
		0, 0, 95, 96, 5, 10, 0, 0, 96, 97, 6, 6, -1, 0, 97, 107, 1, 0, 0, 0, 98,
		99, 5, 12, 0, 0, 99, 100, 3, 14, 7, 0, 100, 101, 5, 14, 0, 0, 101, 102,
		3, 10, 5, 0, 102, 103, 5, 6, 0, 0, 103, 104, 5, 12, 0, 0, 104, 105, 6,
		6, -1, 0, 105, 107, 1, 0, 0, 0, 106, 58, 1, 0, 0, 0, 106, 63, 1, 0, 0,
		0, 106, 69, 1, 0, 0, 0, 106, 75, 1, 0, 0, 0, 106, 80, 1, 0, 0, 0, 106,
		90, 1, 0, 0, 0, 106, 98, 1, 0, 0, 0, 107, 13, 1, 0, 0, 0, 108, 109, 6,
		7, -1, 0, 109, 110, 5, 35, 0, 0, 110, 131, 6, 7, -1, 0, 111, 112, 5, 16,
		0, 0, 112, 131, 6, 7, -1, 0, 113, 114, 5, 17, 0, 0, 114, 131, 6, 7, -1,
		0, 115, 116, 5, 36, 0, 0, 116, 131, 6, 7, -1, 0, 117, 118, 5, 30, 0, 0,
		118, 119, 3, 14, 7, 0, 119, 120, 5, 31, 0, 0, 120, 121, 6, 7, -1, 0, 121,
		131, 1, 0, 0, 0, 122, 123, 5, 27, 0, 0, 123, 124, 3, 14, 7, 7, 124, 125,
		6, 7, -1, 0, 125, 131, 1, 0, 0, 0, 126, 127, 5, 18, 0, 0, 127, 128, 3,
		14, 7, 6, 128, 129, 6, 7, -1, 0, 129, 131, 1, 0, 0, 0, 130, 108, 1, 0,
		0, 0, 130, 111, 1, 0, 0, 0, 130, 113, 1, 0, 0, 0, 130, 115, 1, 0, 0, 0,
		130, 117, 1, 0, 0, 0, 130, 122, 1, 0, 0, 0, 130, 126, 1, 0, 0, 0, 131,
		159, 1, 0, 0, 0, 132, 133, 10, 5, 0, 0, 133, 134, 7, 0, 0, 0, 134, 135,
		3, 14, 7, 6, 135, 136, 6, 7, -1, 0, 136, 158, 1, 0, 0, 0, 137, 138, 10,
		4, 0, 0, 138, 139, 7, 1, 0, 0, 139, 140, 3, 14, 7, 5, 140, 141, 6, 7, -1,
		0, 141, 158, 1, 0, 0, 0, 142, 143, 10, 3, 0, 0, 143, 144, 7, 2, 0, 0, 144,
		145, 3, 14, 7, 4, 145, 146, 6, 7, -1, 0, 146, 158, 1, 0, 0, 0, 147, 148,
		10, 2, 0, 0, 148, 149, 5, 19, 0, 0, 149, 150, 3, 14, 7, 3, 150, 151, 6,
		7, -1, 0, 151, 158, 1, 0, 0, 0, 152, 153, 10, 1, 0, 0, 153, 154, 5, 20,
		0, 0, 154, 155, 3, 14, 7, 2, 155, 156, 6, 7, -1, 0, 156, 158, 1, 0, 0,
		0, 157, 132, 1, 0, 0, 0, 157, 137, 1, 0, 0, 0, 157, 142, 1, 0, 0, 0, 157,
		147, 1, 0, 0, 0, 157, 152, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157,
		1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 15, 1, 0, 0, 0, 161, 159, 1, 0,
		0, 0, 7, 30, 39, 53, 106, 130, 157, 159,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PcgParserInit initializes any static state used to implement PcgParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPcgParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PcgParserInit() {
	staticData := &PcgParserStaticData
	staticData.once.Do(pcgParserInit)
}

// NewPcgParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPcgParser(input antlr.TokenStream) *PcgParser {
	PcgParserInit()
	this := new(PcgParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PcgParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Pcg.g4"

	return this
}

// PcgParser tokens.
const (
	PcgParserEOF           = antlr.TokenEOF
	PcgParserWHITESPACE    = 1
	PcgParserBLOCK_COMMENT = 2
	PcgParserLINE_COMMENT  = 3
	PcgParserPROGRAM       = 4
	PcgParserBEGIN         = 5
	PcgParserEND           = 6
	PcgParserWRITE         = 7
	PcgParserREAD          = 8
	PcgParserELSE          = 9
	PcgParserIF            = 10
	PcgParserVAR           = 11
	PcgParserWHILE         = 12
	PcgParserTHEN          = 13
	PcgParserDO            = 14
	PcgParserWRITELN       = 15
	PcgParserTRUE          = 16
	PcgParserFALSE         = 17
	PcgParserNOT           = 18
	PcgParserAND           = 19
	PcgParserOR            = 20
	PcgParserASSIGN        = 21
	PcgParserLE            = 22
	PcgParserGE            = 23
	PcgParserEQ            = 24
	PcgParserNE            = 25
	PcgParserADD           = 26
	PcgParserSUB           = 27
	PcgParserMUL           = 28
	PcgParserDIV           = 29
	PcgParserLPAR          = 30
	PcgParserRPAR          = 31
	PcgParserCOLON         = 32
	PcgParserLT            = 33
	PcgParserGT            = 34
	PcgParserNUMBER        = 35
	PcgParserID            = 36
)

// PcgParser rules.
const (
	PcgParserRULE_prog     = 0
	PcgParserRULE_optdecls = 1
	PcgParserRULE_decls    = 2
	PcgParserRULE_decl     = 3
	PcgParserRULE_linst    = 4
	PcgParserRULE_insts    = 5
	PcgParserRULE_inst     = 6
	PcgParserRULE_expr     = 7
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetD returns the d rule contexts.
	GetD() IOptdeclsContext

	// GetI returns the i rule contexts.
	GetI() IInstsContext

	// SetD sets the d rule contexts.
	SetD(IOptdeclsContext)

	// SetI sets the i rule contexts.
	SetI(IInstsContext)

	// Getter signatures
	PROGRAM() antlr.TerminalNode
	ID() antlr.TerminalNode
	BEGIN() antlr.TerminalNode
	END() antlr.TerminalNode
	EOF() antlr.TerminalNode
	Optdecls() IOptdeclsContext
	Insts() IInstsContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	d      IOptdeclsContext
	i      IInstsContext
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PcgParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) GetD() IOptdeclsContext { return s.d }

func (s *ProgContext) GetI() IInstsContext { return s.i }

func (s *ProgContext) SetD(v IOptdeclsContext) { s.d = v }

func (s *ProgContext) SetI(v IInstsContext) { s.i = v }

func (s *ProgContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(PcgParserPROGRAM, 0)
}

func (s *ProgContext) ID() antlr.TerminalNode {
	return s.GetToken(PcgParserID, 0)
}

func (s *ProgContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(PcgParserBEGIN, 0)
}

func (s *ProgContext) END() antlr.TerminalNode {
	return s.GetToken(PcgParserEND, 0)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(PcgParserEOF, 0)
}

func (s *ProgContext) Optdecls() IOptdeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptdeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptdeclsContext)
}

func (s *ProgContext) Insts() IInstsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstsContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PcgParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PcgParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(PcgParserPROGRAM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(17)
		p.Match(PcgParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(18)

		var _x = p.Optdecls()

		localctx.(*ProgContext).d = _x
	}
	{
		p.SetState(19)
		p.Match(PcgParserBEGIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)

		var _x = p.Insts()

		localctx.(*ProgContext).i = _x
	}
	{
		p.SetState(21)
		p.Match(PcgParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.Match(PcgParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	root = Node2(Pgm, localctx.(*ProgContext).GetD().GetVal(), localctx.(*ProgContext).GetI().GetVal(), 1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptdeclsContext is an interface to support dynamic dispatch.
type IOptdeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// GetD returns the d rule contexts.
	GetD() IDeclsContext

	// SetD sets the d rule contexts.
	SetD(IDeclsContext)

	// GetVal returns the val attribute.
	GetVal() *Tree

	// SetVal sets the val attribute.
	SetVal(*Tree)

	// Getter signatures
	VAR() antlr.TerminalNode
	Decls() IDeclsContext

	// IsOptdeclsContext differentiates from other interfaces.
	IsOptdeclsContext()
}

type OptdeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    *Tree
	_VAR   antlr.Token
	d      IDeclsContext
}

func NewEmptyOptdeclsContext() *OptdeclsContext {
	var p = new(OptdeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_optdecls
	return p
}

func InitEmptyOptdeclsContext(p *OptdeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_optdecls
}

func (*OptdeclsContext) IsOptdeclsContext() {}

func NewOptdeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptdeclsContext {
	var p = new(OptdeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PcgParserRULE_optdecls

	return p
}

func (s *OptdeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *OptdeclsContext) Get_VAR() antlr.Token { return s._VAR }

func (s *OptdeclsContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *OptdeclsContext) GetD() IDeclsContext { return s.d }

func (s *OptdeclsContext) SetD(v IDeclsContext) { s.d = v }

func (s *OptdeclsContext) GetVal() *Tree { return s.val }

func (s *OptdeclsContext) SetVal(v *Tree) { s.val = v }

func (s *OptdeclsContext) VAR() antlr.TerminalNode {
	return s.GetToken(PcgParserVAR, 0)
}

func (s *OptdeclsContext) Decls() IDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclsContext)
}

func (s *OptdeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptdeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PcgParser) Optdecls() (localctx IOptdeclsContext) {
	localctx = NewOptdeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PcgParserRULE_optdecls)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PcgParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)

			var _m = p.Match(PcgParserVAR)

			localctx.(*OptdeclsContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)

			var _x = p.Decls()

			localctx.(*OptdeclsContext).d = _x
		}
		localctx.(*OptdeclsContext).val = Node1(Dblock, localctx.(*OptdeclsContext).GetD().GetVal(), (func() int {
			if localctx.(*OptdeclsContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*OptdeclsContext).Get_VAR().GetLine()
			}
		}()))

	case PcgParserBEGIN:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*OptdeclsContext).val = nil

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclsContext is an interface to support dynamic dispatch.
type IDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetD returns the d rule contexts.
	GetD() IDeclContext

	// GetDs returns the ds rule contexts.
	GetDs() IDeclsContext

	// SetD sets the d rule contexts.
	SetD(IDeclContext)

	// SetDs sets the ds rule contexts.
	SetDs(IDeclsContext)

	// GetVal returns the val attribute.
	GetVal() *Tree

	// SetVal sets the val attribute.
	SetVal(*Tree)

	// Getter signatures
	Decl() IDeclContext
	Decls() IDeclsContext

	// IsDeclsContext differentiates from other interfaces.
	IsDeclsContext()
}

type DeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    *Tree
	d      IDeclContext
	ds     IDeclsContext
}

func NewEmptyDeclsContext() *DeclsContext {
	var p = new(DeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_decls
	return p
}

func InitEmptyDeclsContext(p *DeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_decls
}

func (*DeclsContext) IsDeclsContext() {}

func NewDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclsContext {
	var p = new(DeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PcgParserRULE_decls

	return p
}

func (s *DeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclsContext) GetD() IDeclContext { return s.d }

func (s *DeclsContext) GetDs() IDeclsContext { return s.ds }

func (s *DeclsContext) SetD(v IDeclContext) { s.d = v }

func (s *DeclsContext) SetDs(v IDeclsContext) { s.ds = v }

func (s *DeclsContext) GetVal() *Tree { return s.val }

func (s *DeclsContext) SetVal(v *Tree) { s.val = v }

func (s *DeclsContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *DeclsContext) Decls() IDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclsContext)
}

func (s *DeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PcgParser) Decls() (localctx IDeclsContext) {
	localctx = NewDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PcgParserRULE_decls)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)

			var _x = p.Decl()

			localctx.(*DeclsContext).d = _x
		}
		localctx.(*DeclsContext).val = NodeN(Dblock, []*Tree{localctx.(*DeclsContext).GetD().GetVal()}, localctx.(*DeclsContext).GetD().GetVal().line)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)

			var _x = p.Decl()

			localctx.(*DeclsContext).d = _x
		}
		{
			p.SetState(36)

			var _x = p.Decls()

			localctx.(*DeclsContext).ds = _x
		}
		localctx.(*DeclsContext).val = NodeN(Dblock, append([]*Tree{localctx.(*DeclsContext).GetD().GetVal()}, localctx.(*DeclsContext).GetDs().GetVal().ss...), localctx.(*DeclsContext).GetD().GetVal().line)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// GetTyp returns the typ token.
	GetTyp() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetTyp sets the typ token.
	SetTyp(antlr.Token)

	// GetVal returns the val attribute.
	GetVal() *Tree

	// SetVal sets the val attribute.
	SetVal(*Tree)

	// Getter signatures
	COLON() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    *Tree
	id     antlr.Token
	typ    antlr.Token
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_decl
	return p
}

func InitEmptyDeclContext(p *DeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_decl
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PcgParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) GetId() antlr.Token { return s.id }

func (s *DeclContext) GetTyp() antlr.Token { return s.typ }

func (s *DeclContext) SetId(v antlr.Token) { s.id = v }

func (s *DeclContext) SetTyp(v antlr.Token) { s.typ = v }

func (s *DeclContext) GetVal() *Tree { return s.val }

func (s *DeclContext) SetVal(v *Tree) { s.val = v }

func (s *DeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(PcgParserCOLON, 0)
}

func (s *DeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PcgParserID)
}

func (s *DeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PcgParserID, i)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PcgParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PcgParserRULE_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)

		var _m = p.Match(PcgParserID)

		localctx.(*DeclContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(PcgParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)

		var _m = p.Match(PcgParserID)

		localctx.(*DeclContext).typ = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*DeclContext).val = Node2(Dvar,
		Leaf(Did, localctx.(*DeclContext).GetId().GetText(), localctx.(*DeclContext).GetId().GetLine()),
		Leaf(Did, localctx.(*DeclContext).GetTyp().GetText(), localctx.(*DeclContext).GetTyp().GetLine()),
		localctx.(*DeclContext).GetId().GetLine())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILinstContext is an interface to support dynamic dispatch.
type ILinstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetI returns the i rule contexts.
	GetI() IInstContext

	// GetRest returns the rest rule contexts.
	GetRest() ILinstContext

	// SetI sets the i rule contexts.
	SetI(IInstContext)

	// SetRest sets the rest rule contexts.
	SetRest(ILinstContext)

	// GetVal returns the val attribute.
	GetVal() []*Tree

	// SetVal sets the val attribute.
	SetVal([]*Tree)

	// Getter signatures
	Inst() IInstContext
	Linst() ILinstContext

	// IsLinstContext differentiates from other interfaces.
	IsLinstContext()
}

type LinstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    []*Tree
	i      IInstContext
	rest   ILinstContext
}

func NewEmptyLinstContext() *LinstContext {
	var p = new(LinstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_linst
	return p
}

func InitEmptyLinstContext(p *LinstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_linst
}

func (*LinstContext) IsLinstContext() {}

func NewLinstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinstContext {
	var p = new(LinstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PcgParserRULE_linst

	return p
}

func (s *LinstContext) GetParser() antlr.Parser { return s.parser }

func (s *LinstContext) GetI() IInstContext { return s.i }

func (s *LinstContext) GetRest() ILinstContext { return s.rest }

func (s *LinstContext) SetI(v IInstContext) { s.i = v }

func (s *LinstContext) SetRest(v ILinstContext) { s.rest = v }

func (s *LinstContext) GetVal() []*Tree { return s.val }

func (s *LinstContext) SetVal(v []*Tree) { s.val = v }

func (s *LinstContext) Inst() IInstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstContext)
}

func (s *LinstContext) Linst() ILinstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILinstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILinstContext)
}

func (s *LinstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PcgParser) Linst() (localctx ILinstContext) {
	localctx = NewLinstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PcgParserRULE_linst)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)

			var _x = p.Inst()

			localctx.(*LinstContext).i = _x
		}
		localctx.(*LinstContext).val = []*Tree{localctx.(*LinstContext).GetI().GetVal()}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)

			var _x = p.Inst()

			localctx.(*LinstContext).i = _x
		}
		{
			p.SetState(50)

			var _x = p.Linst()

			localctx.(*LinstContext).rest = _x
		}
		localctx.(*LinstContext).val = append([]*Tree{localctx.(*LinstContext).GetI().GetVal()}, localctx.(*LinstContext).GetRest().GetVal()...)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstsContext is an interface to support dynamic dispatch.
type IInstsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ILinstContext

	// SetL sets the l rule contexts.
	SetL(ILinstContext)

	// GetVal returns the val attribute.
	GetVal() *Tree

	// SetVal sets the val attribute.
	SetVal(*Tree)

	// Getter signatures
	Linst() ILinstContext

	// IsInstsContext differentiates from other interfaces.
	IsInstsContext()
}

type InstsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    *Tree
	l      ILinstContext
}

func NewEmptyInstsContext() *InstsContext {
	var p = new(InstsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_insts
	return p
}

func InitEmptyInstsContext(p *InstsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_insts
}

func (*InstsContext) IsInstsContext() {}

func NewInstsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstsContext {
	var p = new(InstsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PcgParserRULE_insts

	return p
}

func (s *InstsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstsContext) GetL() ILinstContext { return s.l }

func (s *InstsContext) SetL(v ILinstContext) { s.l = v }

func (s *InstsContext) GetVal() *Tree { return s.val }

func (s *InstsContext) SetVal(v *Tree) { s.val = v }

func (s *InstsContext) Linst() ILinstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILinstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILinstContext)
}

func (s *InstsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PcgParser) Insts() (localctx IInstsContext) {
	localctx = NewInstsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PcgParserRULE_insts)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)

		var _x = p.Linst()

		localctx.(*InstsContext).l = _x
	}
	localctx.(*InstsContext).val = NodeN(Iblock, localctx.(*InstsContext).GetL().GetVal(), localctx.(*InstsContext).GetL().GetVal()[0].line)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstContext is an interface to support dynamic dispatch.
type IInstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// GetK returns the k token.
	GetK() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetK sets the k token.
	SetK(antlr.Token)

	// GetE returns the e rule contexts.
	GetE() IExprContext

	// GetC returns the c rule contexts.
	GetC() IExprContext

	// GetT returns the t rule contexts.
	GetT() IInstsContext

	// GetF returns the f rule contexts.
	GetF() IInstsContext

	// GetB returns the b rule contexts.
	GetB() IInstsContext

	// SetE sets the e rule contexts.
	SetE(IExprContext)

	// SetC sets the c rule contexts.
	SetC(IExprContext)

	// SetT sets the t rule contexts.
	SetT(IInstsContext)

	// SetF sets the f rule contexts.
	SetF(IInstsContext)

	// SetB sets the b rule contexts.
	SetB(IInstsContext)

	// GetVal returns the val attribute.
	GetVal() *Tree

	// SetVal sets the val attribute.
	SetVal(*Tree)

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode
	Expr() IExprContext
	LPAR() antlr.TerminalNode
	RPAR() antlr.TerminalNode
	WRITE() antlr.TerminalNode
	WRITELN() antlr.TerminalNode
	READ() antlr.TerminalNode
	THEN() antlr.TerminalNode
	ELSE() antlr.TerminalNode
	END() antlr.TerminalNode
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllInsts() []IInstsContext
	Insts(i int) IInstsContext
	DO() antlr.TerminalNode
	AllWHILE() []antlr.TerminalNode
	WHILE(i int) antlr.TerminalNode

	// IsInstContext differentiates from other interfaces.
	IsInstContext()
}

type InstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    *Tree
	id     antlr.Token
	e      IExprContext
	k      antlr.Token
	c      IExprContext
	t      IInstsContext
	f      IInstsContext
	b      IInstsContext
}

func NewEmptyInstContext() *InstContext {
	var p = new(InstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_inst
	return p
}

func InitEmptyInstContext(p *InstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_inst
}

func (*InstContext) IsInstContext() {}

func NewInstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstContext {
	var p = new(InstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PcgParserRULE_inst

	return p
}

func (s *InstContext) GetParser() antlr.Parser { return s.parser }

func (s *InstContext) GetId() antlr.Token { return s.id }

func (s *InstContext) GetK() antlr.Token { return s.k }

func (s *InstContext) SetId(v antlr.Token) { s.id = v }

func (s *InstContext) SetK(v antlr.Token) { s.k = v }

func (s *InstContext) GetE() IExprContext { return s.e }

func (s *InstContext) GetC() IExprContext { return s.c }

func (s *InstContext) GetT() IInstsContext { return s.t }

func (s *InstContext) GetF() IInstsContext { return s.f }

func (s *InstContext) GetB() IInstsContext { return s.b }

func (s *InstContext) SetE(v IExprContext) { s.e = v }

func (s *InstContext) SetC(v IExprContext) { s.c = v }

func (s *InstContext) SetT(v IInstsContext) { s.t = v }

func (s *InstContext) SetF(v IInstsContext) { s.f = v }

func (s *InstContext) SetB(v IInstsContext) { s.b = v }

func (s *InstContext) GetVal() *Tree { return s.val }

func (s *InstContext) SetVal(v *Tree) { s.val = v }

func (s *InstContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PcgParserASSIGN, 0)
}

func (s *InstContext) ID() antlr.TerminalNode {
	return s.GetToken(PcgParserID, 0)
}

func (s *InstContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InstContext) LPAR() antlr.TerminalNode {
	return s.GetToken(PcgParserLPAR, 0)
}

func (s *InstContext) RPAR() antlr.TerminalNode {
	return s.GetToken(PcgParserRPAR, 0)
}

func (s *InstContext) WRITE() antlr.TerminalNode {
	return s.GetToken(PcgParserWRITE, 0)
}

func (s *InstContext) WRITELN() antlr.TerminalNode {
	return s.GetToken(PcgParserWRITELN, 0)
}

func (s *InstContext) READ() antlr.TerminalNode {
	return s.GetToken(PcgParserREAD, 0)
}

func (s *InstContext) THEN() antlr.TerminalNode {
	return s.GetToken(PcgParserTHEN, 0)
}

func (s *InstContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PcgParserELSE, 0)
}

func (s *InstContext) END() antlr.TerminalNode {
	return s.GetToken(PcgParserEND, 0)
}

func (s *InstContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(PcgParserIF)
}

func (s *InstContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(PcgParserIF, i)
}

func (s *InstContext) AllInsts() []IInstsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstsContext); ok {
			len++
		}
	}

	tst := make([]IInstsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstsContext); ok {
			tst[i] = t.(IInstsContext)
			i++
		}
	}

	return tst
}

func (s *InstContext) Insts(i int) IInstsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstsContext)
}

func (s *InstContext) DO() antlr.TerminalNode {
	return s.GetToken(PcgParserDO, 0)
}

func (s *InstContext) AllWHILE() []antlr.TerminalNode {
	return s.GetTokens(PcgParserWHILE)
}

func (s *InstContext) WHILE(i int) antlr.TerminalNode {
	return s.GetToken(PcgParserWHILE, i)
}

func (s *InstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PcgParser) Inst() (localctx IInstContext) {
	localctx = NewInstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PcgParserRULE_inst)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)

			var _m = p.Match(PcgParserID)

			localctx.(*InstContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Match(PcgParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)

			var _x = p.expr(0)

			localctx.(*InstContext).e = _x
		}

		localctx.(*InstContext).val = Node2(Iassign,
			Leaf(Did, localctx.(*InstContext).GetId().GetText(), localctx.(*InstContext).GetId().GetLine()),
			localctx.(*InstContext).GetE().GetVal(), localctx.(*InstContext).GetId().GetLine())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)

			var _m = p.Match(PcgParserWRITE)

			localctx.(*InstContext).k = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(PcgParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)

			var _x = p.expr(0)

			localctx.(*InstContext).e = _x
		}
		{
			p.SetState(66)
			p.Match(PcgParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InstContext).val = Node1(Iwrite, localctx.(*InstContext).GetE().GetVal(), localctx.(*InstContext).GetK().GetLine())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)

			var _m = p.Match(PcgParserWRITELN)

			localctx.(*InstContext).k = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(PcgParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)

			var _x = p.expr(0)

			localctx.(*InstContext).e = _x
		}
		{
			p.SetState(72)
			p.Match(PcgParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InstContext).val = Node1(Iwriteln, localctx.(*InstContext).GetE().GetVal(), localctx.(*InstContext).GetK().GetLine())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)

			var _m = p.Match(PcgParserREAD)

			localctx.(*InstContext).k = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(PcgParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)

			var _m = p.Match(PcgParserID)

			localctx.(*InstContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(PcgParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InstContext).val = Node1(Iread, Leaf(Did, localctx.(*InstContext).GetId().GetText(), localctx.(*InstContext).GetId().GetLine()), localctx.(*InstContext).GetK().GetLine())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(80)

			var _m = p.Match(PcgParserIF)

			localctx.(*InstContext).k = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)

			var _x = p.expr(0)

			localctx.(*InstContext).c = _x
		}
		{
			p.SetState(82)
			p.Match(PcgParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)

			var _x = p.Insts()

			localctx.(*InstContext).t = _x
		}
		{
			p.SetState(84)
			p.Match(PcgParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)

			var _x = p.Insts()

			localctx.(*InstContext).f = _x
		}
		{
			p.SetState(86)
			p.Match(PcgParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(PcgParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InstContext).val = Node3(Iifelse, localctx.(*InstContext).GetC().GetVal(), localctx.(*InstContext).GetT().GetVal(), localctx.(*InstContext).GetF().GetVal(), localctx.(*InstContext).GetK().GetLine())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(90)

			var _m = p.Match(PcgParserIF)

			localctx.(*InstContext).k = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)

			var _x = p.expr(0)

			localctx.(*InstContext).c = _x
		}
		{
			p.SetState(92)
			p.Match(PcgParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)

			var _x = p.Insts()

			localctx.(*InstContext).t = _x
		}
		{
			p.SetState(94)
			p.Match(PcgParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(PcgParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InstContext).val = Node2(Iif, localctx.(*InstContext).GetC().GetVal(), localctx.(*InstContext).GetT().GetVal(), localctx.(*InstContext).GetK().GetLine())

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(98)

			var _m = p.Match(PcgParserWHILE)

			localctx.(*InstContext).k = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)

			var _x = p.expr(0)

			localctx.(*InstContext).c = _x
		}
		{
			p.SetState(100)
			p.Match(PcgParserDO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)

			var _x = p.Insts()

			localctx.(*InstContext).b = _x
		}
		{
			p.SetState(102)
			p.Match(PcgParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(PcgParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InstContext).val = Node2(Iwhile, localctx.(*InstContext).GetC().GetVal(), localctx.(*InstContext).GetB().GetVal(), localctx.(*InstContext).GetK().GetLine())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// Get_TRUE returns the _TRUE token.
	Get_TRUE() antlr.Token

	// Get_FALSE returns the _FALSE token.
	Get_FALSE() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// Set_TRUE sets the _TRUE token.
	Set_TRUE(antlr.Token)

	// Set_FALSE sets the _FALSE token.
	Set_FALSE(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetE returns the e rule contexts.
	GetE() IExprContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetE sets the e rule contexts.
	SetE(IExprContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// GetVal returns the val attribute.
	GetVal() *Tree

	// SetVal sets the val attribute.
	SetVal(*Tree)

	// Getter signatures
	NUMBER() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	RPAR() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	SUB() antlr.TerminalNode
	NOT() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	ADD() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    *Tree
	e1     IExprContext
	n      antlr.Token
	_TRUE  antlr.Token
	_FALSE antlr.Token
	id     antlr.Token
	e      IExprContext
	op     antlr.Token
	e2     IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PcgParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PcgParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetN() antlr.Token { return s.n }

func (s *ExprContext) Get_TRUE() antlr.Token { return s._TRUE }

func (s *ExprContext) Get_FALSE() antlr.Token { return s._FALSE }

func (s *ExprContext) GetId() antlr.Token { return s.id }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetN(v antlr.Token) { s.n = v }

func (s *ExprContext) Set_TRUE(v antlr.Token) { s._TRUE = v }

func (s *ExprContext) Set_FALSE(v antlr.Token) { s._FALSE = v }

func (s *ExprContext) SetId(v antlr.Token) { s.id = v }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) GetE1() IExprContext { return s.e1 }

func (s *ExprContext) GetE() IExprContext { return s.e }

func (s *ExprContext) GetE2() IExprContext { return s.e2 }

func (s *ExprContext) SetE1(v IExprContext) { s.e1 = v }

func (s *ExprContext) SetE(v IExprContext) { s.e = v }

func (s *ExprContext) SetE2(v IExprContext) { s.e2 = v }

func (s *ExprContext) GetVal() *Tree { return s.val }

func (s *ExprContext) SetVal(v *Tree) { s.val = v }

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PcgParserNUMBER, 0)
}

func (s *ExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(PcgParserTRUE, 0)
}

func (s *ExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(PcgParserFALSE, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(PcgParserID, 0)
}

func (s *ExprContext) LPAR() antlr.TerminalNode {
	return s.GetToken(PcgParserLPAR, 0)
}

func (s *ExprContext) RPAR() antlr.TerminalNode {
	return s.GetToken(PcgParserRPAR, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(PcgParserSUB, 0)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(PcgParserNOT, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(PcgParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(PcgParserDIV, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(PcgParserADD, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(PcgParserLT, 0)
}

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(PcgParserLE, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(PcgParserGT, 0)
}

func (s *ExprContext) GE() antlr.TerminalNode {
	return s.GetToken(PcgParserGE, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(PcgParserEQ, 0)
}

func (s *ExprContext) NE() antlr.TerminalNode {
	return s.GetToken(PcgParserNE, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(PcgParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(PcgParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PcgParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PcgParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, PcgParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PcgParserNUMBER:
		{
			p.SetState(109)

			var _m = p.Match(PcgParserNUMBER)

			localctx.(*ExprContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).val = Leaf(Eint, localctx.(*ExprContext).GetN().GetText(), localctx.(*ExprContext).GetN().GetLine())

	case PcgParserTRUE:
		{
			p.SetState(111)

			var _m = p.Match(PcgParserTRUE)

			localctx.(*ExprContext)._TRUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).val = Leaf(Ebool, "true", (func() int {
			if localctx.(*ExprContext).Get_TRUE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRUE().GetLine()
			}
		}()))

	case PcgParserFALSE:
		{
			p.SetState(113)

			var _m = p.Match(PcgParserFALSE)

			localctx.(*ExprContext)._FALSE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).val = Leaf(Ebool, "false", (func() int {
			if localctx.(*ExprContext).Get_FALSE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FALSE().GetLine()
			}
		}()))

	case PcgParserID:
		{
			p.SetState(115)

			var _m = p.Match(PcgParserID)

			localctx.(*ExprContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).val = Leaf(Did, localctx.(*ExprContext).GetId().GetText(), localctx.(*ExprContext).GetId().GetLine())

	case PcgParserLPAR:
		{
			p.SetState(117)
			p.Match(PcgParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)

			var _x = p.expr(0)

			localctx.(*ExprContext).e = _x
		}
		{
			p.SetState(119)
			p.Match(PcgParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).val = localctx.(*ExprContext).GetE().GetVal()

	case PcgParserSUB:
		{
			p.SetState(122)

			var _m = p.Match(PcgParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)

			var _x = p.expr(7)

			localctx.(*ExprContext).e = _x
		}
		localctx.(*ExprContext).val = Node1(Eneg, localctx.(*ExprContext).GetE().GetVal(), localctx.(*ExprContext).GetOp().GetLine())

	case PcgParserNOT:
		{
			p.SetState(126)

			var _m = p.Match(PcgParserNOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)

			var _x = p.expr(6)

			localctx.(*ExprContext).e = _x
		}
		localctx.(*ExprContext).val = Node1(Enot, localctx.(*ExprContext).GetE().GetVal(), localctx.(*ExprContext).GetOp().GetLine())

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PcgParserRULE_expr)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(133)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PcgParserMUL || _la == PcgParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(134)

					var _x = p.expr(6)

					localctx.(*ExprContext).e2 = _x
				}

				var k Kind
				if localctx.(*ExprContext).GetOp().GetTokenType() == PcgParserMUL {
					k = Emul
				} else {
					k = Ediv
				}
				localctx.(*ExprContext).val = Node2(k, localctx.(*ExprContext).GetE1().GetVal(), localctx.(*ExprContext).GetE2().GetVal(), localctx.(*ExprContext).GetOp().GetLine())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PcgParserRULE_expr)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(138)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PcgParserADD || _la == PcgParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(139)

					var _x = p.expr(5)

					localctx.(*ExprContext).e2 = _x
				}

				var k Kind
				if localctx.(*ExprContext).GetOp().GetTokenType() == PcgParserADD {
					k = Eadd
				} else {
					k = Esub
				}
				localctx.(*ExprContext).val = Node2(k, localctx.(*ExprContext).GetE1().GetVal(), localctx.(*ExprContext).GetE2().GetVal(), localctx.(*ExprContext).GetOp().GetLine())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PcgParserRULE_expr)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(143)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25832718336) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(144)

					var _x = p.expr(4)

					localctx.(*ExprContext).e2 = _x
				}

				var k Kind
				switch localctx.(*ExprContext).GetOp().GetTokenType() {
				case PcgParserLT:
					k = Elt
				case PcgParserLE:
					k = Ele
				case PcgParserGT:
					k = Egt
				case PcgParserGE:
					k = Ege
				case PcgParserEQ:
					k = Eeq
				case PcgParserNE:
					k = Ene
				}
				localctx.(*ExprContext).val = Node2(k, localctx.(*ExprContext).GetE1().GetVal(), localctx.(*ExprContext).GetE2().GetVal(), localctx.(*ExprContext).GetOp().GetLine())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PcgParserRULE_expr)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(148)

					var _m = p.Match(PcgParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(149)

					var _x = p.expr(3)

					localctx.(*ExprContext).e2 = _x
				}
				localctx.(*ExprContext).val = Node2(Eand, localctx.(*ExprContext).GetE1().GetVal(), localctx.(*ExprContext).GetE2().GetVal(), localctx.(*ExprContext).GetOp().GetLine())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PcgParserRULE_expr)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(153)

					var _m = p.Match(PcgParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(154)

					var _x = p.expr(2)

					localctx.(*ExprContext).e2 = _x
				}
				localctx.(*ExprContext).val = Node2(Eor, localctx.(*ExprContext).GetE1().GetVal(), localctx.(*ExprContext).GetE2().GetVal(), localctx.(*ExprContext).GetOp().GetLine())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *PcgParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PcgParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

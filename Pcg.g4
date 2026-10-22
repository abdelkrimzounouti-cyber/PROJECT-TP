// pcg.g4
grammar Pcg;

@header {
    // we don't actually need this
}

// Pseudo tokens


// Lexical units (tokens)



// 1. WHITESPACE & COMMENTS (Must be handled first)
WHITESPACE     : [ \r\n\t]+ -> skip;
BLOCK_COMMENT  : '(*' .*? '*)' -> skip;
LINE_COMMENT   : '--' ~[\r\n]* -> skip;

PROGRAM : 'program';
BEGIN : 'begin';
END  : 'end';
WRITE  : 'write';
READ  : 'read';
ELSE    : 'else';
IF    : 'if';
VAR    : 'var';
WHILE  : 'while';
THEN : 'then';
DO   : 'do';

WRITELN : 'writeln';
TRUE    : 'true';
FALSE   : 'false';
NOT     : 'not';
AND     : 'and';
OR      : 'or';

ASSIGN   : '<-';
LE    : '<=';
GE   : '>=';
EQ  : '==';
NE : '<>';

ADD : '+';
SUB  : '-';
MUL  : '*';
DIV  : '/';
LPAR : '(';
RPAR : ')';
COLON : ':'; 
LT       : '<';
GT  : '>';

fragment DIGIT : [0-9];
NUMBER   : ('-')? DIGIT+;

ID : [a-zA-Z_] [a-zA-Z_0-9]*; 



// Grammar productions

prog
    : PROGRAM ID d=optdecls BEGIN i=insts END EOF
      { root = Node2(Pgm, $d.val, $i.val, 1) }
    ;

optdecls returns [*Tree val]
    : VAR d=decls
      { $val = Node1(Dblock, $d.val, $VAR.line) }
    | 
      { $val = nil }
    ;

decls returns [*Tree val]
    : d=decl
      { $val = NodeN(Dblock, []*Tree{$d.val}, $d.val.line) }
    | d=decl ds=decls
      { $val = NodeN(Dblock, append([]*Tree{$d.val}, $ds.val.ss...), $d.val.line) }
    ;

decl returns [*Tree val]
    : id=ID COLON typ=ID
      {
          $val = Node2(Dvar,
              Leaf(Did, $id.GetText(), $id.GetLine()),
              Leaf(Did, $typ.GetText(), $typ.GetLine()),
              $id.GetLine())
      }
    ;

linst returns [[]*Tree val]
    : i=inst
      { $val = []*Tree{$i.val} }
    | i=inst rest=linst
      { $val = append([]*Tree{$i.val}, $rest.val...) }
    ;

insts returns [*Tree val]
    : l=linst
      { $val = NodeN(Iblock, $l.val, $l.val[0].line) }
    ;

inst returns [*Tree val]
    : id=ID ASSIGN e=expr
      {
          $val = Node2(Iassign,
              Leaf(Did, $id.GetText(), $id.GetLine()),
              $e.val, $id.GetLine())
      }
    | k=WRITE LPAR e=expr RPAR
      { $val = Node1(Iwrite, $e.val, $k.GetLine()) }
    | k=WRITELN LPAR e=expr RPAR
      { $val = Node1(Iwriteln, $e.val, $k.GetLine()) }
    | k=READ LPAR id=ID RPAR
      { $val = Node1(Iread, Leaf(Did, $id.GetText(), $id.GetLine()), $k.GetLine()) }
    | k=IF c=expr THEN t=insts ELSE f=insts END IF
      { $val = Node3(Iifelse, $c.val, $t.val, $f.val, $k.GetLine()) }
    | k=IF c=expr THEN t=insts END IF
      { $val = Node2(Iif, $c.val, $t.val, $k.GetLine()) }
    | k=WHILE c=expr DO b=insts END WHILE
      { $val = Node2(Iwhile, $c.val, $b.val, $k.GetLine()) }
    ;

expr returns [*Tree val]
    : n=NUMBER
      { $val = Leaf(Eint, $n.GetText(), $n.GetLine()) }
    | TRUE
      { $val = Leaf(Ebool, "true", $TRUE.line) }
    | FALSE
      { $val = Leaf(Ebool, "false", $FALSE.line) }
    | id=ID
      { $val = Leaf(Did, $id.GetText(), $id.GetLine()) }
    | LPAR e=expr RPAR
      { $val = $e.val }
    | op=SUB e=expr
      { $val = Node1(Eneg, $e.val, $op.GetLine()) }
    | op=NOT e=expr
      { $val = Node1(Enot, $e.val, $op.GetLine()) }
    | e1=expr op=(MUL | DIV) e2=expr
      {
          var k Kind
          if $op.GetTokenType() == PcgParserMUL {
              k = Emul
          } else {
              k = Ediv
          }
          $val = Node2(k, $e1.val, $e2.val, $op.GetLine())
      }
    | e1=expr op=(ADD | SUB) e2=expr
      {
          var k Kind
          if $op.GetTokenType() == PcgParserADD {
              k = Eadd
          } else {
              k = Esub
          }
          $val = Node2(k, $e1.val, $e2.val, $op.GetLine())
      }
    | e1=expr op=(LT | LE | GT | GE | EQ | NE) e2=expr
      {
          var k Kind
          switch $op.GetTokenType() {
          case PcgParserLT: k = Elt
          case PcgParserLE: k = Ele
          case PcgParserGT: k = Egt
          case PcgParserGE: k = Ege
          case PcgParserEQ: k = Eeq
          case PcgParserNE: k = Ene
          }
          $val = Node2(k, $e1.val, $e2.val, $op.GetLine())
      }
    | e1=expr op=AND e2=expr
      { $val = Node2(Eand, $e1.val, $e2.val, $op.GetLine()) }
    | e1=expr op=OR e2=expr
      { $val = Node2(Eor, $e1.val, $e2.val, $op.GetLine()) }
    ;

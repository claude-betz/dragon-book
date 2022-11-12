// 3.5.1 a

{%
	/* definitions of manifest constants
	LT, LE, EQ, NE, GT, GE,
	IF, THEN, ELSE, ID, NUMBER, RELOP
	*/
%}

/* regular definition */
delim	[ \t\n]
ws	[delim]+
letter 	[A-Za-z]
digit	[0-9]
id	{letter}({letter}{digit})*
number	{digit}+(\.{digit}+)?(E[+-]?{digit}+)?

%%

{ws}	{/* no action and no return */}
if	{return(IF);}
then	{return(THEN);}
else	{return(ELSE);}
{id}	{yylval = (int) installID(); return(ID);}
{num}	{yylval = (int) installNum(); return(NUMBER);}
"<"	{yylval = LT; return(RELOP);}
"<="	{yylval = LE; return(RELOP);}
"="	{yylval = EQ; return(RELOP);}
"<>"	{yylval = NE; return(RELOP);}
">"	{yylval = LT; return(RELOP);}
">="	{yylval = LT; return(RELOP);}

%%

int installID() {
	/*
	function to install the lexeme, whose first character is pointed
	to by yytext, and whose length is yyleng, into the symbol table and
	return a pointer thereto.
	*/
}

int installNum() {
	/*
	similar to installID, but puts numerical constants into a separate
	table.
	*/
}

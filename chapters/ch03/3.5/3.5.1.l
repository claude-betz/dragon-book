/* implement changes to Figure 3.23 for
	- 3.5.1 a
	- 3.5.1 b
	- 3.5.1 c
	- 3.5.1 d
*/

{%
	/* definitions of manifest constants
	LT, LE, EQ, NE, GT, GE,
	IF, THEN, ELSE, WHILE, STRING, ID, NUMBER, RELOP
	*/
%}

/* regular definition */
delim	[ \t\n]
ws	[delim]+
letter 	[A-Za-z_]
digit	[0-9]
id	{letter}({letter}{digit})*
number	{digit}+(\.{digit}+)?(E[+-]?{digit}+)?
string	\"(\\"|[^"])*\"

%%

{ws}		{/* no action and no return */}
if  		{return(IF);}
then		{return(THEN);}
else		{return(ELSE);}
whil		{return(WHILE);}
{id}		{yylval = (int) installID(); return(ID);}
{num		{yylval = (int) installNum(); return(NUMBER);}
{string}	{yylval = (int) installString(); return(STRING);}
"<"		{yylval = LT; return(RELOP);}
"<=		{yylval = LE; return(RELOP);}
"==		{yylval = EQ; return(RELOP);}
"!=		{yylval = NE; return(RELOP);}
">"		{yylval = LT; return(RELOP);}
">=		{yylval = LT; return(RELOP);}

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

int installString() {
	/*
	similar to above but separate table
	*/
}

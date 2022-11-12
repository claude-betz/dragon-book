/*
	Write a Lex program that copies a file, replacing each non-empty sequence of whitespace
	by a single blank.
*/

%{
#include<stdio.h>
#include<string.h>

char replace_with[] = " ";
%}

/* regular definitions */
space	[ ]
delim	[ \t\n]
ws	{delim}+

/* rules section */
%%
{ws}	{ fprintf(yyout, "%s", replace_with); } 
%%

int main(int k, char **argcv) {

	yyin=fopen(argcv[1], "r");
	yyout=fopen("3.5.1-test-result.txt", "w");
	yylex();

	return 0;
}

int yywrap() {
	return 1;
}

/*
	3.5.3 Write a Lex program that copies a C program, replacing each instance of the
	keyword "float" by "double"
*/

%{
#include<stdio.h>
#include<string.h>

char replace[] = "double";
%}

%% 

float	{ fprintf(yyout, "%s", replace); }

%%

int main(int k, char **argcv) {
	yyin = fopen(argcv[1], "r");	
	yyout = fopen("3.5.3-test-result.txt", "w");
	yylex();
	return 0;
}

int yywrap() {
	return 1;
}

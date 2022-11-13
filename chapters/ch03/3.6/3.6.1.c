/*
 * 3.6.1.c
 *
 * Show how, given the failure function for the KMP algorithm, we can construct from a 
 * keyword b_1, b_2, ..., b_n an n+1 state DFA that recognises .*b_1b_2...b_n, where the dot
 * stands for "any character". Moreover this DFA can be constructed in O(n) time.
 * */

#include<stdio.h>
#include<string.h>
#include<stdbool.h>

bool DFA(char text[], char keyword[], int failure_function[]) {
	int length_kw = strlen(keyword);
	int length_text = strlen(text);

	// states, each state stores character it accepts
	char states[length_kw+1];

	// states array
	memcpy(states+1, keyword, length_kw*sizeof(char));
	
	// counter for characters
	int i = 0;
	int curr_state = 0;
	while (i < length_text) {
		// if we can transition to next state
		if (text[i] == states[curr_state+1]) {
			curr_state = curr_state+1;
			i = i+1;
		} else {
			if (curr_state != 0) {
				curr_state = failure_function[curr_state];
			} else {
				i = i+1;
			}
		} 
	}

	printf("curr_state: %d, length_kw: %d", curr_state, length_kw);
	if (curr_state == length_kw) {
		return true;
	}
	return false;
}

int main() {
	char text[] = "hhdhaaa";
	char keyword[] = "aaa";
	int ff[4] = {0,0,1,2};

	bool res = DFA(text, keyword, ff);
	printf("%s", res ? "true":"false");
}

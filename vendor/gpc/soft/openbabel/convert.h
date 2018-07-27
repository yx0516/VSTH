#include <stdlib.h>   // go using C.free

struct state
{
	int         status;
	const char* msg;
	const char* data;
};

void ConvertMol(char* inMol, struct state* ret);

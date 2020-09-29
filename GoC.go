package GoC

/*
#include <stdio.h>

extern void GoPrint();
void Cprint() {
	char buffer[128];
	printf("This is Cprint!\n");
	scanf("%s",buffer);
	GoPrint(buffer);
}
*/
import "C"

func CallC() {
	C.Cprint()
}

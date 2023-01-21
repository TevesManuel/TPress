#include <stdio.h>
#include <TPress.h>

int main()
{
    const char * text;
    
    printf("Escribe el texto a simular: ");
    scanf("%[^\n]%*c", text);
    TypeText(text);
    return 0;
}
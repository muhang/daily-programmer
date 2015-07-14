#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main(int argc, char *argv[]) 
{
    char str[500];

    printf("Enter the string...\n");

    gets(str);

    /* SPLIT STRING BY WHITESPACE (http://stackoverflow.com/questions/11198604/c-split-string-into-an-array-of-strings) */
    char ** res  = NULL;
    char *  p    = strtok (str, " ");
    int n_spaces = 0, i;
/* split string and append tokens to 'res' */ 
    while (p) {
        res = realloc (res, sizeof (char*) * ++n_spaces);

        if (res == NULL) {
            exit (-1); /* memory allocation failed */
        }

        res[n_spaces-1] = p;

        p = strtok (NULL, " ");
    }

    /* realloc one extra element for the last NULL */

    res = realloc (res, sizeof (char*) * (n_spaces+1));
    res[n_spaces] = 0;

    int vert = 0;
    int left = 0;

    for (i = 0; i < n_spaces; i++) {
        
        if (strlen(res[i]) > 0 && res[i] != NULL) {

            if (vert == 1) {

                for(int k = 0; k < strlen(res[i]); ++k) {
                    if (k == 0) {
                        printf("%c\n", res[i][k]); 
                    } else {
                        for(int l = 0; l < left; ++l) {
                            printf(" ");
                        }
                        printf("%c\n", res[i][k]);
                    }
                } 
                vert = 0;
            } else {
                if (left > 0) {
                    for (int j = 0; j < left; ++j) {
                        printf(" ");
                    }
                }
                left += strlen(res[i]);
                
                printf("%s", res[i]);

                vert = 1;
            }
        }
    }

    free(res);
    
    return 0;
}

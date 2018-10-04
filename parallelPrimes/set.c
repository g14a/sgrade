
#include <stdio.h> 
#include <math.h> 

  
void printPowerSet(int *set, int set_size) 
{ 
    /*set_size of power set of a set with set_size 
      n is (2**n -1)*/
    unsigned int pow_set_size = pow(2, set_size); 
    int counter, j; 
    

    for(counter = 0; counter < pow_set_size; counter++) 
    { 
      for(j = 0; j < set_size; j++) 
       { 
          /* Check if jth bit in the counter is set 
             If set then pront jth element from set */
          if(counter & (1<<j)) 
            printf("%d", set[j]); 
       } 
       printf("\n"); 
    } 
 
} 
  
/*Driver program to test printPowerSet*/
int main() 
{ 
    int set[50];

    for (int i=0;i<50;i++) {
      set[i] = i;
    } 

    printPowerSet(set, 30); 

    return 0; 
}

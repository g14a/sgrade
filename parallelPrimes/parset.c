
#include <stdio.h> 
#include <math.h> 
#include <omp.h>
  
long double printPowerSet(int *set, int set_size) 
{ 
    /*set_size of power set of a set with set_size 
      n is (2**n -1)*/
    unsigned int pow_set_size = pow(2, set_size); 
    int counter, j; 
    
    long double start = omp_get_wtime();

#pragma omp parallel for
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

    long double end = omp_get_wtime()-start;

    return end;
} 
  
/*Driver program to test printPowerSet*/
int main() 
{ 
    int set[50];

    for (int i=0;i<50;i++) {
      set[i] = i;
    } 

    long double t = printPowerSet(set, 50); 

    printf("%Lf", t);
    return 0; 
}

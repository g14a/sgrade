# Intel Threading Challenge ( taken as S Grade )

### Problem -> [Intel site thread](https://software.intel.com/en-us/forums/p1-a2-consecutive-primes)

### The exact problem statement wasn't available on the thread, so the discussion forum was completely gone through and was implemented using [Go](https://golang.org/).

### Prime numbers are supposed to be generated within a limit. Then, sums of each subset of primes in range up to that prime are checked if they are possible perfect powers.

## Serial Approach
Doing it in the routine, serial way. Emitting prime numbers until a limit and then store them in a list. Then iterate through the list and apply the Perfect Power Algorithm given below.

## Parallel Approach

There are 3 filters.

1. Generating numbers from the lower limit to upper limit. 
2. Filtering the prime numbers is done parallely using [Channels](https://tour.golang.org/concurrency/2).
3. Computation of Perfect Powers is parallelized with respect to each individual independent power in the inner most for loop.

Read more about [Concurrency in Go with Goroutines](https://tour.golang.org/concurrency/1).
   
## Documentation of the Code.

Helper functions are put inside the /helpers directory of the root directory.

Helper functions include :

1.
        func Isprime() {
            
            // returns true if a number is prime.
        
        }

2.
    
        func Generate(start int, ch chan<- int, thresh int) {

            // inputs all numbers from start upto thresh in the channel ch

        } 

3. 
    
        func Filter(in <-chan int, out chan<-int, wg *sync.WaitGroup) {

            // filters elements from the 'in' channel into the 'out' channel if they're prime.

        }

4.
     
        func IsPower(num, power int) {

         // checks if num is a perfect power for an individual 'power'

        }


5.
        func GetNthroot(a, b int) {

            //returns the root of the perfect power found above.

        }


6.
        func FileCreate(name string) *os.File {

            // creates a file with the name 'name'

        }


Read more about [sync.WaitGroups](https://golang.org/pkg/sync/#WaitGroup).

## The Perfect Power Computation

[Algorithm](https://github.com/gowtham-munukutla/sgrade/blob/a0fd40f0891f10fcf7480f7ef2ccd9f31c13fb37/parallelPrimes/gos/par1.go#L64).

The parallelized part is mentioned here below.

```
    var wg sync.WaitGroup
    
    go func(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	    for power := 2; power <= maxPower; power++ {
		if helpers.IsPower(sum, power) {
    		s := strconv.Itoa(prime) + ":" + strconv.Itoa(primes[elem]) + " = " + strconv.Itoa(sum) + " = " + strconv.Itoa(helpers.GetNthroot(sum, power)) + "**" + strconv.Itoa(power)
				io.WriteString(fp, s+"\n")
	        }
	    }
    }(&wg)
```

## Performance Statistics

|Limit   | Power Limit | Serial Time in Sec |  Parallel Time in Secs |
|--------|-------------|--------------------|------------------------|
|20000	 |5	           |1.595569219         |      1.828454819       |
|20000	 |10	       |2.63408235          |      1.857581038       |
|20000	 |15	       |3.672684378         |      1.874498921       |
|20000	 |20	       |4.708801813         |      1.911036053       |
|20000	 |22	       |5.125772613         |      1.914186214       |
|20000	 |25	       |5.747006907         |      1.923525398       |
|20000	 |30	       |6.78421919          |      1.980672387       |
|20000	 |35	       |7.819467612         |      2.024711574       |
|20000	 |50	       |10.93049763         |      2.323714389       |


|Limit   | Power Limit | Serial Time in Sec |  Parallel Time in Secs |
|--------|-------------|--------------------|------------------------|
|50000	 |5	           |13.422268529        |      15.638489655      |
|50000	 |10	       |18.756378307        |      15.781230999      |
|50000	 |15	       |24.103983503        |      16.005617183      |
|50000	 |20	       |29.43290857         |      16.231497568      |
|50000	 |22	       |31.575538906        |      16.321308708      |
|50000	 |25	       |34.78270055         |      16.48725101       |
|50000	 |30	       |40.121898136        |      16.764135946      |



|Limit   | Power Limit | Serial Time in Sec |  Parallel Time in Secs |
|--------|-------------|--------------------|------------------------|
|100000	 |5	       |65.122981912        |      88.069495932      |
|100000	 |10	       |93.743481887        |      88.86182371       |
|100000	 |15	       |112.387322679       |      89.784588735      |
|100000	 |18	       |123.646906166       |      90.600418026      |
|100000	 |20	       |131.040610217       |      91.138923087      |
|100000	 |22	       |138.474489187       |      91.779524448      |
|100000	 |25	       |149.773055499       |      92.876253655      |
|100000	 |30	       |168.309619647       |      93.965321005      |
|100000	 |32	       |175.769323417       |      93.312082881      |
|100000	 |35	       |186.921259909       |      93.312082881      |


## Performance Graphs

![Prime Generation](https://i.imgur.com/5D9PDNC.png)

### With various Limits
![](https://i.imgur.com/h5IhJTP.png)
![](https://i.imgur.com/M1hTj9x.png)
![](https://i.imgur.com/rDxW3r1.png)

# go-tsp
**Travelling salesman problem solved with Genetic Algorithm using Go**.
  
## How to run
The program can be ran 2 ways
1. go run main.go  
or
2. go build : an executable file is generated  
**ie:** go-tsp.exe. Run the file.   

## Results
The program will create a folder called **tsp**(*can be defined in program),
 which stores the results based on the seed number provided. The results are 
 images(.png) of the tours that outperform their predecessors.  
**ie:** ./ tsp / [seed number] / [images....png]  
    
**Sample Results:**  
![Alt text](/ReadmeAssets/1504372704/output.gif)

## Documentation
The program can be tweaked by modifying the global variables in 
- main.go
- geneticAlgorithm/geneticAlgBase.go

## Resources
- [Thorough read on genetic algorithm on "tutorialspoint"](https://www.tutorialspoint.com/genetic_algorithms/index.htm) 
- [Java implementation on "theprojectspot"](http://www.theprojectspot.com/tutorial-post/applying-a-genetic-algorithm-to-the-travelling-salesman-problem/5)
- [Additional reading material on "analyticsvidhya](https://www.analyticsvidhya.com/blog/2017/07/introduction-to-genetic-algorithm/)
A Bayes Filter

Things to do.

Design:
	Eventually this software will have these classes

	Source
		- a class that can iterate on the textfiles in a directory (we will exclude binary files, subdirectories and hidden content) or a single file, or a single user-provided byte array

	Tokenizer
		- it will be used to take the entire file or start 
		- it will tokenize the 
			- if you provide a stopword list you will 

	Learning Functions	
		- Naive Bayes Trainer -> nuff said
		- Stopword Creator -> it will generate a stopword list that you can use on the trainer and the filter

	Filter Functions
		- Naive Bayes Filter
		- Print-To-Screen
	
Steps

1. Write it
	a. ~~read a file~~
	b. create a printfile function
		b1. Make printfile work by using a Source Class
	c. create a learn function
		c1. ~~create a fulltext vs mail format switch~~
		c2. parse on a directory
		c3. record status in a file
	d. create a scoring function
	e. test this stuff with some spam testcase

2. Put it on github



# Gofilter
A Naive Bayes Filter

This software is unfinished.

# Design
Eventually this software will have these classes

##SourceDocuments
A class that will iterate on the textfiles in a directory (I will exclude subdirectories and files starting with .), a single file, or a single user-provided byte array.

##Tokenizer
It will be used to take the entire file or start after the headers section (if you provide it with an RFC2822 email message).
It will tokenize the words inside the file, providing both words and their occurrence.
If you provide a stopword list it will remove them out.

##Learning Functions	
* Naive Bayes Trainer
	'nuff said
* Stopword Creator 
	it will generate a stopword list that you can use on the trainer and the filter

##Filter Functions
* Naive Bayes Filter
* Print-To-Screen for debugging purposes
	
Steps

1. Write it
  1. ~~read a file~~
  2. create a printfile function
	1. Make printfile work by using a Source Class
  3. create a learn function
 	 1. ~~create a fulltext vs mail format switch~~
 	 2. parse on a directory
  4. record status in a file
  5. create a scoring function
  6. test this stuff with some spam testcase

2. ~~put it on github~~


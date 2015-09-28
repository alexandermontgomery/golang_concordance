# Requirements
This was built and tested on Golang 1.3. I will leave the setup of a golang environment out of the scope of this document. You can read more about installing go on the [go website](https://golang.org/doc/install). 

I always like to use a virtualization environment such as Docker instead of installing software directly on my machine. [https://hub.docker.com/_/golang/](This is the official go docker repo)
# Building the Application
After setting up your Go environment and checking out this repo, simply run 

    go install golang_concordance

# Testing the Application
To run the tests run

    go test golang_concordance
    
# Running the Application
Currently the application only receives input from stdin. The tool to pipe the text to the built executable is at the users discretion. Here is an example using [cat](http://www.linfo.org/cat.html)

    cat some_file_with_text.txt | golang_concordance

# Assumptions
* Only romance languages will be fed to this program
* Different formats of the same dictionary word (e.g. "labels", "label", "label's") are different words in the concordance
* A sentence structure is delimted by the following pattern: punctuation, white space delimiter, a capital letter

# Exceptions
* There can only be 4294967295 words in the inputted text
* There can only be 4294967295 sentences in inputted text
* Certain sentence structures cannot be handled by the sentence splitting rule above. For example, "There are many states with large cities that are not the capital city, e.g. New York City". A human recognizes this as one sentence but the program will split on the ". N" after "e.g."
* Dashes between words need to be surrounded by spaces otherwise they will be viewed as one word. For example, "The famous Merriam-Webster dictionary was written by two individuals. Publication rights to the dictionary was obtained in 1843 by Merriam - Webster died shortly before that".
* Certain words that seem like words to a human (for instance the programming language C++) are not covered by the regex. It would be trivial to add additional characters to the regex, but by expanding the regex, the edgecases will also expand

# Roadmap
* Consider parallelizing the processing of strings using a concurrency framework
* Build in a mechanism to better manage memory. Most likely use some sort of data store such as Redis to hold concordance calculations while building a Concordance for extremely large texts.
* Allow the various parsing regexes to be configured at runtime via command line arguments
* I used a regex for this. I generally do not like to use regexes because at times their behavior can be confusing. With more time I would have further tested and honed the regexes, and considered finding a development strategy that does not use a regex
* I would increase usability by allowing a handful of different places to specify the input text. Most likely, a file, a resource URL, as well as the command line.
* I would allow users to choose their desired output format (JSON, XML, plain text to CLI, HTML page)

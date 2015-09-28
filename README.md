# Building 

# Running

# Testing

# Assumptions
* Only romance languages will be fed to this program
* Different formats of the same dictionary word (e.g. "labels", "label", "label's") are different words in the concordance
* A sentence structure is delimted by the following patter: punctuation, white space delimiter, a capital letter

# Exceptions
* There can only be 4294967295 words in the inputted text
* There can only be 4294967295 sentences in inputted text
* Certain sentence structures cannot be handled by the sentence splitting rule above. For example, "There are many states with large cities that are not the capital city, e.g. New York City". A human recognizes this as one sentence but the program will split on the ". N" after "e.g."
* Dashes between words need to be surrounded by spaces otherwise they will be viewed as one word. For example, "The famous Merriam-Webster dictionary was written by two individuals. Publication rights to the dictionary was obtained in 1843 by Merriam - Webster died shortly before that".
* Certain words that seem like words to a human (for instance the programming language C++) are not covered by the regex. It would be trivial to add additional characters to the regex, but by expanding the regex, the edgecases will also expand

# With more time
* Consider parallelizing the processing of strings using a concurrency framework
* Allowed the various constants to be configured at runtime via command line
* I used a regex for this. I generally do not like to use regexes because at times their behavior can be confusing. With more time I would have further tested and honed the regexes, and considered finding a development strategy that does not use a regex
* I would increase usability by allowing a handful of different places to specify the input text. Most likely, a file, a resource URL, as well as the command line.
* I would have outputted an interactive HTML page of the concordance so users can see the sentences associated with a word
#                   PROJECT DOCUMENTATION

This project is based on Go Language, and it's a text transforming program that modifies text that is poorly structured into a proper structured text.

------------------------------------
### FEATURES OF THIS PROGRAM

1. **Text Transforming**: it has features like fixing poorly writing punctuations, quotes, articles and gives the user a power to change the casing of any words into an UPPERCASE, lowercase or Titlecase and also have addition features of coverting the case of multiple text too by using some special markers e.g (up) for uppercasing, (low) for lowercasing and (cap) for Tile case, and adding number inside the maker bracket allowes the user to transform mulitiple text together with that marker.

2. **Base Transforming**: this feature allow the user to convert base text (e.g Binary and Hexadecimal) into a decimal base with special markers like (bin) and (hex), transforming the text before the marker into decimal

## FILES FUNCTION

 |*FILES* | *CATEGORY* | *USES* |
 |------|----------|------|
 |articles.go| helper function| it has a function that fix articles that that are poorly structured with vowel converting A to An when the next word begins with a vowel and converts An to a when the next word begins with a consonant letter.
 |base.go| helper function| it has a function that converts text written in base form of binary or hecxadecimal to a decimal base using a marker (bin) or (hex) manipulating the previous word before the marker into the decimal base.
 |casing.go| helper function| it has a function that changes text cases of words using a marker of (up) for *UPPERCASE*, (low) for *lowercase* and (cap) for *Titlecase* by changing the case of the word before the marker... the marker can also contain numbers in it (e.g (up, 3)) using the number to manipulate all the words before the marker falling within the count of the index.
 |complier.go| helper function| it has the Orchestrator function that coordinate and arrange all the functions in a single function to be easily called into the main (i.e helpers.Compiler).
 |punct.go| helper function| it has a function to fix punctuations in text that are poorly written.
 |quotes.go|helper function| it has a function to fix single or double quotes in text that are poorly witten.
 |main.go| generic function| it's the heart and engine of the program it controls the program and make it executable through the function main, it also possess a *os* read and write function to read user input, capture the data and store it in a memory then apply all the functions and rewrite it into an output file, clean and refined.
 |go.mod| Go modules| this a file for root of Go dependency system, it identifies the program as a module which is collection of GO packages together and keeps track of every external packages this program use.
 |sample.txt| text file (input)| it's a file where the user input their text for it to be refined.
 |result.txt| text file (output)| it's a file where the user sees the modified output of the initial input, cleaned and refined.
 -------------------------------





### HOW IT WORKS

This program carries two text files (sample.txt and result.txt)... The sample.txt accepts the text from the user, structured or poorly structured and to run the program the user uses a command in the terminal which is *(go run . sample.txt result.txt)* after hitting ENTER, the inputed text passes through the program functions(i.e. helper functions) and it identifies any poorly structured text and it polish or refine it into a modified text, to see the result the user can go into the result.txt file or type in a command to see the result directly in the terminal which is (cat result.txt)










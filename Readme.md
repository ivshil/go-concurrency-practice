# Project created to practice Concurency in GoLang

- Simple
  
        Create a function called "printMessage" that takes a string as input and prints it to the console.

        In the main function, create a goroutine that invokes the "printMessage" function with the argument "Hello, World!".

        After starting the goroutine, print the message "Goroutine started!" to the console from the main function.

        Use a delay or wait mechanism to ensure that the goroutine has enough time to execute before the main function terminates.

        Verify that the output displays "Goroutine started!" before the message "Hello, World!".

        Example Output: Goroutine started! Hello, World!

-  Practical Concurency
  
        Define a function called "processFile" that takes a filename (string) as input and simulates some processing on the file. For this example, you can simply print the filename and a message indicating that processing is being done.

        In the main function, create a slice of filenames (at least 5 filenames) that represent files to be processed concurrently.

        Use a loop to iterate through the slice of filenames.

        For each filename, launch a goroutine that calls the "processFile" function with the respective filename as an argument.

        Use a delay or wait mechanism (such as time.Sleep()) to ensure that all goroutines have enough time to complete their processing before the main function terminates.

        Example Output: Assuming the filenames in the slice are ["file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"], the output could be:

        Processing file: file1.txt Processing file: file2.txt Processing file: file3.txt Processing file: file4.txt Processing file: file5.txt

- Note:

        You can modify the "processFile" function to include more meaningful processing, such as reading the contents of the file, performing calculations, or any other desired operations. The actual execution order and interleaving of the messages may vary due to the concurrent nature of goroutines.

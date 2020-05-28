This is a simple program that compares the content of two files and prints out the percentage they match. This program was done for Tero Karvinen's Go-Programming course.
Link to course page: http://terokarvinen.com/2020/go-programming-course-2020-w22/

=== Made by Arttu Pesonen ===

You can run this program in Linux by typing ./filematcher --f1 *file1 name* --f2 *file2 name*
You can add --i to enable img mode (Work in progress. Matches 2 same images 100% but still testing left to do for two similar images)
Supported image types are PNG and JPEG/JPG
In Windows just type .\filematcher but otherwise the syntax is the same.

Sources I that helped me greatly:
- http://terokarvinen.com/2020/go-programming-course-2020-w22/
- https://golang.org/pkg/image/
- https://gobyexample.com/reading-files
- https://www.teachwithict.com/binary-representation-of-images.html
- https://socketloop.com/tutorials/golang-convert-an-image-file-to-byte
- https://golang.org/pkg/bufio/
- https://github.com/nfnt/resize
- https://pkg.go.dev/github.com/nfnt/resize?tab=doc
- https://stackoverflow.com/questions/22945486/golang-converting-image-image-to-byte
- https://golang.org/pkg/image/png/
- https://stackoverflow.com/questions/46437169/png-encoding-with-go-is-slow

I am still not sure what percentage of similarity could be considered for example as a sign of plagiarism, or img identification since I haven't tested enough files yet.
With txt files where one had "aaaa" and the second had "aaab" the program printed out 75% so at least the math should be correct.

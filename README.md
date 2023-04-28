# Extended library

This project tries to extend the standard library of Go with some useful functionalities.
It is composed of multiple packages, each of them with a different purpose.

Currently, I don't have plans to extend it further, but any suggestions are welcome.

## [Data structures](datastructures/README.md)
The initial idea was to organize commonly used data structures in one place which would
implement common interfaces.
In most Go codes built-in types (slices and maps) are enough, 
but sometimes you need more specific structures (linked list, sets, etc.).

## [Streams](streams/README.md)
Later on, the project was extended with support for streams. 
They can be used completely independently of the rest of the project, but my data 
structures are also compatible with them.

Unfortunately, I wasn't able to implement the stream in a way that I wanted to.
The problem was that Go doesn't support generic methods, so methods like `Map` or `Reduce`
needed to remain as generic functions.
All functions are usable, but they aren't readable as in Java.
(I hope that Go will support generic method in the future)

## [Misc](misc/README.md)
In the meanwhile, a package with commonly used types and function was created.
Things that don't belong to any other package, but you probably need in every project.
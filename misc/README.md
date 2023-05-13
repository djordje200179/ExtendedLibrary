# Misc

A set of types and functions that don't fit into other packages and which
are useful for nearly any project.

## Optional
This package provides a type `Optional` which is a wrapper around a value
that may or may not be present. 
It is similar to Java's `Optional` or Ruby's `Maybe` type.

Alternative to using this package is storing the value as a pointer and
checking for `nil` value, but that approach increases number of 
memory indirections and risks null pointer dereference.

## Functions
This package provides a set of common function and function types.
Some of them were created just for my own needs, but most of them
are needed for "every day" programming (predicates, comparators, etc).

## Pair
Simple generic type that holds two values of any type.

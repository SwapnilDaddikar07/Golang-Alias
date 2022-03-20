# Golang-Alias


Problem statement.

Our application has a package named Animals.
This package has concrete implementations for the Animal Interface which resides in the Interfaces package.
The Animals package like mentioned contains concrete implementations like Tiger and Lion. But the catch is , it also contains an Error type.
The Animal interface returns the above mentioned error.
During refactoring,  the Error from Animals package was moved to Errors Package , as that is the right place for the error to reside.
Now , we have 2 errors , 1 in Errors package and one in animals package.
Our interface returns error from the Animals package, now if we want to move to the new implementation that we have in Errors package , we need to change the interface.
If we do that , there is no gradual change, unlike constants , variables and functions which can be changed gradually. (old implementation internally calling the new one.)
So what is the way out?
We can move the implementation to Errors package, and say that the old error in Animals package is same as the error in Errors package.
We can do this by  
type Error = Errors.AnimalErrors

Above is called type alias.

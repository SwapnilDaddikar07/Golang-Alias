package Animals

import "alias/Errors"

//type AnimalErrors struct {
//	ErrorCode    int
//	ErrorMessage string
//}
//
//func (e AnimalErrors) SetError() {
//	e.ErrorMessage = "error message"
//}

//The above is the first iteration/first implementation. AnimalErrors was in the Animals package.
//We then decided , we need to move error out from Animals package and add to the error package. (code cleaning/refactoring.)
//The issue is , the animal interface returns Animals.Error i.e the above error.
//We cannot change the interface implementation , because all implementations will be impacted (type is different even if underlying type is same) ,and we need to change all code at once.
//Even if we do the below mentioned change , where we assign the old Error to the Errors.AnimalErrors, this will still not help solve incremental changes.
//As soon as we change the interface, evertyhing has to be changed.

//type Error Errors.AnimalErrors

//So this is the preferred way using type alias, we say that the old Error is exactly same as the new Error in errors package.
//We now can change interfaces and implementations incrementally. and move to the new error.

//Try to directly change the interface from returning Animals.Error to Errors.Error , there is no gradual change , all implementations will have to change together.
//Now , use the below method , where u assign the old error to the new error.
//Now interface can be chnaged to return new Errors.Error , and it wont cause any issue on the implementation. The implementations can now gradually start returning Errors.Error and then finally sunset the AnimalErrors.Error

type Error = Errors.AnimalErrors
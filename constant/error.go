package constant

type Error string

const (

	FailedToParseData				Error = "Oops! We encountered an issue while processing your data. Please try again!"
	FailedToHashPassword			Error = "We ran into an error while securing your password. Please try again!"
	FailedToInsertData      		Error = "Sorry, we couldn’t add the new data. Please try again!"
	FailedToLoadUserData    		Error = "We couldn’t load your account information. Please try again!"
	FailedToGenerateTokenAccess  	Error = "We couldn’t generate an access token. Please try again!"
	InvalidAccountError          	Error = "The email or password entered is incorrect. Please try again!"
	DuplicateDataError				Error = "This email is already taken. Please choose another one!"

)
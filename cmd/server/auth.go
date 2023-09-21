package server

type AuthController struct {
}

func InitAuthController() *AuthController {
	return &AuthController{}
}

/*TODO:
User requests a challenge
	Gets challenge from Auth service
		Returns user with public key of node, encrypted message and the nonce.
	User tries to decrypt the challenge by generating a shared key using public key of node and their private key
	User then sends this deciphered text back.
		Auth verifies the deciphered key with the key it has.
		if it matches, user is authenticated.

TODO: What happens when user is authenticated?
	How does the signin/signup works?
*/

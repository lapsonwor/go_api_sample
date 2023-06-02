package goEth

import (
	"fmt"
	"testing"
)

func TestVerifySign(t *testing.T) {
	publicAddr := "0x10Af5c28EaA42765e319b73eAEa60b4F668aA3d0"
	message := `E2hWNKKHg1uYOEYnw+F/y0h7Q+YreWnV5jnwhA+nLBnwPyNKxWa6noOIsc7xKAhvA5woQ6R5b7yfa/rjDWZ8Cp+w4t78J+kWzTohjOV1P2I3fFkb9WDhvqXHYWo6/aQXpnkeozPaaKuhzRUd329B07vk2Y8qPIdmdvFDkzaz+L0MtVsdktzGNeWe70aZ4PFsk7u7qsp5Mrey9z43zoZFs52di6gW/z+fjt0xilZz06ZYeM/l1UYXvPr/U1UOfO0UyvgSUSnaIA/IXPxb+hsKn1cDciVksS9AFqY7fkPYR8hmBhS+ApIYDpb5OIlQDl+gqvJUmkQCqxXIfnx4wgVk1g==`
	signatureHash := "0xaab8462591f04588d1f6452b676d268339c9d575f967d16bd00e04823c59983913489514dd5efb8fe8b9b564f09011987d994522648ee7a0fef248296fce146a1b"
	match := VerifySign(publicAddr, signatureHash, message)
	fmt.Println("match", match)
}
func TestVerifySignEg(t *testing.T) {
	// Establish the parameters to use for Argon2.
	privateKey := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	message := `goQ4aWonIO1KcJ6WzUDMq5pwim9jefouMg49pskOfjomWFBzqWluaVYbuXmj1qz/bnNJw/NGd+w/MAH2gsOtn70qoC88zqbfpKy1Lw/kYO7aTo4EXj/3v6CM9ZgsIrx/DFeUh2uJxY4aSwLLtVHrC+VUWCz/NCMt5e+XWp4co2k4vvsgTRK1UvvPAPEH82x0N85inJ7XG1EPhQE67R7RrCBJzY6iw1yE9bgSzwSaXC2Hvlvrk+JpRSrKJCeLTVNRQn+1gobzyrTl91EAbJGkvmj/IAUtMl6Icv3a7jwcb4vk+D6wxwINgBZJ2EYJ3Z17D3G6LcVp2oVMd/3ueoU3bw==`

	verified := VerifySignWithPrivate(privateKey, message)
	fmt.Println("verified", verified)
	fmt.Println("Test completed")

}
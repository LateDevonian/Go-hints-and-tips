
//zak's workshop calculator for my reference
package tshirtcalculator

// Size represents a T-Shirt's size
type Size int

const (
    // S represents the sizing for a small T-Shirt
    S Size = 0

    // M represents the sizing for a medium T-Shirt
    M Size = 1

    // L represents the sizing for a large T-Shirt
    L Size = 2
)

// Color represents
type Color int

const (
    // Beige is probably the best color
    Beige Color = 0

    // Khaki is a knockoff beige
    Khaki Color = 1

    // Cream is a tasteful color
    Cream Color = 2

    // CosmicLatte would be a good name for a cat. See https://en.wikipedia.org/wiki/Beige
    CosmicLatte Color = 3

    // UnbleachedSilk is perhaps the most descriptive name for a color ever.
    UnbleachedSilk Color = 4

    // Blue is a pretty cool color
    Blue Color = 5
)

const (
    // BasePrice is the base price for a non large, blue and text covered T-Shirt
    BasePrice = 3.50

    // TextFee is the fee for having text printed on the shirt
    TextFee = 5.50

    // BlueFee is the extra fee for having a blue shirt
    BlueFee = 10.0
)

// CalculatePrice generates the price of a T-shirt based on specifications
func CalculatePrice(size Size, color Color, text string) float64 {
    price := BasePrice
    if text != "" {
        price += TextFee
    }

    if color == Blue {
        price += BlueFee
    }

    return price
}
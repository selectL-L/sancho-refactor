package utils

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	mathrand "math/rand"
	"strings"
	"time"
)

// RandomUtils provides utilities for generating random values
type RandomUtils struct {
	rng *mathrand.Rand
}

// NewRandomUtils creates a new RandomUtils instance
func NewRandomUtils() *RandomUtils {
	return &RandomUtils{
		rng: mathrand.New(mathrand.NewSource(time.Now().UnixNano())),
	}
}

// GetRandomInt generates a random integer between min and max (inclusive)
func (r *RandomUtils) GetRandomInt(min, max int) (int, error) {
	if min >= max {
		return 0, fmt.Errorf("min must be less than max")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + min, nil
}

// GetRandomIntInsecure generates a random integer using math/rand (less secure but faster)
func (r *RandomUtils) GetRandomIntInsecure(min, max int) int {
	return r.rng.Intn(max-min+1) + min
}

// GetRandomFloat generates a random float between min and max
func (r *RandomUtils) GetRandomFloat(min, max float64) (float64, error) {
	if min >= max {
		return 0, fmt.Errorf("min must be less than max")
	}

	// crypto/rand doesn't have a Float64() function, so we need to create our own
	bigInt, err := rand.Int(rand.Reader, big.NewInt(1<<53))
	if err != nil {
		return 0, err
	}

	// Convert to a float64 between 0 and 1
	f := float64(bigInt.Int64()) / float64(1<<53)

	return min + f*(max-min), nil
}

// GetRandomFloatInsecure generates a random float using math/rand (less secure but faster)
func (r *RandomUtils) GetRandomFloatInsecure(min, max float64) float64 {
	return min + r.rng.Float64()*(max-min)
}

// RollDice rolls a dice with the given number of sides
func (r *RandomUtils) RollDice(sides int) (int, error) {
	if sides < 1 {
		return 0, fmt.Errorf("dice must have at least 1 side")
	}

	return r.GetRandomInt(1, sides)
}

// ParseDiceRoll parses a dice roll expression (e.g. "2d6+3")
func (r *RandomUtils) ParseDiceRoll(expr string) (int, []int, error) {
	// Remove all spaces
	expr = strings.ReplaceAll(expr, " ", "")

	// Check if it's just a number
	if num, err := parseSingleNumber(expr); err == nil {
		return num, []int{num}, nil
	}

	// Split by operators
	operators := []string{"+", "-", "*", "^", "_"}
	var op string
	var parts []string

	for _, operator := range operators {
		if strings.Contains(expr, operator) {
			op = operator
			parts = strings.SplitN(expr, operator, 2)
			break
		}
	}

	// If no operator found, try to parse as a dice roll
	if op == "" {
		return r.parseDiceNotation(expr)
	}

	// Otherwise, evaluate each part
	leftVal, leftDice, err := r.ParseDiceRoll(parts[0])
	if err != nil {
		return 0, nil, err
	}

	rightVal, rightDice, err := r.ParseDiceRoll(parts[1])
	if err != nil {
		return 0, nil, err
	}

	// Combine dice results
	allDice := append(leftDice, rightDice...)

	// Apply operator
	result := 0
	switch op {
	case "+":
		result = leftVal + rightVal
	case "-":
		result = leftVal - rightVal
	case "*":
		result = leftVal * rightVal
	case "^":
		result = int(math.Pow(float64(leftVal), float64(rightVal)))
	case "_":
		result = int(math.Pow(float64(rightVal), float64(leftVal)))
	}

	return result, allDice, nil
}

// parseDiceNotation parses a dice notation expression (e.g. "2d6")
func (r *RandomUtils) parseDiceNotation(expr string) (int, []int, error) {
	parts := strings.Split(expr, "d")
	if len(parts) != 2 {
		return 0, nil, fmt.Errorf("invalid dice notation: %s", expr)
	}

	// Parse number of dice
	var count int
	if parts[0] == "" {
		count = 1
	} else {
		var err error
		count, err = parseSingleNumber(parts[0])
		if err != nil {
			return 0, nil, fmt.Errorf("invalid dice count: %s", parts[0])
		}
	}

	// Parse number of sides
	sides, err := parseSingleNumber(parts[1])
	if err != nil {
		return 0, nil, fmt.Errorf("invalid dice sides: %s", parts[1])
	}

	// Validate
	if count < 1 {
		return 0, nil, fmt.Errorf("dice count must be at least 1")
	}
	if sides < 1 {
		return 0, nil, fmt.Errorf("dice sides must be at least 1")
	}

	// Roll the dice
	sum := 0
	rolls := make([]int, count)

	for i := 0; i < count; i++ {
		roll, err := r.RollDice(sides)
		if err != nil {
			return 0, nil, err
		}

		rolls[i] = roll
		sum += roll
	}

	return sum, rolls, nil
}

// parseSingleNumber parses a single numeric value
func parseSingleNumber(s string) (int, error) {
	var result big.Int
	_, ok := result.SetString(s, 10)
	if !ok {
		return 0, fmt.Errorf("invalid number: %s", s)
	}

	return int(result.Int64()), nil
}

// ComposeRollResult formats a dice roll result
func (r *RandomUtils) ComposeRollResult(total int, dice []int) string {
	if len(dice) == 1 && dice[0] == total {
		return fmt.Sprintf("%d", total)
	}

	diceStr := make([]string, len(dice))
	for i, d := range dice {
		diceStr[i] = fmt.Sprintf("%d", d)
	}

	return fmt.Sprintf("%d (%s)", total, strings.Join(diceStr, " "))
}

// ProcessRollCommand parses and executes a roll command
func (r *RandomUtils) ProcessRollCommand(input string) (string, error) {
	// Extract the dice expression
	expr := strings.TrimPrefix(input, ".roll ")
	expr = strings.TrimSpace(expr)

	// Parse and roll
	total, dice, err := r.ParseDiceRoll(expr)
	if err != nil {
		return "", err
	}

	// Format the result
	result := r.ComposeRollResult(total, dice)

	// Check if the output would be too long
	if len("Your roll is "+result+".") > 2000 {
		return fmt.Sprintf("%d", total), nil
	}

	return result, nil
}

// GetRandomElement selects a random element from a slice
func (r *RandomUtils) GetRandomElement(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("slice is empty")
	}

	index, err := r.GetRandomInt(0, len(slice)-1)
	if err != nil {
		return "", err
	}

	return slice[index], nil
}

// GetRandomElementInsecure selects a random element using math/rand (less secure but faster)
func (r *RandomUtils) GetRandomElementInsecure(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	return slice[r.rng.Intn(len(slice))]
}

// GetWeightedRandom selects a random element based on weights
func (r *RandomUtils) GetWeightedRandom(options []string, weights []float64) (string, error) {
	if len(options) != len(weights) {
		return "", fmt.Errorf("options and weights must have the same length")
	}

	if len(options) == 0 {
		return "", fmt.Errorf("options is empty")
	}

	// Calculate sum of weights
	var sum float64
	for _, w := range weights {
		sum += w
	}

	// Get a random value between 0 and sum
	n, err := r.GetRandomFloat(0, sum)
	if err != nil {
		return "", err
	}

	// Find the corresponding option
	var cumulativeWeight float64
	for i, w := range weights {
		cumulativeWeight += w
		if n <= cumulativeWeight {
			return options[i], nil
		}
	}

	// Fallback to last option
	return options[len(options)-1], nil
}

// GetDeterministicRandom gets a deterministic "random" value based on input
func (r *RandomUtils) GetDeterministicRandom(input string, salt string, max int) int {
	// Create a deterministic seed from the input and salt
	seed := int64(0)
	for i, c := range input {
		seed += int64(c) * int64(i+1)
	}
	for i, c := range salt {
		seed += int64(c) * int64(i+1)
	}

	// Create a deterministic RNG
	rng := mathrand.New(mathrand.NewSource(seed))

	return rng.Intn(max)
}

// PseudoRandomResponse gets a consistent response for 8ball-like commands
func (r *RandomUtils) PseudoRandomResponse(question string, userID string, responses []string) string {
	today := time.Now().YearDay() + time.Now().Year()*365

	// Generate deterministic index
	index := r.GetDeterministicRandom(strings.ToLower(question), fmt.Sprintf("%d%s", today, userID), len(responses))

	return responses[index]
}

// Shuffle randomizes the order of elements in a slice
func (r *RandomUtils) Shuffle(slice []string) error {
	for i := len(slice) - 1; i > 0; i-- {
		j, err := r.GetRandomInt(0, i)
		if err != nil {
			return err
		}
		slice[i], slice[j] = slice[j], slice[i]
	}
	return nil
}

// ShuffleInsecure randomizes the order of elements in a slice using math/rand
func (r *RandomUtils) ShuffleInsecure(slice []string) {
	r.rng.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// GetRandomBytes generates random bytes using crypto/rand
func (r *RandomUtils) GetRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateToken creates a random token of specified length
func (r *RandomUtils) GenerateToken(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		n, err := r.GetRandomInt(0, len(chars)-1)
		if err != nil {
			return "", err
		}
		result[i] = chars[n]
	}

	return string(result), nil
}

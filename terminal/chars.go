package terminal

var ALPHABET = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var NUMBERS = []byte("0123456789")
var SPECIAL_CHARS = []byte("!@#$%^&*()_+-=[]{}|;:,.<>?~")
var ALL_CHARS = append(ALPHABET, NUMBERS...)

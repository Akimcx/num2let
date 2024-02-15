package main

import (
	"testing"
)

type testSuit struct {
	input int; expected string
}

func TestSingleDigit(t *testing.T) {
	tests := []testSuit {
		{0, "zéro"},
		{1, "un"},
		{2, "deux"},
		{3, "trois"},
		{4, "quatre"},
		{5, "cinq"},
		{6, "six"},
		{7, "sept"},
		{8, "huit"},
		{9, "neuf"},
	}

	test(tests,t)
}



func TestMultiplesOfTen(t *testing.T) {
	tests := []testSuit{
		{10, "dix"},
		{20, "vingt"},
		{30, "trente"},
		{40, "quarante"},
		{50, "cinquante"},
		{60, "soixante"},
		{70, "soixante-dix"},
		{80, "quatre-vingts"},
		{90, "quatre-vingt-dix"},
	}

	test(tests,t)
}

func TestTeens(t *testing.T) {
	tests := []testSuit{
		{11, "onze"},
		{12, "douze"},
		{13, "treize"},
		{14, "quatorze"},
		{15, "quinze"},
		{16, "seize"},
		{17, "dix-sept"},
		{18, "dix-huit"},
		{19, "dix-neuf"},
	}

	test(tests,t)
}

func TestTwoDigitNumbers(t *testing.T) {
	tests := []testSuit{
		{21, "vingt et un"},
		{34, "trente-quatre"},
		{48, "quarante-huit"},
		{57, "cinquante-sept"},
		{69, "soixante-neuf"},
		{73, "soixante-treize"},
		{88, "quatre-vingt-huit"},
		{94, "quatre-vingt-quatorze"},
	}

	test(tests,t)
}

func TestThreeDigitNumbers(t *testing.T) {
	tests := []testSuit{
		{100, "cent"},
		{234, "deux cent trente-quatre"},
		{567, "cinq cent soixante-sept"},
		{789, "sept cent quatre-vingt-neuf"},
		{911, "neuf cent onze"},
		{999, "neuf cent quatre-vingt-dix-neuf"},
	}

	test(tests,t)
}

func TestFourDigitNumbers(t *testing.T) {
	tests := []testSuit{
		{1000, "mille"},
		{2001, "deux mille un"},
		{3456, "trois mille quatre cent cinquante-six"},
		{6789, "six mille sept cent quatre-vingt-neuf"},
		{9123, "neuf mille cent vingt-trois"},
		{9999, "neuf mille neuf cent quatre-vingt-dix-neuf"},
	}

	test(tests,t)
}

// Todo add support for 1_000_000_000 numbers
func TestLargeNumbers(t *testing.T) {
	tests := []testSuit{
		{10000, "dix mille"},
		{56789, "cinquante-six mille sept cent quatre-vingt-neuf"},
		{123456, "cent vingt-trois mille quatre cent cinquante-six"},
		{789012, "sept cent quatre-vingt-neuf mille douze"},
		{1000000, "un million"},
		{9876543, "neuf millions huit cent soixante-seize mille cinq cent quarante-trois"},
		// {1000000000, "un milliard"},
		// {1234567890, "un milliard deux cent trente-quatre million cinq cent soixante-sept mille huit cent quatre-vingt-dix"},
	}
	test(tests,t)
}

func test(values []testSuit,t *testing.T) {
	for _, test := range values {
		t.Run("", func(t *testing.T) {
			result := simpleNumber(test.input)
			if result != test.expected {
				t.Errorf("For input %d, expected %s, but got %s", test.input, test.expected, result)
			}
		})
	}
}

package utils

import (
	"../testutils"
	"log"
	"testing"
)

func TestEquivalence(t *testing.T) {

	preBuiltString := "His money is twice tainted: 'taint yours and 'taint mine."
	preBuilt := ComputeHmac256(preBuiltString)

	halfAString := "His money is twice tainted:"
	adJoinedString1 := halfAString + " 'taint yours and 'taint mine."
	adJoinedToBuild := ComputeHmac256(adJoinedString1)

	testutils.AssertTrue(t, preBuiltString == adJoinedString1, "Messages are not equivalent")
	testutils.AssertTrue(t, IsHMACEqual(preBuilt, adJoinedToBuild), "Message hashes are not equivalent")
}

func TestInEquivalence(t *testing.T) {

	preBuiltString := "His money is twice tainted: 'taint yours and 'taint mine."
	preBuilt := ComputeHmac256(preBuiltString)

	halfAString := "His money is twice tainted:"
	adJoinedString2 := halfAString + " 'taint yours and 'taint mine. It is someone else'"
	adJoinedToBuild := ComputeHmac256(adJoinedString2)

	testutils.AssertFalse(t, preBuiltString == adJoinedString2, "Messages are equivalent")
	testutils.AssertFalse(t, IsHMACEqual(preBuilt, adJoinedToBuild), "Message hashes are equivalent")

	log.Default().Println(HMACToString(preBuilt)[:7] + "... is not equal to: " + HMACToString(adJoinedToBuild)[:7] + "...")
}

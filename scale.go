package main

import (
	"math/rand"
	"time"
)

type key struct {
	Index int
	Name  string
}

type scale struct {
	ScaleName string
	ModeName  string
	Notes     []string
	Key       key
}

type scaleResponse struct {
	ScaleName string
	ModeName  string
	Notes     []string
}

func getNotes() []string {
	return []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
}

func getScales() []string {
	return []string{"Major", "Chromatic", "Blues", "Harmonic Minor", "Melodic Minor"}
}
func getScaleStructures() map[string][]int {
	return map[string][]int{
		"Major":          {2, 2, 1, 2, 2, 2, 1},
		"Chromatic":      {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		"Blues":          {3, 2, 1, 1, 3, 2},
		"Harmonic Minor": {2, 1, 2, 2, 1, 3, 1},
		"Melodic Minor":  {2, 1, 2, 2, 2, 2, 1},
	}
}

func getModeNames() map[string][]string {
	return map[string][]string{
		"Major":          {"Major", "Dorian", "Phrygian", "Lydian", "Mixolydian", "Aolian (natural minor)", "Locrian"},
		"Chromatic":      {"Chromatic"},
		"Blues":          {"Blues"},
		"Harmonic Minor": {"Harmonic Minor", "Locrian Natural 6", "Ionion #5", "Dorian #4", "Phrygian Dominant", "Lydian #2", "Super Locrian bb7"},
		"Melodic Minor":  {"Melodic Minor", "Dorian b2", "Lydian #5", "Lydian Dominant", "Mixolydian b6", "Aeolian b5", "Altered Scale (Super Locrian)"},
	}
}

func getRandomScale() string {
	scales := getScales()
	rand.Seed(time.Now().Unix())
	return scales[rand.Intn(len(scales))]
}

func getRandomModeFromScale(scale string) (int, string) {
	modeNames := getModeNames()
	modes := modeNames[scale]
	rand.Seed(time.Now().Unix())
	modeIndex := rand.Intn(len(modes))
	return modeIndex, modes[modeIndex]
}

func getNotesOfScale(scaleStructure []int, key int, modeIndex int) []string {
	notes := getNotes()
	notes = rotateStringSlice(notes, key)
	totalNotes := len(notes)
	scaleLength := len(scaleStructure)
	notesOfScale := make([]string, scaleLength+1)
	currentIndex := 0
	if modeIndex > 0 {
		for _, v := range scaleStructure[:modeIndex] {
			currentIndex += v
		}
	}
	scaleStructure = rotateIntSlice(scaleStructure, modeIndex)
	notesOfScale[0] = notes[currentIndex%totalNotes]
	for i, step := range scaleStructure {
		currentIndex += step
		currentIndex = currentIndex % totalNotes
		nextNote := notes[currentIndex]
		notesOfScale[i+1] = nextNote
	}
	return notesOfScale
}

func getRandomKey() key {
	keys := getNotes()
	rand.Seed(time.Now().Unix())
	keyIndex := rand.Intn(len(keys))
	return key{
		Index: keyIndex,
		Name:  keys[keyIndex],
	}
}

func generateRandomScale() scaleResponse {
	scale := getRandomScale()
	scaleStructures := getScaleStructures()
	scaleStructure := scaleStructures[scale]
	modeIndex, modeName := getRandomModeFromScale(scale)
	key := getRandomKey()
	notesOfScale := getNotesOfScale(scaleStructure, key.Index, modeIndex)
	return scaleResponse{
		ScaleName: key.Name + " " + scale,
		ModeName:  notesOfScale[0] + " " + modeName,
		Notes:     notesOfScale,
	}
}

func rotateIntSlice(slice []int, rotationAmount int) []int {
	return append(slice[rotationAmount:], slice[:rotationAmount]...)
}

func rotateStringSlice(slice []string, rotationAmount int) []string {
	return append(slice[rotationAmount:], slice[:rotationAmount]...)
}

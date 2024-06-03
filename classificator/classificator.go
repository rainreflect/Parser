package classificator

import (
	"fmt"
	"strings"

	"github.com/jbrukh/bayesian"
	"github.com/rainreflect/parser/parser"
)

// Function to preprocess text (tokenization, lowercasing)
func preprocess(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

// Function to train and test the Bayesian classifier
func RunBayesianClassifier(technicalDocs, literaryDocs, testDocs string) {
	// Define classes
	const (
		Technical bayesian.Class = "Technical"
		Literary  bayesian.Class = "Literary"
	)
	classifier := bayesian.NewClassifier(Technical, Literary)

	// Split documents by delimiter (e.g., newline)
	technicalDocsArray := strings.Split(technicalDocs, "\n")
	literaryDocsArray := strings.Split(literaryDocs, "\n")
	testDocsArray := strings.Split(testDocs, "\n")

	// Preprocess and train the classifier
	for _, doc := range technicalDocsArray {
		classifier.Learn(preprocess(doc), Technical)
	}
	for _, doc := range literaryDocsArray {
		classifier.Learn(preprocess(doc), Literary)
	}

	// Convert frequencies to probabilities
	classifier.ConvertTermsFreqToTfIdf()

	// Classify test documents
	for _, doc := range testDocsArray {
		scores, likely, _ := classifier.LogScores(preprocess(doc))
		fmt.Printf("Document: \"%s\"\n", doc)
		fmt.Printf("Scores: Technical: %.2f, Literary: %.2f\n", scores[0], scores[1])
		fmt.Printf("Classified as: %s\n\n", classifier.Classes[likely])
	}

	// Evaluation of the classifier (for simplicity, using the same test data)
	correct := 0
	for _, doc := range technicalDocsArray {
		_, likely, _ := classifier.LogScores(preprocess(doc))
		if classifier.Classes[likely] == Technical {
			correct++
		}
	}
	for _, doc := range literaryDocsArray {
		_, likely, _ := classifier.LogScores(preprocess(doc))
		if classifier.Classes[likely] == Literary {
			correct++
		}
	}

	total := len(technicalDocsArray) + len(literaryDocsArray)
	accuracy := float64(correct) / float64(total)
	fmt.Printf("Accuracy: %.2f%%\n", accuracy*100)
}

func Classificate(inputTexts string) {

	var techTexts string
	var notTechTexts string

	trainingTechUrls := []string{
		"https://theappsolutions.com/blog/tips/google-maps-tips/",
		"https://theappsolutions.com/blog/marketing/make-look-your-app-cool/",
		"https://theappsolutions.com/blog/development/epic-vs-cerner/",
	}

	trainingNoTechUrls := []string{
		"https://theappsolutions.com/blog/tips/google-maps-tips/",
		"https://theappsolutions.com/blog/marketing/10-ways-build-push-notification/",
		"https://theappsolutions.com/blog/marketing/make-look-your-app-cool/",
		"https://theappsolutions.com/blog/development/epic-vs-cerner/",
	}

	techItem, err := parser.ParseData(trainingTechUrls)
	if err != nil {
		panic(err)
	}

	notTechItem, err := parser.ParseData(trainingNoTechUrls)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(trainingTechUrls); i++ {
		techTexts = techTexts + " " + techItem[i].Text
		notTechTexts = notTechTexts + " " + notTechItem[i].Text
	}
	RunBayesianClassifier(techTexts, notTechTexts, inputTexts)
}

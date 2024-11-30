package basicMl

import (
	"fmt"
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwitworth/golearn/base"
)

func basicMl() {
	irisData, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	trainData, testData := base.InstancesTrainTestSplit(irisData, 0.7)

	model := ensemble.NewRandomForest(100)

	err = model.Fit(trainData)
	if err != nil {
		log.Fatal(err)
	}

	predictions := model.Predict(testData)

	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	fmt.Println("Confusion Matrix:\n", confusionMat)

	accuracy := evaluation.GetAccuracy(confusionMat)
	fmt.Printf("Accuracy: %.2f%%\n", accuracy*100)

}

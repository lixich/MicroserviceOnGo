package main

import (
	"fmt"
	"strings"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
	"gonum.org/v1/gonum/mat"
)

var tree base.Classifier

func init() {
	// Random Trees
	tree = trees.NewRandomTree(0)
	tree.Load("model.txt")
}

func CreateModel() {
	rawData, err := base.ParseCSVToInstances("AntiFraudDataSet.csv", true)
	if err != nil {
		panic(err)
	}
	// Train, test.
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	tree = trees.NewRandomTree(2)
	// Fit.
	err = tree.Fit(trainData)
	if err != nil {
		panic(err)
	}
	// Predict.
	predictions, err := tree.Predict(testData)
	if err != nil {
		panic(err)
	}
	// Performance.
	fmt.Println("RandomTree Performance")
	cf, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(cf))

	// Save.
	err = tree.Save("model.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to save model: %s", err.Error()))
	}
}

func СompletenessСheckTransaction(s string) bool {
	// TODO Check keys, format (bad implementation).
	var res bool = true
	for _, key := range TransactionKeys[:len(TransactionKeys)-1] {
		if !strings.Contains(s, key) {
			res = false
		}
	}
	return res
}

func PredictTransaction(t Transaction) Transaction {
	data := mat.NewDense(1, 7, []float64{
		t.Sum,
		float64(t.YearEnd),
		t.Hour,
		float64(t.WeekDay),
		float64(t.HaveDeviceId),
		float64(t.EuropeAsiaCountryIp),
		float64(t.Refund),
	})
	XORdata := base.InstancesFromMat64(1, 7, data)
	classAttr := base.GetAttributeByName(XORdata, "6")
	XORdata.AddClassAttribute(classAttr)
	for index, key := range TransactionKeys {
		atr := XORdata.AllAttributes()[index]
		atr.SetName(key)
	}
	predictions, err := tree.Predict(XORdata)
	if err != nil {
		panic(err)
	}
	if predictions.RowString(0) == "1.00" {
		t.Refund = 1
	} else {
		t.Refund = 0
	}
	return t
}

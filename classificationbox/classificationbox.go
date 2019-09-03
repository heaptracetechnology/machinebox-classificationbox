package classificationbox

import (
	"context"
	"encoding/json"
	result "github.com/heaptracetechnology/machinebox-classificationbox/result"
	"github.com/machinebox/sdk-go/classificationbox"
	"net/http"
	"os"
)

//ClassificationBox struct
type ClassificationBox struct {
	ID        string                      `json:"id,omitempty"`
	Name      string                      `json:"name"`
	Ngrams    int                         `json:"ngrams,omitempty"`
	Skipgrams int                         `json:"skipgrams,omitempty"`
	Classes   []string                    `json:"classes,omitempty"`
	ModelID   string                      `json:"modelId,omitempty"`
	Class     string                      `json:"class,omitempty"`
	Inputs    []classificationbox.Feature `json:"inputs,omitempty"`
	Limit     int                         `json:"limit,omitempty"`
}

//Feature struct
type Feature struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

//Message struct
type Message struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

//CreateModel MachineBox-ClassificationBox
func CreateModel(responseWriter http.ResponseWriter, request *http.Request) {

	address := os.Getenv("ADDRESS")

	decoder := json.NewDecoder(request.Body)
	var param ClassificationBox
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client := classificationbox.New(address)

	ctx, _ := context.WithCancel(context.Background())

	model := classificationbox.Model{
		ID:   param.ID,
		Name: param.Name,
		Options: &classificationbox.ModelOptions{
			Ngrams:    param.Ngrams,
			Skipgrams: param.Skipgrams,
		},
		Classes: param.Classes,
	}

	newModel, newModelErr := client.CreateModel(ctx, model)
	if newModelErr != nil {
		result.WriteErrorResponseString(responseWriter, newModelErr.Error())
		return
	}

	bytes, _ := json.Marshal(newModel)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//TeachModel MachineBox-ClassificationBox
func TeachModel(responseWriter http.ResponseWriter, request *http.Request) {

	address := os.Getenv("ADDRESS")

	decoder := json.NewDecoder(request.Body)
	var param ClassificationBox
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client := classificationbox.New(address)

	ctx, _ := context.WithCancel(context.Background())

	teach := classificationbox.Example{
		Class:  param.Class,
		Inputs: param.Inputs,
	}

	teachModelErr := client.Teach(ctx, param.ModelID, teach)
	if teachModelErr != nil {
		result.WriteErrorResponseString(responseWriter, teachModelErr.Error())
		return
	}

	message := Message{true, "success", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//GetModel MachineBox-ClassificationBox
func GetModel(responseWriter http.ResponseWriter, request *http.Request) {

	address := os.Getenv("ADDRESS")

	decoder := json.NewDecoder(request.Body)
	var param ClassificationBox
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client := classificationbox.New(address)

	ctx, _ := context.WithCancel(context.Background())

	model, newModelErr := client.GetModel(ctx, param.ModelID)
	if newModelErr != nil {
		result.WriteErrorResponseString(responseWriter, newModelErr.Error())
		return
	}

	bytes, _ := json.Marshal(model)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//MakePredictions MachineBox-ClassificationBox
func MakePredictions(responseWriter http.ResponseWriter, request *http.Request) {

	address := os.Getenv("ADDRESS")

	decoder := json.NewDecoder(request.Body)
	var param ClassificationBox
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client := classificationbox.New(address)

	ctx, _ := context.WithCancel(context.Background())

	predict := classificationbox.PredictRequest{
		Limit:  param.Limit,
		Inputs: param.Inputs,
	}

	predictResp, predictErr := client.Predict(ctx, param.ModelID, predict)
	if predictErr != nil {
		result.WriteErrorResponseString(responseWriter, predictErr.Error())
		return
	}

	bytes, _ := json.Marshal(predictResp)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//ListingModels MachineBox-ClassificationBox
func ListingModels(responseWriter http.ResponseWriter, request *http.Request) {

	address := os.Getenv("ADDRESS")

	decoder := json.NewDecoder(request.Body)
	var param ClassificationBox
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client := classificationbox.New(address)

	ctx, _ := context.WithCancel(context.Background())

	listModels, listingErr := client.ListModels(ctx)
	if listingErr != nil {
		result.WriteErrorResponseString(responseWriter, listingErr.Error())
		return
	}

	bytes, _ := json.Marshal(listModels)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//DeleteModel MachineBox-ClassificationBox
func DeleteModel(responseWriter http.ResponseWriter, request *http.Request) {

	address := os.Getenv("ADDRESS")

	decoder := json.NewDecoder(request.Body)
	var param ClassificationBox
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client := classificationbox.New(address)

	ctx, _ := context.WithCancel(context.Background())

	deleteErr := client.DeleteModel(ctx, param.ModelID)
	if deleteErr != nil {
		result.WriteErrorResponseString(responseWriter, deleteErr.Error())
		return
	}

	message := Message{true, "Model \"" + param.ModelID + "\" deleted successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

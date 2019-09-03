package classificationbox

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Create model with invalid param", func() {

	classificationBox := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createModel", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateModel)
	handler.ServeHTTP(recorder, request)

	Describe("Create Model", func() {
		Context("create model", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create model with valid param", func() {

	classesList := []string{"class1", "class2", "class3"}
	classificationBox := ClassificationBox{ID: "Test Model ID", Name: "Test ClassifationModel", Ngrams: 1, Skipgrams: 1, Classes: classesList}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createModel", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateModel)
	handler.ServeHTTP(recorder, request)

	Describe("Create Model", func() {
		Context("create model", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Teach model with invalid param", func() {

	classificationBox := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/teachModel", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(TeachModel)
	handler.ServeHTTP(recorder, request)

	Describe("Teach Model", func() {
		Context("teach model", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Teach model with valid param", func() {

	classificationBox := ClassificationBox{ModelID: "omg model", Class: "class"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/teachModel", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(TeachModel)
	handler.ServeHTTP(recorder, request)

	Describe("Teach Model", func() {
		Context("teach model", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get model with invalid param", func() {

	classificationBox := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getModel", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetModel)
	handler.ServeHTTP(recorder, request)

	Describe("Get Model", func() {
		Context("get model", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get model with valid param", func() {

	classificationBox := ClassificationBox{ModelID: "omg model"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getModel", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetModel)
	handler.ServeHTTP(recorder, request)

	Describe("Get Model", func() {
		Context("get model", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Make predictions with invalid param", func() {

	classificationBox := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/makePredictions", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MakePredictions)
	handler.ServeHTTP(recorder, request)

	Describe("Make Predictions", func() {
		Context("Make predictions", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Make predictions with valid param", func() {

	classificationBox := ClassificationBox{ModelID: "omg model", Limit: 10}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/makePredictions", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MakePredictions)
	handler.ServeHTTP(recorder, request)

	Describe("Make Predictions", func() {
		Context("Make predictions", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List models with invalid param", func() {

	classificationBox := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/listModels", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListingModels)
	handler.ServeHTTP(recorder, request)

	Describe("List models", func() {
		Context("List models", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List models with valid param", func() {

	classificationBox := ClassificationBox{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/listModels", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListingModels)
	handler.ServeHTTP(recorder, request)

	Describe("List Models", func() {
		Context("List models", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete Model with invalid param", func() {

	classificationBox := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/deleteModel", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteModel)
	handler.ServeHTTP(recorder, request)

	Describe("Delete Model", func() {
		Context("Delete model", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete Model valid param", func() {

	classificationBox := ClassificationBox{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(classificationBox)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/deleteModel", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteModel)
	handler.ServeHTTP(recorder, request)

	Describe("Delete Model", func() {
		Context("Delete model", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

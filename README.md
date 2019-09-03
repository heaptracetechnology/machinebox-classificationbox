# _Machinebox-Classificationbox_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/machinebox-classificationbox.svg?branch=master)](https://travis-ci.com/omg-services/machinebox-classificationbox)
[![codecov](https://codecov.io/gh/omg-services/machinebox-classificationbox/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/machinebox-classificationbox)


An OMG service for machinebox-classificationbox, it lets you use machine learning to automatically classify various types of data, such as text, images, structured and unstructured data.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Create models
```coffee
machinebox-classificationbox createModel id:'modelID' name:'modelName' ngram:1 skipgrams:1 classes:'["class1", "class2", "class3"]'
#result {"id": "ModelID","name": "ModelName","options": {"ModelOptions"},"classes": ["ClassesList"]}
```
##### Teach model
```coffee
machinebox-classificationbox teachModel modelId:'modelID' class:'className' inputs:'[{"key": "user_age", "type": "number", "value": "25"},{"key": "user_interests", "type": "list", "value": "music,cooking,ml"},{"key": "user_location", "type": "keyword", "value": "London"}]'
#result {"success": true,"message": "success","statusCode": 200}
```
##### Get Model
```coffee
machinebox-classificationbox getModel modelId:'modelID'
#result {"id": "ModelID","name": "ModelName","options": {"ModelOptions"},"classes": ["ClassesList"]}
```
##### Make predictions
```coffee
machinebox-classificationbox makePredictions modelId:'modelID' limit:10 inputs:'[{"key": "user_age", "type": "number", "value": "25"},{"key": "user_interests", "type": "list", "value": "music,cooking,ml"},{"key": "user_location", "type": "keyword", "value": "London"}]'
#result {"classes": [{"id": "class1","score": 0.42109},{"id": "class2","score": 0.293062}]}
```
##### Listing models
```coffee
machinebox-classificationbox listModels
#result [{"id":"omg model1","name":"classifationModel1"},{"id":"omg model2","name":"classifationModel2"}]
```
##### Deleting models
```coffee
machinebox-classificationbox deleteModel modelId:'modelID'
#result {"success": true, "message": "Model "omg model1" deleted successfully","statusCode": 200}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create models
```shell
omg run createModel -a id=<MODEL_ID> -a name=<MODEL_NAME> -a ngram=<N-GRAMS> -a skipgrams=<SKIPGRAMS> -a classes=<LIST_OF CLASSES> -e ADDRESS=<SERVER_IP_ADDRESS_WITH_PORT>
```
##### Teach model
```shell
omg run teachModel -a modelId=<MODEL_ID> -a class=<CLASS_NAME> -a inputs=<LIST_OF_INPUT_FEATURES> -e ADDRESS=<SERVER_IP_ADDRESS_WITH_PORT>
```
##### Get Model
```shell
omg run getModel -a modelId=<MODEL_ID> -e ADDRESS=<SERVER_IP_ADDRESS_WITH_PORT>
```
##### Make predictions
```shell
omg run makePredictions -a modelId=<MODEL_ID> -a limit=<LIMIT> -a inputs=<LIST_OF_INPUT_FEATURES> -e ADDRESS=<SERVER_IP_ADDRESS_WITH_PORT>
```
##### Listing models
```shell
omg run listModels -e ADDRESS=<SERVER_IP_ADDRESS_WITH_PORT>
```
##### Deleting models
```shell
omg run deleteModel -a modelId=<MODEL_ID> -e ADDRESS=<SERVER_IP_ADDRESS_WITH_PORT>
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/machinebox-classificationbox/blob/master/LICENSE).

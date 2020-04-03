
# n-dilation-computer

_Simple microservice to compute a unitary N dilation of a given contraction matrix written by a golang beginner_

This project is licensed under the terms of the MIT license.

If you stumbled across this and are interested, found a bug or have another idea how to contribute, feel free to open an issue or a pull request.

The recipe to calculate the dilation follows the theory mentioned by Béla Szőkefalvi-Nagy in "Analyse harmonique des opérateurs de l'espace de Hilbert" (1967). The matrix square root needed for the dilation is calculated using the Exponential Method for Matrices.

## Usage

Run any common `go` tasks such as `go build`, `go test ./...` or `go run main.go`. The latter will start the server listening on port `:8080`.
This project was built on `go version go1.13.9`.

## Endpoints

### ```POST /dilation```

Post a request to `/dilation` with a json body containing the following fields:

```json
	{
		"value": [0.5,0.25,0.5,0],
		"degree": 2
	}
```

Field explanation:
- value: The contraction in row-major order
- degree: The "N" in the unitary N dilation, the degree N as the upper boundary for which exponents the dilation condition should hold true


The response will be a json containing the dilation:

```json
	{
		"value": [
			0.5,0.25,0,0,0.7031267015263438,-0.07491894857957959,
			0.5,0,0,0,-0.0749189481139183,0.9653430208563805,
			0.815505113452673,-0.1498379036784172,0,0,-0.5,-0.5,
			-0.1498379036784172,0.8529645912349224,0,0,-0.25,-0,
			0,0,1,0,0,0,
			0,0,0,1,0,0
		]
	}
```

If there are fields missing in the request body or the provided `value` does not meet the precondition to calculate the dilation, an error message will be sent. Example:

```json
	{
		"validationError": {
			"value": [
				"value must represent a real contraction"
			]
		}
	}
```

Example with cURL:
```bash
	curl -X POST -H 'Content-Type: application/json' -d "{\"degree\": 2, \"value\": [0.5,0.25,0.5,0]}" http://localhost:8080/dilation
```

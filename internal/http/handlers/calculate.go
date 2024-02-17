package handlers

import (
	"encoding/json"
	"fmt"
	ioreader "io"
	"net/http"

	"github.com/codeaspiras/gotravel-http/internal/altcontext"
	"github.com/codeaspiras/gotravel-http/internal/http/render"
	"github.com/codeaspiras/gotravel/calculator"
)

type CalculateInput struct {
	CostPerLiter     float64 `json:"cost_per_liter" validate:"required,gt=0"`
	DistanceToGo     float64 `json:"distance_to_go" validate:"required,gt=0"`
	DistanceToReturn float64 `json:"distance_to_return" validate:"required,gt=0"`
	DistancePerLiter float64 `json:"distance_per_liter" validate:"required,gt=0"`
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	io, err := altcontext.GetIO(r.Context())
	if err != nil {
		render.Error(nil, w, err, http.StatusInternalServerError)
		return
	}

	validator, err := altcontext.GetValidator(r.Context())
	if err != nil {
		render.Error(io, w, err, http.StatusInternalServerError)
		return
	}

	bodyBytes, err := ioreader.ReadAll(r.Body)
	if err != nil {
		render.Error(io, w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var input CalculateInput
	err = json.Unmarshal(bodyBytes, &input)
	if err != nil {
		render.Error(io, w, err, http.StatusBadRequest)
		return
	}

	errors := validator.Validate(input)
	if len(errors) > 0 {
		render.ValidationError(io, w, errors)
	}

	totalDistance := input.DistanceToGo + input.DistanceToReturn
	result := calculator.CostToCover(
		totalDistance,
		input.DistancePerLiter,
		input.CostPerLiter,
	)

	render.JSON(io, w, http.StatusOK, map[string]interface{}{
		"totalDistance": fmt.Sprintf("%.2f km", totalDistance),
		"cost":          fmt.Sprintf("$%.2f", result),
	})
}

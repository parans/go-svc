package handlers

import (
	"net/http"
)

// Charts returns all charts
// swagger:operation GET /charts charts
// ---
// summary: Find all charts
// description: Returns all charts and their versions with its metadata
// operationId: charts
// produces:
//   - application/json
// responses:
//   '200':
//     description: successful operation
//     schema:
//       type: array
//       items:
//         $ref: '#/definitions/Metadata'
//   '401':
//     description: Authenticatin failed
//   '500':
//     description: Internal server error
func (s *Server) Charts(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// UploadChart uploads a chart package
// swagger:operation POST /charts charts
// ---
// summary: Upload a Kubernetes Chart package
// operationId: uploadChart
// consumes:
//   - multipart/form-data
// produces:
//   - application/json
// parameters:
//   - name: file
//     in: formData
//     description: kubernetes chart archive file. ex. chartname-1.0.0.tgz
//     required: true
//     type: file
// responses:
//   '200':
//     description: successful operation
//     schema:
//       $ref: '#/definitions/Metadata'
//   '400':
//     description: Invalid input
//   '401':
//     description: Authenticatin failed
//   '500':
//     description: Internal server error
func (s *Server) UploadChart(w http.ResponseWriter, r *http.Request) error {
	return nil
}

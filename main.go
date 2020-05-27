// Deploy-Service API
// ---
//     Schemes: https
//     BasePath: /api/deploy/v1
//     Version: 1.0.0
//     Contact: shalin.patel@teradata.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     securityDefinitions:
//       bearer:
//         type: apiKey
//         name: Authorization
//         in: header
//     security:
//       - bearer:
//
// swagger:meta
package main

import (
	"log"
	"github.td.teradata.com/appcenter/backup-restore/cmd"
)

// go:generate swagger generate spec -o ./swaggerui/swagger-spec.json --scan-models --exclude-deps
func main() {
	log.Fatal(cmd.Execute())
}

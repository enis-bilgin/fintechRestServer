/*
 * Backend Test API
 *
 * Service for receiving trades from the outside world
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Internal representation of a trade including internal id
type InternalTrade struct {
	// Unique ID for this trade
	Id string `json:"id"`

	Trade *Trade `json:"trade"`
}

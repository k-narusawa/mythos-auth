package change_email

import (
	"context"
	"fmt"
	"os"

	ory "github.com/ory/kratos-client-go"
)

type ChangeEmailService struct{}

func NewChangeEmailService() *ChangeEmailService {
	return new(ChangeEmailService)
}

func (ces *ChangeEmailService) Invoke(in *ChangeEmailInputData) (out *ChangeEmailOutputData, err error) {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4434", // Kratos Admin API
		},
	}

	apiClient := ory.NewAPIClient(configuration)

	uib := *ory.NewUpdateIdentityBody(
		"preset://email",
		"active",
		map[string]interface{}{
			"email": in.Email,
		},
	)

	ui, resp, err := apiClient.IdentityApi.
		UpdateIdentity(context.Background(), in.Id).
		UpdateIdentityBody(uib).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrontendApi.UpdateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
	}

	return &ChangeEmailOutputData{
		Id:        ui.Id,
		Email:     in.Email,
		CreatedAt: *ui.CreatedAt,
		UpdatedAt: *ui.UpdatedAt,
	}, nil
}

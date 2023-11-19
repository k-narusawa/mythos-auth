package register_user

import (
	"context"
	"fmt"
	"os"

	ory "github.com/ory/kratos-client-go"
)

type RegisterUserService struct{}

func NewRegisterUserService() *RegisterUserService {
	return new(RegisterUserService)
}

func (r *RegisterUserService) Invoke(in *RegisterUserInputData) (out *RegisterUserOutputData, err error) {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4434", // Kratos Admin API
		},
	}

	apiClient := ory.NewAPIClient(configuration)

	cib := *ory.NewCreateIdentityBody(
		"preset://email",
		map[string]interface{}{
			"email": in.Email,
		},
	)

	cib.Credentials = ory.NewIdentityWithCredentials()
	cib.Credentials.SetPassword(ory.IdentityWithCredentialsPassword{
		Config: &ory.IdentityWithCredentialsPasswordConfig{
			Password: ory.PtrString(in.Password),
		},
	})

	ci, resp, err := apiClient.IdentityApi.
		CreateIdentity(context.Background()).
		CreateIdentityBody(cib).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrontendApi.CreateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
	}

	return &RegisterUserOutputData{
		Id:        ci.Id,
		Email:     in.Email,
		CreatedAt: *ci.CreatedAt,
		UpdatedAt: *ci.UpdatedAt,
	}, nil
}

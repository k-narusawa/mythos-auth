package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/labstack/echo/v4"
)

type MailController struct{}

func NewMailController() MailController {
	return MailController{}
}

func (mc *MailController) Send(c echo.Context) error {
	ve := new(VerificationEmailRequest)

	if err := c.Bind(ve); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	fmt.Fprintf(os.Stdout, "TemplateType: %v\n", ve.TemplateType)

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{URL: "http://localhost:8005"}, nil
		})))
	cfg.Region = "ap-northeast-1"

	if err != nil {
		// error handling
	}

	var (
		fromEmailAddress = "mythos@example.com"
		toEmailAddress   = ve.Recipient
		subject          = ve.Subject
		body             = ve.Body
	)

	client := sesv2.NewFromConfig(cfg)

	input := &sesv2.SendEmailInput{
		FromEmailAddress: &fromEmailAddress,
		Destination: &types.Destination{
			ToAddresses: []string{toEmailAddress},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &body,
					},
				},
				Subject: &types.Content{
					Data: &subject,
				},
			},
		},
	}

	res, err := client.SendEmail(ctx, input)
	if err != nil {
		// error handling
	}
	fmt.Println(res.MessageId)
	fmt.Println("success!")

	return c.NoContent(http.StatusNoContent)
}

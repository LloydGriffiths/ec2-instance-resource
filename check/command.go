package check

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Command represents a check command.
type Command struct {
	service *ec2.EC2
}

// Source represents the check source.
type Source struct {
	AccessKeyID     string       `json:"access_key_id"`
	Filters         []ec2.Filter `json:"filters"`
	Region          string       `json:"region"`
	SecretAccessKey string       `json:"secret_access_key"`
	SessionToken    string       `json:"session_token"`
}

// Request represnts the check request.
type Request struct {
	Source  *Source  `json:"source"`
	Version *Version `json:"version"`
}

// Version represnts the resource version.
type Version struct {
	Date      string `json:"date"`
	Instances string `json:"instances"`
}

// New creates and returns a new check command.
func New(source *Source) *Command {
	credentials := aws.NewStaticCredentialsProvider(
		source.AccessKeyID,
		source.SecretAccessKey,
		source.SessionToken,
	)

	config := aws.Config{
		Credentials:      credentials,
		EndpointResolver: endpoints.NewDefaultResolver(),
		Handlers:         defaults.Handlers(),
		HTTPClient:       defaults.HTTPClient(),
		Region:           source.Region,
	}

	return &Command{ec2.New(config)}
}

// Run checks for new version.
func (c *Command) Run(request *Request) ([]*Version, error) {
	r, err := c.instances(request)
	if err != nil {
		return nil, err
	}
	if request.Version != nil && r == request.Version.Instances {
		return []*Version{}, nil
	}
	if len(r) == 0 {
		return []*Version{}, nil
	}
	return []*Version{{Date: time.Now().Format(time.UnixDate), Instances: r}}, nil
}

func (c *Command) instances(request *Request) (string, error) {
	reqd := c.service.DescribeInstancesRequest(&ec2.DescribeInstancesInput{Filters: request.Source.Filters})
	page := reqd.Paginate()
	iids := []string{}

	for page.Next() {
		iids = append(iids, extract(page.CurrentPage().Reservations)...)
	}

	if page.Err() != nil {
		return strings.Join(iids, "\n"), page.Err()
	}

	return strings.Join(iids, "\n"), nil
}

func extract(reservations []ec2.RunInstancesOutput) (instances []string) {
	for _, reservation := range reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, *instance.InstanceId)
		}
	}
	return
}

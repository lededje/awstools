package common

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type SessionFlags struct {
	RoleArn         *string
	RoleExternalID  *string
	RoleSessionName *string
	Region          *string
	MFASerialNumber *string
	MFATokenCode    *string
	Duration        *time.Duration
}

func KingpinSessionFlags() *SessionFlags {
	return &SessionFlags{
		RoleArn:         kingpin.Flag("assume-role-arn", "Role to assume").String(),
		RoleExternalID:  kingpin.Flag("assume-role-external-id", "External ID of the role to assume").String(),
		RoleSessionName: kingpin.Flag("assume-role-session-name", "Role session name").String(),
		Region:          kingpin.Flag("region", "AWS Region").String(),
		MFASerialNumber: kingpin.Flag("mfa-serial-number", "MFA Serial Number").String(),
		MFATokenCode:    kingpin.Flag("mfa-token-code", "MFA Token Code").String(),
		Duration:        kingpin.Flag("session-duration", "Session Duration").Default("1h").Duration(),
	}
}

func NewConfig(region string) *aws.Config {
	if region == "" {
		region = os.Getenv("AWS_REGION")
	}
	if region == "" {
		region = "eu-west-1"
	}

	for _, key := range []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"} {
		value := os.Getenv(key)
		if value != "" {
			if strings.TrimSpace(value) != value {
				Fatalln(fmt.Sprintf("%s has trailing spaces, please check your config", key))
			}
		}
	}

	return &aws.Config{Region: aws.String(region)}
}

func NewSession(region string) *session.Session {
	awsConfig := NewConfig(region)
	return session.New(awsConfig)
}

type SessionTokenProvider struct {
	SessionFlags *SessionFlags
	Session      *session.Session
}

func (p *SessionTokenProvider) Retrieve() (credentials.Value, error) {
	result := credentials.Value{}

	var tokenCode *string
	if *p.SessionFlags.MFATokenCode == "" {
		stdinCode, err := stscreds.StdinTokenProvider()
		if err != nil {
			return result, nil
		}
		tokenCode = aws.String(stdinCode)
	} else {
		tokenCode = p.SessionFlags.MFATokenCode
	}

	input := &sts.GetSessionTokenInput{
		SerialNumber: p.SessionFlags.MFASerialNumber,
		TokenCode:    tokenCode,
	}
	conf := NewConfig(*p.SessionFlags.Region)
	stsClient := sts.New(p.Session, conf)
	output, err := stsClient.GetSessionToken(input)
	if err != nil {
		return result, err
	}

	if output.Credentials == nil {
		return result, errors.New("Could not get credentials")
	}

	return credentials.Value{
		AccessKeyID:     *output.Credentials.AccessKeyId,
		SecretAccessKey: *output.Credentials.SecretAccessKey,
		SessionToken:    *output.Credentials.SessionToken,
	}, nil
}

func (p *SessionTokenProvider) IsExpired() bool {
	return false
}

func OpenSession(sessionFlags *SessionFlags) (*session.Session, *aws.Config) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))
	return sess, AssumeRoleConfig(sessionFlags, sess)
}

func AssumeRoleConfig(sessionFlags *SessionFlags, sess *session.Session) *aws.Config {
	conf := NewConfig(*sessionFlags.Region)
	if sessionFlags.RoleArn != nil && *sessionFlags.RoleArn != "" {
		var creds *credentials.Credentials
		creds = stscreds.NewCredentials(sess, *sessionFlags.RoleArn, func(p *stscreds.AssumeRoleProvider) {
			if *sessionFlags.RoleExternalID != "" {
				p.ExternalID = sessionFlags.RoleExternalID
			}

			if *sessionFlags.RoleSessionName != "" {
				p.RoleSessionName = *sessionFlags.RoleSessionName
			}

			if sessionFlags.Duration != nil {
				p.Duration = *sessionFlags.Duration
			}

			if *sessionFlags.MFASerialNumber != "" {
				p.SerialNumber = sessionFlags.MFASerialNumber
				if len(*sessionFlags.MFATokenCode) == 0 {
					p.TokenProvider = stscreds.StdinTokenProvider
				} else {
					p.TokenCode = sessionFlags.MFATokenCode
				}
			}
		})
		conf.Credentials = creds
	} else if sessionFlags.MFASerialNumber != nil && *sessionFlags.MFASerialNumber != "" {
		conf.Credentials = credentials.NewCredentials(&SessionTokenProvider{
			SessionFlags: sessionFlags,
			Session:      sess,
		})
	}
	return conf
}

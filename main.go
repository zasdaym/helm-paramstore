package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

var version string = "development"

func run() error {
	var (
		showVersion = flag.Bool("version", false, "Show version")
		valuesFile  = flag.String("f", "values.yaml", "Helm values file location")
	)
	flag.Parse()

	if *showVersion {
		_, err := fmt.Printf("version\n")
		return err
	}

	paramstoreFunc, err := getFromAWSParamStore()
	if err != nil {
		return err
	}

	funcMap := template.FuncMap{
		"paramstore": paramstoreFunc,
	}
	content, err := os.ReadFile(*valuesFile)
	if err != nil {
		return err
	}
	tmpl, err := template.New("").Funcs(funcMap).Parse(string(content))
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, nil)
}

// getFromAWSParamStore returns a function to get a value from AWS SSM Parameter Store.
// Closure used to prevent creating new AWS session object repetitively.
func getFromAWSParamStore() (func(string) (string, error), error) {
	sess, err := session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	)
	if err != nil {
		return nil, err
	}

	fn := func(name string) (string, error) {
		svc := ssm.New(sess)
		withDecryption := true
		out, err := svc.GetParameter(&ssm.GetParameterInput{
			Name:           &name,
			WithDecryption: &withDecryption,
		})
		if err != nil {
			return "", err
		}
		return *out.Parameter.Value, nil
	}
	return fn, nil
}

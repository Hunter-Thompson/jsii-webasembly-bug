package main

import (
	"syscall/js"
	"test/dom"
	"test/imports/k8s"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps, appName string) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	labels := map[string]*string{"app": jsii.String(appName)}

	podSpec := &k8s.PodSpec{
		Containers: &[]*k8s.Container{
			{
				Name:  jsii.String(appName),
				Image: jsii.String("nginx:latest"),
				Ports: &[]*k8s.ContainerPort{{ContainerPort: jsii.Number(80)}},
			},
		},
	}

	k8s.NewKubeDeployment(chart, jsii.String("nginx-deployment"), &k8s.KubeDeploymentProps{
		Spec: &k8s.DeploymentSpec{
			Selector: &k8s.LabelSelector{
				MatchLabels: &labels,
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Labels: &labels,
				},
				Spec: podSpec,
			},
		},
	})

	return chart
}

func synth(_ js.Value, _ []js.Value) interface{} {
	appName := dom.GetString("appName", "value")

	app := cdk8s.NewApp(nil)
	NewMyChart(app, "hello-cdk8s-go", nil, appName)
	_ = app.SynthYaml()

	return nil
}

func main() {
	dom.AddEventListener("appName", "input", synth)

	ch := make(chan struct{})
	<-ch
}

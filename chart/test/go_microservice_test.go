package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	appsv1 "k8s.io/api/apps/v1"
)

func TestDeploymentBlueTemplateRendersContainerImage(t *testing.T) {
	helmChartPath := "../"

	options := &helm.Options{
		SetValues: map[string]string{"blue.image": "photosojourn/go-microservice:latest"},
	}

	// Run RenderTemplate to render the template
	// and capture the output.
	output := helm.RenderTemplate(
		t, options, helmChartPath, "go-microservice",
		[]string{"templates/deployment-blue.yaml"})

	//fmt.Println(output)
	// Now we use kubernetes/client-go library to render the
	// template output into the Pod struct. This will
	// ensure the Pod resource is rendered correctly.
	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)
	// Finally, we verify the pod spec is set to the expected
	// container image value
	expectedContainerImage := "photosojourn/go-microservice:latest"
	podContainers := deployment.Spec.Template.Spec.Containers

	if podContainers[0].Image != expectedContainerImage {
		t.Fatalf(
			"Rendered container image (%s) is not expected (%s)",
			podContainers[0].Image,
			expectedContainerImage,
		)
	}
}

func TestDeploymentGreenTemplateRendersContainerImage(t *testing.T) {
	helmChartPath := "../"

	options := &helm.Options{
		SetValues: map[string]string{"blue.image": "photosojourn/go-microservice:latest"},
	}

	// Run RenderTemplate to render the template
	// and capture the output.
	output := helm.RenderTemplate(
		t, options, helmChartPath, "go-microservice",
		[]string{"templates/deployment-green.yaml"})

	//fmt.Println(output)
	// Now we use kubernetes/client-go library to render the
	// template output into the Pod struct. This will
	// ensure the Pod resource is rendered correctly.
	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)
	// Finally, we verify the pod spec is set to the expected
	// container image value
	expectedContainerImage := "photosojourn/go-microservice:latest"
	podContainers := deployment.Spec.Template.Spec.Containers

	if podContainers[0].Image != expectedContainerImage {
		t.Fatalf(
			"Rendered container image (%s) is not expected (%s)",
			podContainers[0].Image,
			expectedContainerImage,
		)
	}
}

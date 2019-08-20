// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var setResourcesOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Select all resources, including uninitialized ones, in the namespace of the specified resource types"},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "-c", Description: "The names of containers in the selected pod templates to change, all containers are selected by default - may use wildcards"},
	prompt.Suggest{Text: "--containers", Description: "The names of containers in the selected pod templates to change, all containers are selected by default - may use wildcards"},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--include-uninitialized", Description: "If true, the kubectl command applies to uninitialized objects. If explicitly set to false, this flag overrides other flags that make the kubectl commands apply to uninitialized objects, e.g., \"--all\". Objects with empty metadata.initializers are regarded as initialized."},
	prompt.Suggest{Text: "--limits", Description: "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges."},
	prompt.Suggest{Text: "--local", Description: "If true, set resources will NOT contact api-server but run locally."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|name|go-template-file|templatefile|template|go-template|jsonpath|jsonpath-file."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|name|go-template-file|templatefile|template|go-template|jsonpath|jsonpath-file."},
	prompt.Suggest{Text: "--record", Description: "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--requests", Description: "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, not including uninitialized ones,supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, not including uninitialized ones,supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}

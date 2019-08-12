package kube

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

func optionCompleter(args []string, long bool) []prompt.Suggest {
	l := len(args)
	if l <= 1 {
		if long {
			return prompt.FilterHasPrefix(optionHelp, "--", false)
		}
		return optionHelp
	}

	var suggests []prompt.Suggest
	commandArgs, _ := excludeOptions(args)
	switch commandArgs[0] {
	case "annotate":
		suggests = annotateOptions
	case "api-resources":
		suggests = apiResourcesOptions
	case "apply":
		suggests = applyOptions
	case "attach":
		suggests = attachOptions
	case "auth":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "can-i":
				suggests = authCanIOptions
			case "reconcile":
				suggests = authReconcileOptions
			}
		}
	case "autoscale":
		suggests = autoscaleOptions
	case "certificate":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "approve":
				suggests = certificateApproveOptions
			case "deny":
				suggests = certificateDenyOptions
			}
		}
	case "cluster-info":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "dump":
				suggests = clusterInfoDumpOptions
			}
		}
	case "config":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "get-contexts":
				suggests = configGetContextsOptions
			case "view":
				suggests = configViewOptions
			case "set-cluster":
				suggests = configSetClusterOptions
			case "set-credentials":
				suggests = configSetCredentialsOptions
			case "set":
				suggests = configSetOptions
			}
		}
	case "convert":
		suggests = convertOptions
	case "cordon":
		suggests = cordonOptions
	case "cp":
		suggests = cpOptions
	case "create":
		suggests = createOptions
	case "delete":
		suggests = deleteOptions
	case "describe":
		suggests = describeOptions
	case "drain":
		suggests = drainOptions
	case "edit":
		suggests = editOptions
	case "exec":
		suggests = execOptions
	case "explain":
		suggests = explainOptions
	case "expose":
		suggests = exposeOptions
	case "get":
		suggests = getOptions
	case "label":
		suggests = labelOptions
	case "logs":
		suggests = logsOptions
	case "patch":
		suggests = patchOptions
	case "port-forward":
		suggests = portForwardOptions
	case "proxy":
		suggests = proxyOptions
	case "replace":
		suggests = replaceOptions
	case "rolling-update":
		suggests = rollingUpdateOptions
	case "rollout":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "history":
				suggests = rolloutHistoryOptions
			case "pause":
				suggests = rolloutPauseOptions
			case "resume":
				suggests = rolloutResumeOptions
			case "status":
				suggests = rolloutStatusOptions
			case "undo":
				suggests = rolloutUndoOptions
			}
		}
	case "run", "run-container":
		suggests = runOptions
	case "scale", "resize":
		suggests = scaleOptions
	case "set":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "env":
				suggests = setEnvOptions
			case "image":
				suggests = setImageOptions
			case "resources":
				suggests = setResourcesOptions
			case "selector":
				suggests = setSelectorOptions
			case "serviceaccount":
				suggests = setServiceaccountOptions
			case "subject":
				suggests = setSubjectOptions
			}
		}
	case "top":
		if len(commandArgs) >= 2 {
			switch commandArgs[1] {
			case "no", "node", "nodes":
				suggests = topNodeOptions
			case "po", "pod", "pods":
				suggests = topPodOptions
			}
		}
	case "uncordon":
		suggests = uncordonOptions
	case "version":
		suggests = versionOptions
	default:
		suggests = optionHelp
	}

	suggests = append(suggests, globalOptions...)
	if long {
		return prompt.FilterContains(
			prompt.FilterHasPrefix(suggests, "--", false),
			strings.TrimLeft(args[l-1], "--"),
			true,
		)
	}
	return prompt.FilterContains(suggests, strings.TrimLeft(args[l-1], "-"), true)
}

var optionHelp = []prompt.Suggest{
	{Text: "-h"},
	{Text: "--help"},
}

var globalOptions = []prompt.Suggest{
	{Text: "--namespace", Description: "temporarily set the namespace for a request"},
	{Text: "-n", Description: "temporarily set the namespace for a request"},
	{Text: "--server", Description: "specify the address and port of the Kubernetes API server"},
	{Text: "-s", Description: "specify the address and port of the Kubernetes API server"},
	{Text: "--user", Description: "take the user if this flag exists."},
	{Text: "--cluster", Description: "take the cluster if this flag exists."},
}

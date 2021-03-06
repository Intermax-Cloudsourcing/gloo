package kubernetes.admission

operations = {"CREATE", "UPDATE"}

deny[msg] {
	input.request.kind.kind == "VirtualService"
	operations[input.request.operation]
	input.request.object.spec.virtualHost.routes[_].options.prefixRewrite
	msg := "prefix re-write not allowed"
}

package v1

/*
type CRDRequest struct {
	types.CRDSpec `json:",inline"`
}*/
/*
func ValidateCRDRequest(request *CRDRequest) (ResponseCode, string) {
	code, msg := ValidateNameByCRDRequest(request)
	if code == ErrClientRequestValidateCode {
		return code, msg
	}
	if request.Group == "" {
		return ErrClientRequestValidateCode, "group invalid"
	}
	if request.Resource == "" {
		return ErrClientRequestValidateCode, "resource invalid"
	}
	if request.Kind == "" {
		return ErrClientRequestValidateCode, "kind invalid"
	}
	if len(request.Versions) == 0 {
		return ErrClientRequestValidateCode, "versions is empty"
	}
	if request.Operator == nil || request.Operator.Name == "" || request.Operator.Namespace == "" {
		return ErrClientRequestValidateCode, "operator invalid"
	}
	return ValidateRequestOKCode, ""
}
*/
/*
func ValidateNameByCRDRequest(request *CRDRequest) (ResponseCode, string) {
	if request.Name == "" {
		return ErrClientRequestValidateCode, "name invalid"
	}
	return ValidateRequestOKCode, ""
}
*/
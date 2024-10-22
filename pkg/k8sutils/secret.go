package k8sutils

import (
	v1 "k8s.io/api/core/v1"
)

const (
	// ApicastSecretLabel is the label that secrets need to have in order to reconcile changes
	ApicastSecretLabel = "apicast.apps.3scale.net/watched-by=apicast"
)

func SecretStringDataFromData(secret *v1.Secret) map[string]string {
	stringData := map[string]string{}

	for k, v := range secret.Data {
		stringData[k] = string(v)
	}
	return stringData
}

func IsSecretWatchedByApicast(secret *v1.Secret) bool {
	if secret == nil {
		return false
	}

	existingLabels := secret.Labels

	if existingLabels != nil {
		if _, ok := existingLabels["apicast.apps.3scale.net/watched-by"]; ok {
			return true
		}
	}

	return false
}

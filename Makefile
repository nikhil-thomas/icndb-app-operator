gen-csv:
	operator-sdk generate csv \
	--csv-channel=alpha \
	--csv-version=0.1.0  \
	--operator-name=icndb-app-operator \
	--update-crds

clean:
	oc delete --ignore-not-found deploy/crds
	oc delete --ignore-not-found -f deploy/role.yaml
	oc delete --ignore-not-found -f deploy/role_binding.yaml
	oc delete --ignore-not-found -f deploy/service_account.yaml

setup: clean
	oc apply -f deploy/crds/icndbfun_v1alpha1_funapp_crd.yaml
	oc apply -f deploy/role.yaml
	oc apply -f deploy/role_binding.yaml
	oc apply -f deploy/service_account.yaml

run-local: setup
	operator-sdk run --local

.PHONY: opo-push-quay-app
push-bundle-to-quay:
ifndef VERSION
	@echo VERSION not set
	@exit 1
endif
ifndef MY_QUAY_NAMESPACE
	@echo QUAY_NAMESPACE not set
	@exit 1
endif
ifndef TOKEN
	@echo TOKEN not set
	@exit 1
endif
	@operator-courier --verbose push  \
		./deploy/olm-catalog/icndb-app-operator \
		${MY_QUAY_NAMESPACE} \
		icndb-app-operator \
		${VERSION}  \
		"${TOKEN}"

PHONY: olm-clean
opo-olm-clean:
	oc delete operatorsource -n openshift-marketplace ${MY_QUAY_NAMESPACE}-operators --ignore-not-found

export define operatorsource
apiVersion: operators.coreos.com/v1
kind: OperatorSource
metadata:
  name: ${MY_QUAY_NAMESPACE}-operators
  namespace: openshift-marketplace
spec:
  type: appregistry
  endpoint: https://quay.io/cnr
  registryNamespace: ${MY_QUAY_NAMESPACE}
  displayName: "${MY_QUAY_NAMESPACE} Operators"
  publisher: "${MY_QUAY_NAMESPACE}"
endef

.PHONY: operator-source
operator-source: opo-olm-clean
	@echo ::::: operator soruce manifest :::::
	@echo "$$operatorsource"
	@echo ::::::::::::::::::::::::::::::
	@echo "$$operatorsource" | oc apply -f -

export define subscription
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: ${MY_QUAY_NAMESPACE}-icndb-app-subscription
  namespace: openshift-operators
spec:
  channel: ${CHANNEL}
  name: icndb-app-operator
  source: ${MY_QUAY_NAMESPACE}-operators
  sourceNamespace: openshift-marketplace
endef

.PHONY: subscription
subscription:
	@echo ::::: subscription soruce manifest :::::
	@echo "$$subscription"
	@echo ::::::::::::::::::::::::::::::
	@echo "$$subscription" | oc apply -f -
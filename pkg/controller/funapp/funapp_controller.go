package funapp

import (
	"context"
	"reflect"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"strings"

	icndbfunv1alpha1 "github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_funapp")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new FunApp Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileFunApp{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("funapp-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource FunApp
	err = c.Watch(&source.Kind{Type: &icndbfunv1alpha1.FunApp{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner FunApp
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &icndbfunv1alpha1.FunApp{},
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &icndbfunv1alpha1.FunApp{},
	})

	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileFunApp implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileFunApp{}

// ReconcileFunApp reconciles a FunApp object
type ReconcileFunApp struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a FunApp object and makes changes based on the state read
// and what is in the FunApp.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileFunApp) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling FunApp")

	// Fetch the FunApp instance
	funapp := &icndbfunv1alpha1.FunApp{}
	err := r.client.Get(context.TODO(), request.NamespacedName, funapp)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	deployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: funapp.Name, Namespace: funapp.Namespace}, deployment)
	if err != nil && errors.IsNotFound(err) {
		// Define a new Deployment
		dep := r.deploymentForFunApp(funapp)
		reqLogger.Info("Creating a new Deployment.", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		err = r.client.Create(context.TODO(), dep)
		if err != nil {
			reqLogger.Error(err, "Failed to create new Deployment.", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return reconcile.Result{}, err
		}
		// Deployment created successfully - return and requeue
		// NOTE: that the requeue is made with the purpose to provide the deployment object for the next step to ensure the deployment size is the same as the spec.
		// Also, you could GET the deployment object again instead of requeue if you wish. See more over it here: https://godoc.org/sigs.k8s.io/controller-runtime/pkg/reconcile#Reconciler
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Deployment.")
		return reconcile.Result{}, err
	}

	size := funapp.Spec.Funpods
	if *deployment.Spec.Replicas != size {
		deployment.Spec.Replicas = &size
		err = r.client.Update(context.TODO(), deployment)
		if err != nil {
			reqLogger.Error(err, "Failed to update Deployment.", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
			return reconcile.Result{}, err
		}
	}

	service := &corev1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: funapp.Name, Namespace: funapp.Namespace}, service)
	if err != nil && errors.IsNotFound(err) {
		// Define a new Service object
		ser := r.serviceForFunApp(funapp)
		reqLogger.Info("Creating a new Service.", "Service.Namespace", ser.Namespace, "Service.Name", ser.Name)
		err = r.client.Create(context.TODO(), ser)
		if err != nil {
			reqLogger.Error(err, "Failed to create new Service.", "Service.Namespace", ser.Namespace, "Service.Name", ser.Name)
			return reconcile.Result{}, err
		}
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Service.")
		return reconcile.Result{}, err
	}

	// Update the Funapp status with the pod names
	// List the pods for this memcached's deployment
	podList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(labelsForFunApp(funapp.Name))
	listOps := &client.ListOptions{
		Namespace:     funapp.Namespace,
		LabelSelector: labelSelector,
	}
	err = r.client.List(context.TODO(), listOps, podList)
	if err != nil {
		reqLogger.Error(err, "Failed to list pods.", "Funapp.Namespace", funapp.Namespace, "Funapp.Name", funapp.Name)
		return reconcile.Result{}, err
	}
	podNames := getPodNames(podList.Items)

	// Update status.Nodes if needed
	if !reflect.DeepEqual(podNames, funapp.Status.Podnames) {
		funapp.Status.Podnames = podNames
		err := r.client.Status().Update(context.TODO(), funapp)
		if err != nil {
			reqLogger.Error(err, "Failed to update Funapp status.")
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil


}

// deploymentForFunApp returns a memcached Deployment object
func (r *ReconcileFunApp) deploymentForFunApp(fa *icndbfunv1alpha1.FunApp) *appsv1.Deployment {

	reqLogger := log.WithValues("Request.Namespace", fa.Namespace, "Request.Name", fa.Name)
	reqLogger.Info("creating deployment FunApp")

	ls := labelsForFunApp(fa.Name)
	replicas := fa.Spec.Funpods

	names := ""
	for _, param := range fa.Spec.Params {
		if strings.EqualFold(param.Key, "Name") {
			if names != "" {
				names +=  ","
			}
			names += param.Value

		}
	}
	args := []string{"--names",  names}

	reqLogger.Info("::::::::::::::::",  "ards:", args )

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fa.Name,
			Namespace: fa.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "nikhilvep/icndb-app:version-exp",
						Name: "icndb-server",
						Args: args,
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8000,
						}},
					}},
				},
			},
		},
	}
	// Set Funapp instance as the owner of the Deployment.
	controllerutil.SetControllerReference(fa, dep, r.scheme)
	return dep
}

// serviceForFunApp function takes in a Funapp object and returns a Service for that object.
func (r *ReconcileFunApp) serviceForFunApp(fa *icndbfunv1alpha1.FunApp) *corev1.Service {
	ls := labelsForFunApp(fa.Name)
	ser := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fa.Name,
			Namespace: fa.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Selector: ls,
			Ports: []corev1.ServicePort{
				{
					Protocol: corev1.ProtocolTCP,
					Port: 8000,
					Name: fa.Name,
				},
			},
		},
	}
	// Set Funapp instance as the owner of the Service.
	controllerutil.SetControllerReference(fa, ser, r.scheme)
	return ser
}

// labelsForFunApp returns the labels for selecting the resources
// belonging to the given memcached CR name.
func labelsForFunApp(name string) map[string]string {
	return map[string]string{"app": "memcached", "memcached_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}
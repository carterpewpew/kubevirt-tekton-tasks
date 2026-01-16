package constants

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type timeouts struct {
	Zero                      *metav1.Duration
	QuickTaskRun              *metav1.Duration
	SmallDVCreation           *metav1.Duration
	DefaultTaskRun            *metav1.Duration
	TaskRunExtraWaitDelay     *metav1.Duration
	PipelineRunExtraWaitDelay *metav1.Duration
	WaitBeforeExecutingVM     *metav1.Duration
	WaitForVMStart            *metav1.Duration
	TenSeconds                *metav1.Duration
}

var Timeouts = timeouts{
	Zero:                      &metav1.Duration{0 * time.Second},
	TaskRunExtraWaitDelay:     &metav1.Duration{3 * time.Minute},
	SmallDVCreation:           &metav1.Duration{10 * time.Minute},
	QuickTaskRun:              &metav1.Duration{1 * time.Minute},
	DefaultTaskRun:            &metav1.Duration{7 * time.Minute},
	WaitBeforeExecutingVM:     &metav1.Duration{15 * time.Second},
	WaitForVMStart:            &metav1.Duration{3 * time.Minute},
	PipelineRunExtraWaitDelay: &metav1.Duration{20 * time.Minute},
	TenSeconds:                &metav1.Duration{10 * time.Second},
}

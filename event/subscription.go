// Resubscribe applies backoff between calls to fn. The time between calls is adapted
// based on the error rate, but will never exceed backoffMax.
func Resubscribe(backoffMax time.Duration, fn ResubscribeFunc) Subscription {
	return ResubscribeErr(backoffMax, func(ctx context.Context, _ error) (Subscription, error) {
		return fn(ctx)
	})
}

// A ResubscribeFunc attempts to establish a subscription.
type ResubscribeFunc func(context.Context) (Subscription, error)

// ResubscribeErr calls fn repeatedly to keep a subscription established. When the
// subscription is established, ResubscribeErr waits for it to fail and calls fn again. This
// process repeats until Unsubscribe is called or the active subscription ends
// successfully.
//
// The difference between Resubscribe and ResubscribeErr is that with ResubscribeErr,
// the error of the failing subscription is available to the callback for logging
// purposes.
//
// ResubscribeErr applies backoff between calls to fn. The time between calls is adapted
// based on the error rate, but will never exceed backoffMax.
func ResubscribeErr(backoffMax time.Duration, fn ResubscribeErrFunc) Subscription {
	s := &resubscribeSub{
		waitTime:   backoffMax / 10,
		backoffMax: backoffMax,
@@ -106,15 +126,18 @@ func Resubscribe(backoffMax time.Duration, fn ResubscribeFunc) Subscription {
	return s
}

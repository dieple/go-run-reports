package notifier

import "log"

func SendUsageWarning(orgID string) {
	log.Printf("[Notify] Org %s is nearing usage limit", orgID)
}

func SendLimitReached(orgID string) {
	log.Printf("[Notify] Org %s has reached usage limit", orgID)
}
package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type LeadStatus string

const (
	LeadStatusNew         LeadStatus = "new"
	LeadStatusWorking     LeadStatus = "working"
	LeadStatusQualified   LeadStatus = "qualified"
	LeadStatusUnqualified LeadStatus = "unqualified"
)

func (ls LeadStatus) IsValid() bool {
	switch ls {
	case LeadStatusNew,
		LeadStatusWorking,
		LeadStatusQualified,
		LeadStatusUnqualified:
		return true
	}
	return false
}

func validateLeadStatus(fl validator.FieldLevel) bool {
	status := LeadStatus(strings.ToLower(fl.Field().String()))
	return status.IsValid()
}

type LeadSource string

const (
	LeadSourceWebsite         LeadSource = "website"
	LeadSourceReferral        LeadSource = "referral"
	LeadSourceLinkedIn        LeadSource = "linkedin"
	LeadSourceFacebook        LeadSource = "facebook"
	LeadSourceTwitter         LeadSource = "twitter"
	LeadSourceInstagram       LeadSource = "instagram"
	LeadSourceYouTube         LeadSource = "youtube"
	LeadSourceGoogleAds       LeadSource = "google ads"
	LeadSourceFacebookAds     LeadSource = "facebook ads"
	LeadSourceInstagramAds    LeadSource = "instagram ads"
	LeadSourceLinkedInAds     LeadSource = "linkedin ads"
	LeadSourceEmailCampaign   LeadSource = "email campaign"
	LeadSourceColdCall        LeadSource = "cold call"
	LeadSourceDirectVisit     LeadSource = "direct visit"
	LeadSourceWebinar         LeadSource = "webinar"
	LeadSourceConference      LeadSource = "conference"
	LeadSourceNetworkingEvent LeadSource = "networking event"
	LeadSourcePartnership     LeadSource = "partnership"
	LeadSourceReseller        LeadSource = "reseller"
	LeadSourceInbound         LeadSource = "inbound"
	LeadSourceOutbound        LeadSource = "outbound"
	LeadSourceLandingPage     LeadSource = "landing page"
	LeadSourceLiveChat        LeadSource = "live chat"
	LeadSourceCustomerSupport LeadSource = "customer support"
	LeadSourceSEO             LeadSource = "seo"
	LeadSourceOrganicSearch   LeadSource = "organic search"
	LeadSourcePaidSearch      LeadSource = "paid search"
	LeadSourceOther           LeadSource = "other"
)

func (ls LeadSource) IsValid() bool {
	switch ls {
	case LeadSourceWebsite,
		LeadSourceReferral,
		LeadSourceLinkedIn,
		LeadSourceFacebook,
		LeadSourceTwitter,
		LeadSourceInstagram,
		LeadSourceYouTube,
		LeadSourceGoogleAds,
		LeadSourceFacebookAds,
		LeadSourceInstagramAds,
		LeadSourceLinkedInAds,
		LeadSourceEmailCampaign,
		LeadSourceColdCall,
		LeadSourceDirectVisit,
		LeadSourceWebinar,
		LeadSourceConference,
		LeadSourceNetworkingEvent,
		LeadSourcePartnership,
		LeadSourceReseller,
		LeadSourceInbound,
		LeadSourceOutbound,
		LeadSourceLandingPage,
		LeadSourceLiveChat,
		LeadSourceCustomerSupport,
		LeadSourceSEO,
		LeadSourceOrganicSearch,
		LeadSourcePaidSearch,
		LeadSourceOther:
		return true
	}
	return false
}

func validateLeadSource(fl validator.FieldLevel) bool {
	source := LeadSource(strings.ToLower(fl.Field().String()))
	return source.IsValid()
}

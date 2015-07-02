package main

import (
	"fmt"

	"github.com/creack/log"
)

var errProducer = log.NewProducer().SetDefaultLevel(log.LevelError)

// Common errors.
var (
	ErrMissingURL = errProducer.NewMessage("Missing URL").SetCode(100).SetName("missing_url")
)

// // Internal Errors, These errors all report as "Internal Error" so as not to
// // divulge any information to the client
// var (
// 	ErrorInternal             = InternalError{"Unknown Error", SimpleError{"unknown_error", m, 999}}
// 	ErrorInvalidPageURL       = InternalError{"Invalid Page URL", SimpleError{"invalid_page_url", m, 998}}
// 	ErrorInvalidReferrerURL   = InternalError{"Invalid Referrer URL", SimpleError{"invalid_referrer_url", m, 997}}
// 	ErrorMongoQuery           = InternalError{"Mongo Query Error", SimpleError{"mongo_query_error", m, 996}}
// 	ErrorQueryString          = InternalError{"Error parsing query string", SimpleError{"invalid_query_string", m, 995}}
// 	ErrorMarshallingJSON      = InternalError{"Could not marshal JSON", SimpleError{"invalid_json", m, 994}}
// 	ErrorSocialData           = InternalError{"Could not parse Social Data", SimpleError{"invalid_social_data", m, 994}}
// 	ErrorMissingEvent         = InternalError{"Missing event", SimpleError{"missing_event", m, 993}} // TODO: this should be a user error
// 	ErrorInvalidEvent         = InternalError{"Invalid event", SimpleError{"invalid_event", m, 992}}
// 	ErrorMissingContentID     = InternalError{"Missing Content ID (cid)", SimpleError{"missing_content_id", m, 991}}
// 	ErrorInvalidContentID     = InternalError{"Invalid Content ID (cid)", SimpleError{"invalid_content_id", m, 990}}
// 	ErrorMissingSocialNetwork = InternalError{"Missing Social Network", SimpleError{"missing_social_network", m, 989}}
// 	ErrorInvalidSocialNetwork = InternalError{"Invalid Social Network", SimpleError{"invalid_social_network", m, 988}}
// 	ErrorMissingConversionID  = InternalError{"Missing Conversion ID", SimpleError{"missing_conversion_id", m, 987}}
// 	ErrorInvalidConversionID  = InternalError{"Invalid Conversion ID", SimpleError{"invalid_conversion_id", m, 986}}
// 	ErrorDatastoreDown        = InternalError{"Datastore down", SimpleError{"datastore_down", m, 985}}
// )

// // Client Errors, These errors give information back to the client to take action
// var (
// 	ErrorMissingURL         = SimpleError{"missing_url", "Missing url", 100}
// 	ErrorInvalidURL         = SimpleError{"invalid_url", "Invalid url", 101}
// 	ErrorMissingDate        = SimpleError{"missing_published_date", "Missing published date (date)", 102}
// 	ErrorInvalidDate        = SimpleError{"invalid_published_date", "Invalid published date (date)", 103}
// 	ErrorMissingAccountID   = SimpleError{"missing_account_id", "Missing Account ID (pid)", 104}
// 	ErrorInvalidAccountID   = SimpleError{"invalid_account_id", "Invalid Account ID (pid)", 105}
// 	ErrorAccountNotFound    = SimpleError{"account_not_found", "Account not found", 106}
// 	ErrorAccountInactive    = SimpleError{"inactive_account", "Account inactive", 107}
// 	ErrorInvalidDomain      = SimpleError{"invalid_domain", "Invalid domain for this account", 108}
// 	ErrorMissingTOSTick     = SimpleError{"missing_tos_tick", "Missing TimeOnSite Tick", 109}
// 	ErrorMissingEngagedTick = SimpleError{"missing_engaged_tick", "Missing Engaged Tick", 110}
// 	ErrorInvalidTickValue   = SimpleError{"invalid_tick_value", "Invalid Tick Value", 111}
// 	ErrorMissingSD          = SimpleError{"missing_ntile", "Missing sd", 112}
// )

func fail() error {
	// return ErrMissingURL
	return fmt.Errorf("FAILERROR")
}

func test() error {
	if err := fail(); err != nil {
		return log.NewError(err)
	}
	return nil
}

func test2() error {
	if err := test(); err != nil {
		return log.NewError(err)
	}
	return nil
}

func main() {
	if err := test2(); err != nil {
		buf, _ := err.(*log.Message).Dump()
		fmt.Printf("%s\n", buf)
		fmt.Printf("%s\n", err)
	}
}

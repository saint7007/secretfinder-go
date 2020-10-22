package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	data, err := ioutil.ReadFile("private.key")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Println("Contents of file:", string(data))
	allRegex := make(map[string]string)
	allRegex["google_api"] = "AIza[0-9A-Za-z-_]{35}"
	allRegex["firebase"] = "AAAA[A-Za-z0-9_-]{7}:[A-Za-z0-9_-]{140}"
	allRegex["google_captcha"] = "6L[0-9A-Za-z-_]{38}|^6[0-9a-zA-Z_-]{39}$"
	allRegex["google_oauth"] = "ya29\\.[0-9A-Za-z\\-_]+"
	allRegex["amazon_aws_access_key_id"] = "AKIA[0-9A-Z]{16}"
	allRegex["amazon_mws_auth_toke"] = "amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
	allRegex["amazon_aws_url"] = "s3\\.amazonaws.com[/]+|[a-zA-Z0-9_-]*\\.s3\\.amazonaws.com"
	allRegex["amazon_aws_url2"] = "([a-zA-Z0-9-\\.\\_]+\\.s3\\.amazonaws\\.com|s3://[a-zA-Z0-9-\\.\\_]|s3-[a-zA-Z0-9-\\.\\_\\//]|s3.amazonaws.com/[a-zA-Z0-9-\\.\\_]|s3.console.aws.amazon.com/s3/buckets/[a-zA-Z0-9-\\.\\_]+)"
	allRegex["facebook_access_token"] = "EAACEdEose0cBA[0-9A-Za-z]+"
	allRegex["authorization_basic"] = "basic [a-zA-Z0-9=:_\\+\\/-]{5,100}"
	allRegex["authorization_bearer"] = "bearer [a-zA-Z0-9_\\-\\.=:_\\+\\//]{5,100}"
	allRegex["authorization_api"] = "api[key|_key|\\s+]+[a-zA-Z0-9_\\-]{5,100}"
	allRegex["mailgun_api_key"] = "key-[0-9a-zA-Z]{32}"
	allRegex["twilio_api_key"] = "SK[0-9a-fA-F]{32}"
	allRegex["twilio_account_sid"] = "AC[a-zA-Z0-9_\\-]{32}"
	allRegex["twilio_app_sid"] = "AP[a-zA-Z0-9_\\-]{32}"
	allRegex["paypal_braintree_access_token"] = "access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}"
	allRegex["square_oauth_secret"] = "sq0csp-[ 0-9A-Za-z\\-_]{43}|sq0[a-z]{3}-[0-9A-Za-z\\-_]{22,43}"
	allRegex["square_access_token"] = "sqOatp-[0-9A-Za-z\\-_]{22}|EAAA[a-zA-Z0-9]{60}"
	allRegex["stripe_standard_api"] = "sk_live_[0-9a-zA-Z]{24}"
	allRegex["stripe_restricted_api"] = "rk_live_[0-9a-zA-Z]{24}"
	allRegex["github_access_token"] = "[a-zA-Z0-9_-]*:[a-zA-Z0-9_\\-]+@github\\.com*"
	allRegex["rsa_private_key"] = "-----BEGIN RSA PRIVATE KEY-----"
	allRegex["ssh_dsa_private_key"] = "-----BEGIN DSA PRIVATE KEY-----"
	allRegex["ssh_dc_private_key"] = "-----BEGIN EC PRIVATE KEY-----"
	allRegex["pgp_private_block"] = "-----BEGIN PGP PRIVATE KEY BLOCK-----"
	allRegex["json_web_token"] = "ey[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*$"
	allRegex["slack_token"] = "\"api_token\":\"(xox[a-zA-Z]-[a-zA-Z0-9-]+)\""
	allRegex["SSH_privKey"] = "([-]+BEGIN [^\\s]+ PRIVATE KEY[-]+[\\s]*[^-]*[-]+END [^\\s]+ PRIVATE KEY[-]+)"
	allRegex["possible_Creds"] = "(?i)(password\\s*[`=:\"]+\\s*[^\\s]+|password is\\s*[`=:\"]*\\s*[^\\s]+|pwd\\s*[`=:\"]*\\s*[^\\s]+|passwd\\s*[`=:\"]+\\s*[^\\s]+)"

	for id, regexstring := range allRegex {
		match, _ := regexp.MatchString(regexstring, string(data))
		//fmt.Println(match)
		if match == true {
			fmt.Println(id, regexstring)
		}

	}

}

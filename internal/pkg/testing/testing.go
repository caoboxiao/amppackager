/*
Package testing provides test constants and utilities for the transformer and printer.

*/
package testing

import "strings"

// TestCase stores a description of the test case, the input HTML, and
// expected output HTML.
type TestCase struct {
	Desc     string
	Input    string
	Expected string
}

// AMPHTML constants.
const (
	// LinkFavicon is an example link tag
	LinkFavicon = "<link href=https://example.com/favicon.ico rel=icon>"

	// LinkResourceHint is an injected tag to preconnect certain domains.
	LinkResourceHint = "<link crossorigin href=https://fonts.gstatic.com rel=\"dns-prefetch preconnect\">"

	// LinkStylesheetGoogleFont is a link tag for a Google Font
	LinkStylesheetGoogleFont = "<link href=https://fonts.googleapis.com/css?family=Roboto rel=stylesheet>"

	// MetaCharset is a required tag for an AMP document.
	MetaCharset = "<meta charset=utf-8>"

	// MetaViewpoert is a required tag for an AMP document.
	MetaViewport = "<meta content=width=device-width,minimum-scale=1,initial-scale=1 name=viewport>"

	// NoscriptAMPBoilerplate is the standard style for <noscript> tag.
	NoscriptAMPBoilerplate = "<noscript><style amp-boilerplate>body{-webkit-animation:none;-moz-animation:none;-ms-animation:none;animation:none}</style></noscript>"

	// ScriptAMPAudio is the script for amp-audio.
	ScriptAMPAudio = "<script async custom-element=amp-audio src=https://cdn.ampproject.org/v0/amp-audio-0.1.js></script>"

	// ScriptAMPDynamicCSSClasses is the script for amp-dynamic-css-class.
	ScriptAMPDynamicCSSClasses = "<script async custom-element=amp-dynamic-css-classes src=https://cdn.ampproject.org/v0/amp-dynamic-css-classes-0.1.js></script>"

	// ScriptAMPExperiment is the script for amp-experiment.
	ScriptAMPExperiment = "<script async custom-element=amp-experiment src=https://cdn.ampproject.org/v0/amp-experiment-0.1.js></script>"

	// ScriptAMPMustache is the script for amp-mustache.
	ScriptAMPMustache = "<script async custom-template=amp-mustache src=https://cdn.ampproject.org/v0/amp-mustache-0.1.js></script>"

	// ScriptAMPRuntime is the AMP script tag.
	ScriptAMPRuntime = "<script async src=https://cdn.ampproject.org/v0.js></script>"

	// ScriptAMP4AdsRuntime is the AMP4Ads script tag.
	ScriptAMP4AdsRuntime = "<script async src=https://cdn.ampproject.org/amp4ads-v0.js></script>"

	// ScriptAMPStory is the script for amp-story.
	ScriptAMPStory = "<script async custom-element=amp-story src=https://cdn.ampproject.org/v0/amp-story-0.1.js></script>"

	// StyleAMP4AdsBoilerplate is the script for amp4ads boilerplate.
	StyleAMP4AdsBoilerplate = "<style amp4ads-boilerplate>body{visibility;hidden}</style>"

	// StyleAMP4EmailBoilerplate is the script for amp4email boilerplate.
	StyleAMP4EmailBoilerplate = "<style amp4email-boilerplate>body{visibility;hidden}</style>"

	// StyleAMPBoilerplate is the standard style.
	StyleAMPBoilerplate = "<style amp-boilerplate>body{-webkit-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-moz-animation:-amp-start 8s steps(1,end) 0s 1 normal both;-ms-animation:-amp-start 8s steps(1,end) 0s 1 normal both;animation:-amp-start 8s steps(1,end) 0s 1 normal both}@-webkit-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-moz-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-ms-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@-o-keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}@keyframes -amp-start{from{visibility:hidden}to{visibility:visible}}</style>"

	// StyleAMPCustom is a customized stylesheet for an AMP document.
	StyleAMPCustom = "<style amp-custom>#lemur { color: #adaaad }</style>"

	// StyleAMPRuntime is an injected tag from server-side rendering.
	StyleAMPRuntime = "<style amp-runtime></style>"

	// Title is a title tag for an AMP document.
	Title = "<title>Hello AMP</title>"
)

// Concat concatenates the given strings together.
func Concat(strs ...string) string {
	return strings.Join(strs, "")
}

// BuildHTML returns AMPHTML with the given body string. Note this isn't
// a valid AMP document.
func BuildHTML(body string) string {
	return Concat(
		"<!doctype html><html ⚡><head>",
		ScriptAMPRuntime,
		LinkFavicon,
		"</head><body>",
		body,
		"</body></html>")
}
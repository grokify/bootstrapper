package bootstrap

import (
	"fmt"
	"strings"

	"github.com/grokify/gotilla/math/bigint"
)

type Alert int

const alertPrefix string = "alert-"

const (
	Primary Alert = iota
	Secondary
	Success
	Danger
	Warning
	Info
	Light
	Dark
)

var alerts = [...]string{
	"Primary",
	"Secondary",
	"Success",
	"Danger",
	"Warning",
	"Info",
	"Light",
	"Dark",
}

func NewAlert(idx int) Alert {
	return Alert(bigint.Int64Mod(int64(idx), int64(len(alerts))))
}

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a Alert) String() string {
	//if Primary <= a && a <= Dark {
	if a >= Primary && a <= Dark {
		return alerts[a]
	}
	a2 := NewAlert(int(a))
	return alerts[a2]
}

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a Alert) Class() string {
	return alertPrefix + strings.ToLower(a.String())
}

type AlertBordered int

const (
	PrimaryB AlertBordered = iota
	SecondaryB
	SuccessB
	DangerB
	WarningB
	InfoB
	DarkB
)

var alertsbordered = [...]string{
	"Primary",
	"Secondary",
	"Success",
	"Danger",
	"Warning",
	"Info",
	"Dark",
}

func NewAlertBordered(idx int) AlertBordered {
	return AlertBordered(bigint.Int64Mod(int64(idx), int64(len(alertsbordered))))
}

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a AlertBordered) String() string {
	if a >= PrimaryB && a <= DarkB {
		return alertsbordered[a]
	}
	a2 := NewAlertBordered(int(a))
	return alertsbordered[a2]
}

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a AlertBordered) Class() string {
	return alertPrefix + strings.ToLower(a.String())
}

const alertDivFormat string = `<div class="alert alert-%s" role="alert">`

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a AlertBordered) DivHtml(innerHtml string) string {
	alertString := a.String()
	begin := fmt.Sprintf(alertDivFormat, strings.ToLower(alertString))
	return begin + innerHtml + "</div>"
}

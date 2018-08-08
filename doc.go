// Package errors implements one opinion as to how errors should work.
//
// Definitions:

//	error: an error is a condition of active degradation in service,
//	and should carry a concrete action that an on-call engineer or ops
//	personnel should carry out. This is probably a link to a runbook.
//
//	warning: a warning indicates a situation that may devolve into an
//	error but for which there is no concrete action to mitigate. This
//	is a sign that maybe the on-call engineer or ops should start keeping
//	an eye on this.
//
//	info: everything else is just information that may or may not be
//	particularly useful.
//
// In a logging package that I wrote, I defined the different logging levels
// as:
//
//	FATAL: the system is in an unsuable state, and cannot continue to run.
//	It's expected that shortly after logging a fatal message, the program
//	will halt.
//
//	CRITICAL: critical conditions. The error, if uncorrected, is likely
//	to cause a fatal condition shortly. An example is running out of disk
//	space. This is something that the ops team should get paged for.
//
//	ERROR: error conditions. A single error doesn't require an ops team to
//	be paged, but repeated errors should typically trigger a page based
//	on threshold triggers. An example is a network failure: it might be a
//	transient failure (these do happen), but most of the time it's self-
//	correcting.
//
//	WARNING: warning conditions. An example of this is a bad request sent
//	to a server. This isn't an error on the part of the program, but it
//	may be indicative of other things. Like errors, the ops team shouldn't
//	be paged for errors, but a page might be triggered if a certain
//	threshold of warnings is reached (which is typically much higher than
//	errors). For example, repeated warnings might be a sign that the
//	system is under attack.
//
//	INFO: informational message. This is a normal log message that is used
//	to deliver information, such as recording requests. Ops teams are
//	never paged for informational messages. This is the default log level.
//
//	DEBUG: debug-level message. These are only used during development or
//	if a deployed system repeatedly sees abnormal errors. It's likely that
//	only a fraction of these are actually being recorded.
//
// This model presumes a fault-tolerant application where the ERROR log level
// doesn't necessarily mean an ongoing degradation in service. There's also
// the distinction of service degradation and system degredation.
package errors

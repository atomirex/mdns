mdns
====

Modified fork of hashicorp/mdns to suit a slightly different use case.

This is more about going into a long running dns-sd monitoring service,
so the need is to keep scanning and populating all entries persistently.

Right now this is a hack - there is no clean backoff for entries, and it
simply polls quite aggressively to try to maintain status.

Also doesn't track things going away.
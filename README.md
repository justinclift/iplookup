# IP Lookup

This is a simple command line utility to display the
country of a given IP address, using the SQLite Geo-IP
database from [here](https://dbhub.io/justinclift/Geo-IP.sqlite).

You need to download the above database first, and have it
in the same directory this command runs from.

### Examples

```
$ iplookup 1.2.3.4
AUS
$ iplookup 4.3.2.1
USA
```

## References

The data for the Geo-IP SQLite database comes from:

&nbsp; &nbsp; http://software77.net/geo-ip/

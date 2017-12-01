ntped
=====
ntped gets time offsets from below NTP servers,
* asia.pool.ntp.org
* euroupe.pool.ntp.org
* north-america.pool.ntp.org
* oceania.pool.ntp.org

Usage
----
To obtain current ntped time, sync with clocks in pool.ntp.org first.
```go
err := ntped.Sync(0, 0)
time := ntped.Now()     // NTPed time in time.Time
```
Sync() queries time offsets (NTP_server_time - your_machine_time) to NTP servers.<br>
Now() returns median_offset + your_machine_time.

Sometimes querying to NTP pool takes time. Drop time consuming NTP servers by timeout.
```go
err := ntped.Sync(0, 5000)  // NTP query expires after 5 seconds
```
If you do not want to drop queries, set timeout as 0.

ntped can retry Sync() when all the queries are dropped.
```go
err := ntped.Sync(3, 2000) // retry Sync() maximum three times
```
On each try(n), Sync() increase timeout(t(n)) lienearly, t(n) = n * t.<br>
For example, when you set timeout(t) as 2000 (2 seconds), timeouts are
* initial try: 2000
* 1st retry:   4000
* 2nd retry:   6000
* 3rd retry:   8000

<br>
If you are not satisfied with Unix() and UnixNano(), you can use UnixMilli() for ntped time.

```go
err := ntped.Sync(0, 0)
ms := ntped.UnixMilli() // ntped.Now() in milli seconds (int64)
```

Dependencies
-----
* github.com/beevik/ntp

License
----
MIT

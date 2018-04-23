ntped
=====
ntped gets time offsets from 16 NTP servers,
* 0.asia.pool.ntp.org ~ 3.asia.pool.ntp.org
* 0.europe.pool.ntp.org ~ 3.europe.pool.ntp.org
* 0.oceania.pool.ntp.org ~ 3.oceania.pool.ntp.org
* 0.north-america.pool.ntp.org ~ 3.north-america.pool.ntp.org

Each zone is supported by at least 100+ servers, respectively.<br>
You can check current pool size here, http://www.pool.ntp.org/zone/@.

Usage
----
To obtain current ntped time, sync with clocks in pool.ntp.org first.
```go
err := ntped.Sync(0, 0)
now := ntped.Now()     // NTPed time in time.Time
```
Sync() queries time offsets(NTP_server_time - your_machine_time) to NTP servers.<br>
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
On each try, Sync() linearly increases timeout, t(n) = t * (n + 1), where *t* is the initial timeout and *n* is the retry number.<br>
In the above example, ```ntped.Sync(3, 2000)```, query timeouts are:
* initial try (n = 0): 2s
* 1st retry   (n = 1): 4s
* 2nd retry   (n = 2): 6s
* 3rd retry   (n = 3): 8s

<br>
If you are not satisfied with Now(), you can use UnixMilli() instead.

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
